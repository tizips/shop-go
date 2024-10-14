package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/order"
	res "project.io/shop/admin/http/response/shop/store/order"
	"project.io/shop/constants/queue"
	"project.io/shop/model"
	"project.io/shop/queue/biz/producer/shop"
)

func ToServiceOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToServiceOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToServiceOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		WithContext(c)

	if request.ID != "" {
		tx = tx.Where("`id`=?", request.ID)
	}

	if request.Order != "" {
		tx = tx.Where("`order_id`=?", request.Order)
	}

	if request.Status != "" {
		tx = tx.Where("`status`=?", request.Status)
	}

	tx.Model(&model.ShpService{}).Count(&responses.Total)

	if responses.Total > 0 {

		var service []model.ShpService

		tx.
			//Preload("Refund").
			Preload("Products.Detail").
			Order("`created_at` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&service)

		responses.Data = make([]res.ToServiceOfPaginate, len(service))

		for idx, item := range service {

			responses.Data[idx] = res.ToServiceOfPaginate{
				ID:        item.ID,
				Order:     item.OrderID,
				Type:      item.Type,
				Details:   make([]res.ToOrderOfDetail, len(item.Products)),
				Status:    item.Status,
				Result:    item.Result,
				Reason:    item.Reason,
				Subtotal:  item.Subtotal,
				Shipping:  item.Shipping,
				Refunds:   item.Refunds,
				Pictures:  item.Pictures,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.ShipmentUser != nil {
				responses.Data[idx].ShipmentUser = &res.ToServiceOfShipment{
					Company: item.ShipmentUser.Company,
					No:      item.ShipmentUser.No,
					Remark:  item.ShipmentUser.Remark,
				}
			}

			if item.ShipmentOrganization != nil {
				responses.Data[idx].ShipmentOrganization = &res.ToServiceOfShipment{
					Company: item.ShipmentOrganization.Company,
					No:      item.ShipmentOrganization.No,
				}
			}

			if item.Refund != nil {
				responses.Data[idx].Refund = &res.ToServiceOfRefund{
					ID:         item.Refund.ID,
					No:         item.Refund.No,
					Channel:    item.Refund.Channel,
					Price:      item.Refund.Money,
					Currency:   item.Refund.Currency,
					RefundedAt: item.Refund.RefundedAt.ToDateTimeString(),
				}
			}

			for key, val := range item.Products {

				if val.Detail != nil {
					responses.Data[idx].Details[key] = res.ToOrderOfDetail{
						ID:             val.Detail.ID,
						Name:           val.Detail.Name,
						Picture:        val.Detail.Picture,
						Specifications: val.Detail.Specifications,
						Price:          val.Detail.Price,
						Quantity:       val.Quantity,
						Prices:         val.Detail.Price * val.Quantity,
					}
				}
			}
		}
	}

	http.Success(ctx, responses)
}

func ToServiceOfLogs(c context.Context, ctx *app.RequestContext) {

	var request req.ToServiceOfLogs

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var logs []model.ShpLog

	facades.Gorm.
		Scopes(scope.Platform(ctx)).
		WithContext(c).
		Order("`created_at` desc").
		Where("`service_id` = ?", request.ID).
		Find(&logs)

	responses := make([]res.ToServiceOfLogs, len(logs))

	for idx, item := range logs {
		responses[idx] = res.ToServiceOfLogs{
			ID:        item.ID,
			Action:    item.Action,
			Content:   item.Content,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}

func DoServiceOfHandle(c context.Context, ctx *app.RequestContext) {

	var request req.DoServiceOfHandle

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).Preload("Products").First(&service, "id=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该售后记录")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fs.Error)
		return
	}

	if service.Status != model.ShpServiceOfStatusPending {
		http.Fail(ctx, "该售后已被处理")
		return
	}

	tx := facades.Gorm.Begin()

	log := model.ShpLog{
		Platform:       service.Platform,
		CliqueID:       service.CliqueID,
		OrganizationID: service.OrganizationID,
		UserID:         service.UserID,
		OrderID:        service.OrderID,
		DetailID:       service.DetailID,
		ServiceID:      &service.ID,
		Action:         "service_" + request.Result,
		Content:        lo.If(request.Result == model.ShpServiceOfResultAgree, "After-Sales Result: Agree").Else("After-Sales Result: Reject") + lo.If(request.Remark != "", "; Reason Remarks: "+request.Remark).Else(""),
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后订单处理失败：%v", result.Error)
		return
	}

	updates := map[string]any{
		"result": request.Result,
		"status": model.ShpServiceOfStatusUser,
	}

	if request.Result == model.ShpServiceOfResultRefuse {
		updates["status"] = model.ShpServiceOfStatusClosed
	}

	if service.Type == model.ShpServiceOfTypeUnReceipt && request.Result == model.ShpServiceOfResultAgree { // 未发货的订单，售后同意后直接退款

		updates["status"] = model.ShpServiceOfStatusFinish

		_ = shop.PublishOrderRefund(queue.ShopOrderRefundMessage{
			ID:      service.OrderID,
			Detail:  service.DetailID,
			Service: &service.ID,
			Refund:  service.Refunds,
			Reason:  "User did not receive the goods; after-sales refund.",
			Details: lo.Map(service.Products, func(item model.ShpServiceDetail, index int) queue.ShopOrderRefundOfDetailMessage {
				return queue.ShopOrderRefundOfDetailMessage{
					ID:       item.DetailID,
					Quantity: item.Quantity,
					Refund:   item.Refund,
				}
			}),
		})
	}

	if result := tx.Model(&model.ShpService{}).Where("`id`=?", service.ID).Updates(updates); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后订单处理失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoServiceOfShipment(c context.Context, ctx *app.RequestContext) {

	var request req.DoServiceOfShipment

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&service, "`id`=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到售后记录")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "售后查询失败，请稍后重试")
		return
	}

	if service.Status != model.ShpServiceOfStatusOrg {
		http.Fail(ctx, "该售后无需商户发货")
		return
	}

	tx := facades.Gorm.Begin()

	log := model.ShpLog{
		Platform:       service.Platform,
		CliqueID:       service.CliqueID,
		OrganizationID: service.OrganizationID,
		UserID:         service.UserID,
		OrderID:        service.OrderID,
		DetailID:       service.DetailID,
		ServiceID:      &service.ID,
		Action:         "service_organization",
		Content:        fmt.Sprintf("The seller has sent out the exchange shipment, waiting for the user to process. Courier Company：%s;Tracking Number：%s", request.Company, request.No),
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "快递发出失败，请稍后重试")
		return
	}

	updates := model.ShpService{
		Status: model.ShpServiceOfStatusConfirmUser,
		ShipmentOrganization: &model.ShpServiceOfShipment{
			Company: request.Company,
			No:      request.No,
			Remark:  request.Remark,
		},
	}

	if result := tx.Model(&service).Select("Status", "ShipmentOrganization").Updates(&updates); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "快递发出失败，请稍后重试")
		return
	}

	_ = shop.PublishServiceFinish(service.ID)

	tx.Commit()

	http.Success[any](ctx)
}

