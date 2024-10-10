package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/constants/queue"
	"project.io/shop/model"
	"project.io/shop/queue/biz/producer/shop"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func DoOrder(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrder

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	lock := facades.Locker.NewMutex(locker.Keys("order", auth.ID(ctx)))

	if err := lock.LockContext(c); err != nil {
		http.Fail(ctx, "Order locking failed. Please try again later.")
		return
	}

	defer lock.UnlockContext(c)

	var shipping model.ShpShipping

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&shipping, "`id`=?", request.Shipping)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "No available shipping companies found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "Shipping company query failed. Please try again later.")
		return
	}

	var carts []model.ShpCart

	facades.Gorm.
		Preload("Product").
		Preload("SKU").
		Find(&carts, "`user_id`=?", auth.ID(ctx))

	if len(carts) == 0 {
		http.Fail(ctx, "Your cart is empty. Please add items before placing an order.")
		return
	}

	if count := lo.CountBy(carts, func(item model.ShpCart) bool { return item.IsInvalid == util.Yes }); count > 0 {
		http.Fail(ctx, "Some products are no longer available, and the order has failed.")
		return
	}

	if count := lo.CountBy(carts, func(item model.ShpCart) bool { return item.SKU == nil }); count > 0 {
		http.Fail(ctx, "Some items have become unavailable, and the order could not be completed.")
		return
	}

	if count := lo.CountBy(carts, func(item model.ShpCart) bool { return item.Price <= 0 }); count > 0 {
		http.Fail(ctx, "Some products have abnormal prices, and the order has failed.")
		return
	}

	if count := lo.CountBy(carts, func(item model.ShpCart) bool { return item.Quantity > item.SKU.Stock }); count > 0 {
		http.Fail(ctx, "Some products are out of stock, and the order has failed.")
		return
	}

	order := model.ShpOrder{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		UserID:         auth.ID(ctx),
		CostShipping:   shipping.Money,
		Prices:         shipping.Money,
		Status:         model.ShpOrderOfStatusPay,
		Remark:         request.Remark,
		IsPaid:         util.No,
		IsInvoice:      util.No,
		IsAppraisal:    util.No,
	}

	details := make([]model.ShpOrderDetail, len(carts))

	for idx, item := range carts {

		details[idx] = model.ShpOrderDetail{
			ID:             facades.Snowflake.Generate().String(),
			Platform:       order.Platform,
			CliqueID:       order.CliqueID,
			OrganizationID: order.OrganizationID,
			UserID:         order.UserID,
			OrderID:        order.ID,
			ProductID:      item.ProductID,
			SkuID:          item.SkuID,
			Name:           item.Name,
			Specifications: item.Specifications,
			Picture:        item.Picture,
			Price:          item.SKU.Price,
			CostPrice:      item.SKU.CostPrice,
			Quantity:       item.Quantity,
			TotalPrice:     item.SKU.Price * item.Quantity,
			CouponPrice:    0,
			CostPrices:     item.SKU.CostPrice * item.Quantity,
			Prices:         0,
			//Weight:         item.SKU.Price * item.Quantity,
			IsInvoiced: util.No,
		}

		details[idx].Prices = details[idx].TotalPrice + details[idx].CostPrice - details[idx].CouponPrice

		order.TotalPrice += details[idx].TotalPrice
		order.CouponPrice += details[idx].CouponPrice
		order.CostPrices += details[idx].CostPrices

		order.Prices += details[idx].Prices
	}

	shipment := model.ShpShipment{
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		OrderID:        order.ID,
		UserID:         order.UserID,
		ShippingID:     shipping.ID,
		Money:          shipping.Money,
		Company:        shipping.Name,
	}

	address := model.ShpOrderAddress{
		ID:             0,
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		OrderID:        order.ID,
		UserID:         order.UserID,
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		Company:        request.Company,
		Country:        request.Country,
		Prefecture:     request.Prefecture,
		City:           request.City,
		Street:         request.Street,
		Detail:         request.Detail,
		Postcode:       request.Postcode,
		Phone:          request.Phone,
		Email:          request.Email,
	}

	log := model.ShpLog{
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		UserID:         order.UserID,
		OrderID:        order.ID,
		Action:         "order",
		Content:        "Place Order",
	}

	tx := facades.Gorm.Begin()

	for _, item := range details {

		if result := tx.Model(&model.ShpSku{}).Where("`id`=? and `stock`>=?", item.SkuID, item.Quantity).Update("stock", gorm.Expr("`stock`-?", item.Quantity)); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "Insufficient stock for the product, unable to place the order.")
			return
		}
	}

	if result := tx.Create(&order); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	if result := tx.Create(&details); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	if result := tx.Create(&address); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	if result := tx.Create(&shipment); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	if result := tx.Delete(&model.ShpCart{}, "`user_id`=?", order.UserID); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	payment := model.ShpPayment{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		UserID:         auth.ID(ctx),
		OrderID:        order.ID,
		Channel:        request.Payment,
		Money:          order.Prices,
		IsConfirmed:    util.No,
		ExpiredAt:      order.CreatedAt.AddMinutes(10),
	}

	if result := tx.Create(&payment); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	tx.Commit()

	if err := shop.PublishOrderClosed(queue.ShopOrderClosedMessage{
		UserID: order.UserID,
		Order:  order.ID,
	}); err != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	responses := res.DoOrder{
		PayID:   payment.ID,
		Channel: request.Payment,
	}

	http.Success(ctx, responses)
}

func ToOrderOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrderOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToOrderOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.
		WithContext(c).
		Where("`user_id`=?", auth.ID(ctx))

	if request.Status != "" {

		if request.Status == "evaluate" {
			tx = tx.Where("`status`=? and `is_appraisal`=?", model.ShpOrderOfStatusCompleted, util.No)
		} else {
			tx = tx.Where("`status`=?", request.Status)
		}
	}

	if request.IsAppraisal > 0 {
		tx = tx.Where("`is_appraisal`=?", request.IsAppraisal)
	}

	tx.Model(&model.ShpOrder{}).Count(&responses.Total)

	if responses.Total > 0 {

		now := carbon.Now()

		var orders []model.ShpOrder

		tx.
			//Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Preload("Details").
			Order("`created_at` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&orders)

		responses.Data = make([]res.ToOrderOfPaginate, len(orders))

		for idx, item := range orders {

			responses.Data[idx] = res.ToOrderOfPaginate{
				ID:          item.ID,
				Details:     make([]res.ToOrderOfDetail, len(item.Details)),
				Prices:      item.Prices,
				Status:      item.Status,
				IsAppraisal: item.IsAppraisal,
				CanService:  util.No,
				CreateAt:    item.CreatedAt.ToDateTimeString(),
			}

			for key, val := range item.Details {
				responses.Data[idx].Details[key] = res.ToOrderOfDetail{
					ID:             val.ProductID,
					Name:           val.Name,
					Picture:        val.Picture,
					Price:          val.Price,
					Quantity:       val.Quantity,
					Specifications: val.Specifications,
				}
			}

			//if item.Organization != nil {
			//	responses.Data[idx].Organization = res.ToOrderOfOrganization{
			//		ID:   *item.OrganizationID,
			//		Name: item.Organization.Name,
			//	}
			//}

			if item.CompletedAt != nil && item.CompletedAt.AddDays(7).Gt(now) {
				responses.Data[idx].CanService = util.Yes
			}
		}
	}

	http.Success(ctx, responses)
}

func ToOrderOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrderOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var order model.ShpOrder

	fo := facades.Gorm.
		//Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Preload("Details").
		Preload("Address").
		Preload("Payment").
		//Preload("Invoice").
		//Preload("Shipment").
		//Preload("Shipping").
		//Preload("Services.Products").
		Preload("Logs", func(t *gorm.DB) *gorm.DB { return t.Order("`id` desc") }).
		First(&order, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Order not found.")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "Order query failed. Please try again later.")
		return
	}

	responses := res.ToOrderOfInformation{
		ID:           order.ID,
		CostShipping: order.CostShipping,
		TotalPrice:   order.TotalPrice,
		CouponPrice:  order.CouponPrice,
		Prices:       order.Prices,
		Refund:       order.Refund,
		Status:       order.Status,
		Remark:       order.Remark,
		IsInvoice:    order.IsInvoice,
		IsAppraisal:  order.IsAppraisal,
		CanService:   util.No,
		CreateAt:     order.CreatedAt.ToDateTimeString(),
	}

	now := carbon.Now()

	if order.CompletedAt != nil && order.CompletedAt.AddDays(7).Gt(now) {
		responses.CanService = util.Yes
	}

	//if order.Organization != nil {
	//	responses.Organization = res.ToOrderOfOrganization{
	//		ID:   order.Organization.ID,
	//		Name: order.Organization.Name,
	//	}
	//}

	if order.Payment != nil {
		responses.Payment = &res.ToOrderOfPayment{
			ID:      order.Payment.ID,
			Channel: order.Payment.Channel,
		}

		if order.Payment.No != nil {
			responses.Payment.No = *order.Payment.No
		}

		if order.Payment.PaidAt != nil {
			responses.Payment.PaidAt = order.Payment.PaidAt.ToDateTimeString()
		}
	}

	//if order.Shipping != nil {
	//	responses.Shipping = order.Shipping.Name
	//}

	if order.Address != nil {
		responses.Address = res.ToOrderOfAddress{
			FirstName:  order.Address.FirstName,
			LastName:   order.Address.LastName,
			Company:    order.Address.Company,
			Country:    order.Address.Country,
			Prefecture: order.Address.Prefecture,
			City:       order.Address.City,
			Street:     order.Address.Street,
			Detail:     order.Address.Detail,
			Postcode:   order.Address.Postcode,
			Phone:      order.Address.Phone,
			Email:      order.Address.Email,
		}
	}

	if len(order.Details) > 0 {

		responses.Details = make([]res.ToOrderOfDetail, len(order.Details))

		for idx, item := range order.Details {
			responses.Details[idx] = res.ToOrderOfDetail{
				ID:       item.ID,
				Name:     item.Name,
				Picture:  item.Picture,
				Price:    item.Price,
				Quantity: item.Quantity,
				//Prices:         item.Prices,
				Specifications: item.Specifications,
				Refund:         item.Refund,
				Returned:       item.Returned,
			}
		}
	}

	//if order.IsInvoice == util.Yes && order.Invoice != nil {
	//	responses.Invoice = &res.ToOrderOfInvoice{
	//		Type:      order.Invoice.Type,
	//		Header:    order.Invoice.Header,
	//		No:        order.Invoice.No,
	//		Bank:      order.Invoice.Bank,
	//		Card:      order.Invoice.Card,
	//		Address:   order.Invoice.Address,
	//		Telephone: order.Invoice.Telephone,
	//		Status:    order.Invoice.Status,
	//		Files:     order.Invoice.Files,
	//		Reason:    order.Invoice.Reason,
	//		Remark:    order.Invoice.Remark,
	//		CreatedAt: order.Invoice.CreatedAt.ToDateTimeString(),
	//	}
	//}

	if len(order.Logs) > 0 {

		responses.Logs = make([]res.ToOrderOfLog, len(order.Logs))

		for idx, item := range order.Logs {
			responses.Logs[idx] = res.ToOrderOfLog{
				Action:    item.Action,
				Content:   item.Content,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	//if order.Shipment != nil {
	//
	//	responses.Shipment = &res.ToOrderOfShipment{
	//		Company:   order.Shipment.Company,
	//		No:        order.Shipment.No,
	//		Remark:    order.Shipment.Remark,
	//		CreatedAt: order.Shipment.CreatedAt.ToDateTimeString(),
	//	}
	//}
	//
	//if len(order.Services) > 0 {
	//
	//	responses.Services = make([]res.ToOrderOfService, len(order.Services))
	//
	//	for idx, item := range responses.Details {
	//
	//		for _, val := range order.Services {
	//
	//			if lo.ContainsBy(val.Products, func(val model.ShpServiceProduct) bool { return item.ID == val.DetailID }) {
	//
	//				responses.Details[idx].IsServiced = util.Yes
	//
	//				if val.Status != model.ShpServiceOfStatusFinish && val.Status != model.ShpServiceOfStatusClosed {
	//					responses.Details[idx].Service = &val.ID
	//				}
	//			}
	//		}
	//	}
	//
	//	for idx, item := range order.Services {
	//
	//		responses.Services[idx] = res.ToOrderOfService{
	//			ID:     item.ID,
	//			Type:   item.Type,
	//			Status: item.Status,
	//			Detail: item.DetailID,
	//			Reason: item.Reason,
	//			Details: lo.Map(item.Products, func(val model.ShpServiceProduct, index int) res.ToOrderOfServiceWithDetail {
	//				return res.ToOrderOfServiceWithDetail{
	//					ID:       val.DetailID,
	//					Quantity: val.Quantity,
	//				}
	//			}),
	//			CreatedAt: item.CreatedAt.ToDateTimeString(),
	//		}
	//	}
	//}

	http.Success(ctx, responses)
}
