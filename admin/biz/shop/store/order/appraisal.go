package order

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/order"
	res "project.io/shop/admin/http/response/shop/store/order"
	"project.io/shop/model"
)

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

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	if len(request.Stars) > 0 {

		if request.Type == "product" {
			tx.Where("star_product IN (?)", request.Stars)
		} else if request.Type == "shipment" {
			tx.Where("star_shipment IN (?)", request.Stars)
		} else {
			tx.Where("`star_product` IN (?) or `star_shipment` IN (?)", request.Stars, request.Stars)
		}
	}

	tx.Model(&model.ShpAppraisal{}).Count(&responses.Total)

	if responses.Total > 0 {

		var appraisals []model.ShpAppraisal

		tx.
			Preload("User", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Order("`created_at` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&appraisals)

		responses.Data = make([]res.ToAppraisalOfPaginate, len(appraisals))

		for idx, item := range appraisals {

			responses.Data[idx] = res.ToAppraisalOfPaginate{
				ID:           item.ID,
				Order:        item.OrderID,
				StarProduct:  item.StarProduct,
				StarShipment: item.StarShipment,
				Remark:       item.Remark,
				Pictures:     item.Pictures,
				CreatedAt:    item.CreatedAt.ToDateTimeString(),
			}

			if item.User != nil {
				responses.Data[idx].FirstName = item.User.FirstName
				responses.Data[idx].LastName = item.User.LastName
				responses.Data[idx].Email = item.User.Email
			}
		}
	}

	http.Success(ctx, responses)
}
