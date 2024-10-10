package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/basic"
	res "project.io/shop/admin/http/response/shop/store/basic"
	"project.io/shop/model"
)

func DoBannerOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	banner := model.ShpBanner{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Picture:        request.Picture,
		Name:           request.Name,
		Description:    request.Description,
		Button:         request.Button,
		Target:         request.Target,
		URL:            request.URL,
		Order:          request.Order.Order,
		IsEnable:       request.IsEnable,
	}

	if result := facades.Gorm.Create(&banner); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.ShpBanner

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	banner.Name = request.Name
	banner.Description = request.Description
	banner.Button = request.Button
	banner.Picture = request.Picture
	banner.Target = request.Target
	banner.URL = request.URL
	banner.Order = request.Order.Order
	banner.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&banner); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.ShpBanner

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&banner); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.ShpBanner

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&banner).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToBannerOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToBannerOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToBannerOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	tx.Model(&model.ShpBanner{}).Count(&responses.Total)

	if responses.Total > 0 {

		var banners []model.ShpBanner

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&banners)

		responses.Data = make([]res.ToBannerOfPaginate, len(banners))

		for idx, item := range banners {
			responses.Data[idx] = res.ToBannerOfPaginate{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				Button:      item.Button,
				Picture:     item.Picture,
				Target:      item.Target,
				URL:         item.URL,
				Order:       item.Order,
				IsEnable:    item.IsEnable,
				CreatedAt:   item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