func DoServiceOfFinish(c context.Context, ctx *app.RequestContext) {

	var request req.DoServiceOfFinish

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).Preload("Products").First(&service, "`id`=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到售后记录")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "售后查询失败，请稍后重试")
		return
	}

	if service.Status != model.ShpServiceOfStatusConfirmOrg {
		http.Fail(ctx, "该售后无需商户确认")
		return
	}

	tx := facades.Gorm.Begin()

	log := model.ShpLog{
		Platform:       service.Platform,
		CliqueID:       service.CliqueID,
		OrganizationID: service.OrganizationID,
		UserID:         service.UserID,
		OrderID:        service.OrderID,
		DetailID:       service.DetailID,
		ServiceID:      &service.ID,
		Action:         "service_finish",
		Content:        "The seller has completed the after-sales order.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后完结失败，请稍后重试")
		return
	}

	if result := tx.Model(&service).Update("status", model.ShpServiceOfStatusFinish); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后完结失败，请稍后重试")
		return
	}

	if service.Type == model.ShpServiceOfTypeRefund {

		_ = shop.PublishOrderRefund(queue.ShopOrderRefundMessage{
			ID:      service.OrderID,
			Detail:  service.DetailID,
			Service: &service.ID,
			Refund:  service.Refunds,
			Reason:  lo.If(service.Reason != "", service.Reason).Else("The user has requested a return and refund."),
			Details: lo.Map(service.Products, func(item model.ShpServiceDetail, index int) queue.ShopOrderRefundOfDetailMessage {
				return queue.ShopOrderRefundOfDetailMessage{
					ID:       item.DetailID,
					Quantity: item.Quantity,
					Refund:   item.Refund,
				}
			}),
		})
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoServiceOfClosed(c context.Context, ctx *app.RequestContext) {

	var request req.DoServiceOfClosed

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&service, "`id`=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到售后记录")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "售后查询失败，请稍后重试")
		return
	}

	if service.Status == model.ShpServiceOfStatusPending || service.Status == model.ShpServiceOfStatusFinish || service.Status == model.ShpServiceOfStatusClosed {
		http.Fail(ctx, "该售后无法关闭")
		return
	}

	tx := facades.Gorm.Begin()

	log := model.ShpLog{
		Platform:       service.Platform,
		CliqueID:       service.CliqueID,
		OrganizationID: service.OrganizationID,
		UserID:         service.UserID,
		OrderID:        service.OrderID,
		DetailID:       service.DetailID,
		ServiceID:      &service.ID,
		Action:         "service_closed",
		Content:        "The seller has closed the after-sales order.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后关闭失败，请稍后重试")
		return
	}

	if result := tx.Model(&service).Update("status", model.ShpServiceOfStatusClosed); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "售后关闭失败，请稍后重试")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}
