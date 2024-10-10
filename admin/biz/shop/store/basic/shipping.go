package basic

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/basic"
	res "project.io/shop/admin/http/response/shop/store/basic"
	"project.io/shop/model"
)

func DoShippingOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoShippingOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	shipping := model.ShpShipping{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Money:          request.Money,
		Query:          request.Query,
		Order:          request.Order.Order,
		IsEnable:       request.IsEnable,
	}

	if result := facades.Gorm.Create(&shipping); result.Error != nil {
		http.Fail(ctx, fmt.Sprintf("创建失败：%v", result.Error))
		return
	}

	http.Success[any](ctx)
}

func DoShippingOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoShippingOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var shipping model.ShpShipping

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&shipping, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	shipping.Name = request.Name
	shipping.Money = request.Money
	shipping.Query = request.Query
	shipping.Order = request.Order.Order
	shipping.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&shipping); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoShippingOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoShippingOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var shipping model.ShpShipping

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&shipping, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&shipping); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoShippingOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoShippingOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var shipping model.ShpShipping

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&shipping, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&shipping).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToShippingOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToShippingOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToShippingOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	tx.Model(&model.ShpShipping{}).Count(&responses.Total)

	if responses.Total > 0 {

		var shippings []model.ShpShipping

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` asc").
			Find(&shippings)

		responses.Data = make([]res.ToShippingOfPaginate, len(shippings))

		for idx, item := range shippings {
			responses.Data[idx] = res.ToShippingOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Money:     item.Money,
				Query:     item.Query,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToShippingOfOpening(c context.Context, ctx *app.RequestContext) {

	var shippings []model.ShpShipping

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `id` asc").Find(&shippings, "`is_enable`=?", util.Yes)

	responses := make([]res.ToShippingOfOpening, len(shippings))

	for idx, item := range shippings {
		responses[idx] = res.ToShippingOfOpening{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}
