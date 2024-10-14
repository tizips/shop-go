package shop

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToServiceOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToServiceOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToServiceOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.
		WithContext(c).
		Where("`user_id`=?", auth.ID(ctx))

	tx.Model(&model.ShpService{}).Count(&responses.Total)

	if responses.Total > 0 {

		var services []model.ShpService

		tx.
			Preload("Products.Detail").
			Order("`created_at` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&services)

		responses.Data = make([]res.ToServiceOfPaginate, len(services))

		for idx, item := range services {

			responses.Data[idx] = res.ToServiceOfPaginate{
				ID:        item.ID,
				Type:      item.Type,
				Status:    item.Status,
				Reason:    item.Reason,
				Subtotal:  item.Subtotal,
				Shipping:  item.Shipping,
				Refund:    item.Refunds,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			responses.Data[idx].Details = make([]res.ToServiceOfDetail, len(item.Products))

			for key, val := range item.Products {

				if val.Detail != nil {
					responses.Data[idx].Details[key] = res.ToServiceOfDetail{
						ID:             val.DetailID,
						Name:           val.Detail.Name,
						Picture:        val.Detail.Picture,
						Price:          val.Detail.Price,
						Quantity:       val.Quantity,
						Refund:         val.Detail.Price * val.Quantity,
						Specifications: val.Detail.Specifications,
						CreatedAt:      val.CreatedAt.ToDateTimeString(),
					}
				}
			}
		}
	}

	http.Success(ctx, responses)
}

func ToServiceOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToServiceOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.
		Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Preload("Products.Detail").
		Preload("Logs", func(t *gorm.DB) *gorm.DB { return t.Order("`created_at` desc") }).
		First(&service, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "No after-sales service found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "After-sales query failed. Please try again later.")
		return
	}

	responses := res.ToServiceOfInformation{
		ID:       service.ID,
		Order:    service.OrderID,
		Type:     service.Type,
		Status:   service.Status,
		Reason:   service.Reason,
		Pictures: service.Pictures,
		Details:  make([]res.ToServiceOfDetail, len(service.Products)),
		Subtotal: service.Subtotal,
		Shipping: service.Shipping,
		Refund:   service.Refunds,
		Logs: lo.Map(service.Logs, func(item model.ShpLog, _ int) res.ToServiceOfLog {
			return res.ToServiceOfLog{
				Action:    item.Action,
				Content:   item.Content,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}),
		CreatedAt: service.CreatedAt.ToDateTimeString(),
	}

	if service.Organization != nil {
		responses.Organization = &res.ToServiceOfOrganization{
			ID:   *service.OrganizationID,
			Name: service.Organization.Name,
		}
	}

	for key, val := range service.Products {

		if val.Detail != nil {
			responses.Details[key] = res.ToServiceOfDetail{
				ID:             val.DetailID,
				Name:           val.Detail.Name,
				Picture:        val.Detail.Picture,
				Price:          val.Detail.Price,
				Quantity:       val.Quantity,
				Refund:         val.Detail.Price * val.Quantity,
				Specifications: val.Detail.Specifications,
				CreatedAt:      val.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func DoServiceOfCancel(c context.Context, ctx *app.RequestContext) {

	var request req.DoServiceOfCancel

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var service model.ShpService

	fs := facades.Gorm.
		Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		First(&service, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "No after-sales service found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "After-sales query failed. Please try again later.")
		return
	}

	if service.Status != model.ShpServiceOfStatusPending {
		http.Fail(ctx, "The after-sales service has already been processed and cannot be canceled.")
		return
	}

	tx := facades.Gorm.Begin()

	if result := tx.Delete(&service, "id=?", service.ID); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Failed to cancel after-sales service. Please try again later.")
		return
	}

	log := model.ShpLog{
		Platform:       service.Platform,
		CliqueID:       service.CliqueID,
		OrganizationID: service.OrganizationID,
		UserID:         service.UserID,
		OrderID:        service.OrderID,
		DetailID:       service.DetailID,
		ServiceID:      &service.ID,
		Action:         "service_cancel",
		Content:        "User canceled after-sales request.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Failed to cancel after-sales service. Please try again later.")
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

	fs := facades.Gorm.First(&service, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "No after-sales service found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "After-sales query failed. Please try again later.")
		return
	}

	if service.Status != model.ShpServiceOfStatusUser {
		http.Fail(ctx, "No shipping is required from the user for this after-sales service.")
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
		Action:         "service_user",
		Content:        fmt.Sprintf("The user has sent out the shipment and is waiting for the seller to process it. Courier Company：%v;Tracking Number：%v", request.Company, request.No),
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Shipment dispatch failed. Please try again later.")
		return
	}

	updates := model.ShpService{
		Status: model.ShpServiceOfStatusOrg,
		ShipmentUser: &model.ShpServiceOfShipment{
			Company: request.Company,
			No:      request.No,
			Remark:  request.Remark,
		},
	}

	if service.Type == model.ShpServiceOfTypeRefund {
		updates.Status = model.ShpServiceOfStatusConfirmOrg
	}

	if result := tx.Model(&service).Select("Status", "ShipmentUser").Updates(&updates); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Shipment dispatch failed. Please try again later.")
		return
	}

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

	fs := facades.Gorm.First(&service, "`id`=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "No after-sales service found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "After-sales query failed. Please try again later.")
		return
	}

	if service.Status != model.ShpServiceOfStatusConfirmUser {
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
		Content:        "User has completed the after-sales order.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "After-sales completion failed. Please try again later.")
		return
	}

	if result := tx.Model(&service).Update("status", model.ShpServiceOfStatusFinish); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "After-sales completion failed. Please try again later.")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}
