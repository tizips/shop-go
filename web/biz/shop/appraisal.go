package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func DoAppraisalOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoAppraisalOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var order model.ShpOrder

	fo := facades.Gorm.First(&order, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Order not found.")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "Order query failed. Please try again later.")
		return
	}

	if order.IsAppraisal == util.Yes {
		http.Fail(ctx, "This order has already been reviewed and cannot be reviewed again.")
		return
	} else if order.Status != model.ShpOrderOfStatusReceived {
		http.Fail(ctx, "This order has not yet been received and cannot be reviewed.")
		return
	}

	tx := facades.Gorm.Begin()

	appraisal := model.ShpAppraisal{
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		UserID:         order.UserID,
		OrderID:        order.ID,
		StarProduct:    request.StarProduct,
		StarShipment:   request.StarShipment,
		Remark:         request.Remark,
		Pictures:       request.Pictures,
	}

	if result := tx.Create(&appraisal); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order review failed. Please try again later.")
		return
	}

	log := model.ShpLog{
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		UserID:         order.UserID,
		OrderID:        order.ID,
		Action:         "appraisal",
		Content:        "User submitted a review.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order review failed. Please try again later.")
		return
	}

	if result := tx.Model(&model.ShpOrder{}).Where("`id`=? and `user_id`=? and is_appraisal=?", appraisal.OrderID, appraisal.UserID, util.No).Update("is_appraisal", util.Yes); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order review failed. Please try again later.")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoAppraisalOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoAppraisalOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var appraisal model.ShpAppraisal

	fa := facades.Gorm.First(&appraisal, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该评价")
		return
	} else if fa.Error != nil {
		http.Fail(ctx, "评价查询失败，请稍后重试")
		return
	}

	tx := facades.Gorm.Begin()

	if result := tx.Delete(&appraisal, "`id`=?", appraisal.ID); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "订单评价失败，请稍后重试")
		return
	}

	log := model.ShpLog{
		Platform:       appraisal.Platform,
		CliqueID:       appraisal.CliqueID,
		OrganizationID: appraisal.OrganizationID,
		UserID:         appraisal.UserID,
		OrderID:        appraisal.OrderID,
		Action:         "appraisal",
		Content:        "用户删除评价",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "订单评价失败，请稍后重试")
		return
	}

	if result := tx.Model(&model.ShpOrder{}).Where("`id`=? and `user_id`=? and is_appraisal=?", appraisal.OrderID, appraisal.UserID, util.Yes).Update("is_appraisal", util.No); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "订单评价失败，请稍后重试")
		return
	}

	tx.Commit()
}

func ToAppraisalOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToAppraisalOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToAppraisalOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.WithContext(c).Where("`user_id`=?", auth.ID(ctx))

	tx.Model(&model.ShpAppraisal{}).Count(&responses.Total)

	if responses.Total > 0 {

		var appraisals []model.ShpAppraisal

		tx.
			Preload("Organization", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Preload("Details").
			Order("`created_at` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&appraisals)

		responses.Data = make([]res.ToAppraisalOfPaginate, len(appraisals))

		for idx, item := range appraisals {

			responses.Data[idx] = res.ToAppraisalOfPaginate{
				ID: item.ID,
				Organization: res.ToOrderOfOrganization{
					ID: *item.OrganizationID,
				},
				Details: lo.Map(item.Details, func(val model.ShpOrderDetail, index int) res.ToOrderOfDetail {
					return res.ToOrderOfDetail{
						ID:             val.ID,
						Name:           val.Name,
						Picture:        val.Picture,
						Price:          val.Price,
						Quantity:       val.Quantity,
						Specifications: val.Specifications,
					}
				}),
				StarProduct:  item.StarProduct,
				StarShipment: item.StarShipment,
				Pictures:     item.Pictures,
				Remark:       item.Remark,
				CreatedAt:    item.CreatedAt.ToDateTimeString(),
			}

			if item.Organization != nil {
				responses.Data[idx].Organization.Name = item.Organization.Name
			}
		}
	}

	http.Success(ctx, responses)
}
