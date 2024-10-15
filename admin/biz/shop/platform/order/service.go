package order

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/platform/order"
	res "project.io/shop/admin/http/response/shop/platform/order"
	"project.io/shop/model"
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
			Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
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

			if item.Organization != nil {
				responses.Data[idx].Organization = item.Organization.Name
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
