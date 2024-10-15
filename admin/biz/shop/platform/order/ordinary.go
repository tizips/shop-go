package order

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/platform/order"
	res "project.io/shop/admin/http/response/shop/platform/order"
	"project.io/shop/model"
)

func ToOrdinaryOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrdinaryOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToOrdinaryOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c)

	if request.ID != "" {
		tx = tx.Where("`id`=?", request.ID)
	}

	if request.Status != "" {
		tx = tx.Where("`status`=?", request.Status)
	} else {
		tx = tx.Where("`status` NOT IN (?)", []string{model.ShpOrderOfStatusClosed})
	}

	tx.Model(&model.ShpOrder{}).Count(&responses.Total)

	if responses.Total > 0 {

		var orders []model.ShpOrder

		tx.
			Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Preload("Payment", func(t *gorm.DB) *gorm.DB { return t.Where("`is_confirmed`=?", util.Yes) }).
			Preload("Shipment").
			Preload("Details").
			Preload("Services", func(t *gorm.DB) *gorm.DB {
				return t.
					Preload("Products").
					Where("`status` not IN (?)", []string{model.ShpServiceOfStatusFinish, model.ShpServiceOfStatusClosed})
			}).
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` desc").
			Find(&orders)

		responses.Data = make([]res.ToOrdinaryOfPaginate, len(orders))

		for idx, item := range orders {

			responses.Data[idx] = res.ToOrdinaryOfPaginate{
				ID:          item.ID,
				TotalPrice:  item.CostPrices,
				Prices:      item.CostPrices,
				CouponPrice: item.CouponPrice,
				Details: lo.Map(item.Details, func(val model.ShpOrderDetail, index int) res.ToOrderOfDetail {
					return res.ToOrderOfDetail{
						ID:             val.ID,
						Name:           val.Name,
						Picture:        val.Picture,
						Specifications: val.Specifications,
						Price:          val.Price,
						Quantity:       val.Quantity,
						TotalPrice:     val.TotalPrice,
						Prices:         val.Prices,
						Refund:         val.Refund,
						Returned:       val.Returned,
					}
				}),
				IsPaid:    item.IsPaid,
				IsInvoice: item.IsInvoice,
				Status:    item.Status,
				Refund:    item.Refund,
				CreatedAt: item.CreatedAt.String(),
			}

			if item.Organization != nil {
				responses.Data[idx].Organization = item.Organization.Name
			}

			if item.Payment != nil {
				responses.Data[idx].Payment = &res.ToOrderOfPayment{
					ID:       item.Payment.ID,
					No:       *item.Payment.No,
					Channel:  item.Payment.Channel,
					Money:    item.Payment.Money,
					Currency: item.Payment.Currency,
					PaidAt:   item.Payment.PaidAt.ToDateTimeString(),
				}
			}

			if item.Shipment != nil {
				responses.Data[idx].Shipping = item.Shipment.Company
			}

			if len(item.Services) > 0 {

				for key, val := range responses.Data[idx].Details {

					mark := false

					for _, value := range item.Services {

						for _, v := range value.Products {

							if v.DetailID == val.ID {

								mark = true

								responses.Data[idx].Details[key].Service = &res.ToOrderOfService{
									ID:       val.ID,
									Type:     value.Type,
									Quantity: v.Quantity,
									Status:   value.Status,
								}

								break
							}
						}

						if mark {
							break
						}
					}
				}
			}
		}
	}

	http.Success(ctx, responses)
}

func ToOrdinaryOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrdinaryOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var order model.ShpOrder

	fo := facades.Gorm.
		Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		//Preload("Shipping").
		Preload("Details").
		Preload("Address").
		Preload("Payment", func(t *gorm.DB) *gorm.DB { return t.Where("`is_confirmed`=?", util.Yes) }).
		Preload("Shipment").
		//Preload("Invoice").
		Preload("Services", func(t *gorm.DB) *gorm.DB {
			return t.
				Preload("Products").
				Where("`status` not IN (?)", []string{model.ShpServiceOfStatusFinish, model.ShpServiceOfStatusClosed})
		}).
		First(&order, "`id`=?", request.ID)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该订单")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fo.Error)
		return
	}

	responses := res.ToOrdinaryOfInformation{
		ID: order.ID,
		Details: lo.Map(order.Details, func(val model.ShpOrderDetail, index int) res.ToOrderOfDetail {
			return res.ToOrderOfDetail{
				ID:             val.ID,
				Name:           val.Name,
				Picture:        val.Picture,
				Specifications: val.Specifications,
				Price:          val.Price,
				Quantity:       val.Quantity,
				TotalPrice:     val.TotalPrice,
				Prices:         val.Prices,
				Refund:         val.Refund,
				Returned:       val.Returned,
			}
		}),
		Address:      res.ToOrdinaryOfAddress{},
		Invoice:      nil,
		Shipment:     nil,
		CostShipping: order.CostShipping,
		TotalPrice:   order.TotalPrice,
		CouponPrice:  order.CouponPrice,
		Prices:       order.Prices,
		Refund:       order.Refund,
		Status:       order.Status,
		Remark:       order.Remark,
		IsInvoice:    order.IsInvoice,
		CreateAt:     order.CreatedAt.ToDateTimeString(),
	}

	if order.Organization != nil {
		responses.Organization = order.Organization.Name
	}

	if order.Shipment != nil {
		responses.Shipping = order.Shipment.Company
	}

	//if order.Invoice != nil {
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
	//	}
	//}

	if order.Shipment != nil {
		responses.Shipment = &res.ToOrderOfShipment{
			Company:   order.Shipment.Company,
			No:        order.Shipment.No,
			Remark:    order.Shipment.Remark,
			CreatedAt: order.Shipment.CreatedAt.ToDateTimeString(),
		}
	}

	if order.Address != nil {
		responses.Address = res.ToOrdinaryOfAddress{
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

	if order.Payment != nil {
		responses.Payment = &res.ToOrderOfPayment{
			ID:       order.Payment.ID,
			No:       *order.Payment.No,
			Channel:  order.Payment.Channel,
			Money:    order.Payment.Money,
			Currency: order.Payment.Currency,
			PaidAt:   order.Payment.PaidAt.ToDateTimeString(),
		}
	}

	if len(order.Services) > 0 {

		for key, val := range responses.Details {

			mark := false

			for _, value := range order.Services {

				for _, v := range value.Products {

					if v.DetailID == val.ID {

						mark = true

						responses.Details[key].Service = &res.ToOrderOfService{
							ID:       val.ID,
							Type:     value.Type,
							Quantity: v.Quantity,
							Status:   value.Status,
						}

						break
					}
				}

				if mark {
					break
				}
			}
		}
	}

	http.Success(ctx, responses)
}

func ToOrdinaryOfLogs(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrdinaryOfLogs

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var logs []model.ShpLog

	facades.Gorm.
		WithContext(c).
		Order("`created_at` desc").
		Where("`order_id` = ?", request.ID).
		Find(&logs)

	responses := make([]res.ToOrderOfLogs, len(logs))

	for idx, item := range logs {
		responses[idx] = res.ToOrderOfLogs{
			ID:        item.ID,
			Action:    item.Action,
			Content:   item.Content,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}
