package org

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/hr/clique/org"
	res "project.io/shop/admin/http/response/hr/clique/org"
	"project.io/shop/model"
)

func ToBrands(c context.Context, ctx *app.RequestContext) {

	var brands []model.HRBrand

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&brands)

	responses := make([]res.ToBrands, len(brands))

	for index, item := range brands {
		responses[index] = res.ToBrands{
			ID:        item.ID,
			Name:      item.Name,
			Logo:      item.Logo,
			Order:     item.Order,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}

func DoBrandOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBrandOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	brand := model.HRBrand{
		Platform:       auth.Platform(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Logo:           request.Logo,
		Order:          request.Order.Order,
	}

	if result := facades.Gorm.Create(&brand); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBrandOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBrandOfUpdate

	if err := ctx.Bind(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var brand model.HRBrand

	fb := facades.Gorm.Scopes(scope.Platform(ctx)).First(&brand, "`id`=?", request.ID)

	if errors.Is(fb.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fb.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fb.Error)
		return
	}

	brand.Logo = request.Logo
	brand.Name = request.Name
	brand.Order = request.Order.Order

	if result := facades.Gorm.Save(&brand); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBrandOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoBrandOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if result := facades.Gorm.Scopes(scope.Platform(ctx)).Delete(&model.HRBrand{}, "`id`=?", request.ID); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToBrandOfOpening(c context.Context, ctx *app.RequestContext) {

	var brands []model.HRBrand

	facades.Gorm.Order("`order` asc, `id` asc").Find(&brands, "`platform`=? and `organization_id`=?", authConstants.CodeOfClique, auth.Clique(ctx))

	responses := make([]response.Opening[uint], len(brands))

	for index, item := range brands {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}
