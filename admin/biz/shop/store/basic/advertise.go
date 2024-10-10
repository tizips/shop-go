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
	"github.com/samber/lo"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/basic"
	res "project.io/shop/admin/http/response/shop/store/basic"
	"project.io/shop/model"
)

func ToAdvertiseOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToAdvertiseOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToAdvertiseOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	if request.Pages != "" {
		tx = tx.Where("`page`=?", request.Pages)
	}

	if request.Block != "" {
		tx = tx.Where("`block`=?", request.Block)
	}

	tx.Model(&model.ShpAdvertise{}).Count(&responses.Total)

	if responses.Total > 0 {

		var advertises []model.ShpAdvertise

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` asc").
			Find(&advertises)

		responses.Data = make([]res.ToAdvertiseOfPaginate, len(advertises))

		for idx, item := range advertises {
			responses.Data[idx] = res.ToAdvertiseOfPaginate{
				ID:        item.ID,
				Page:      item.Page,
				Block:     item.Block,
				Title:     item.Title,
				Target:    item.Target,
				URL:       item.URL,
				Thumb:     item.Thumb,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func DoAdvertiseOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoAdvertiseOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if request.Page == model.ShpAdvertiseOfPageHome {

		if !lo.Contains([]string{model.ShpAdvertiseOfBlockNewProduct}, request.Block) {
			http.Fail(ctx, "%s 不在页面 %s 下", request.Block, request.Page)
		}
	}

	advertise := model.ShpAdvertise{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Page:           request.Page,
		Block:          request.Block,
		Title:          request.Title,
		Target:         request.Target,
		URL:            request.URL,
		Thumb:          request.Thumb,
		Order:          request.Order.Order,
		IsEnable:       request.IsEnable,
	}

	if result := facades.Gorm.Create(&advertise); result.Error != nil {
		http.Fail(ctx, "写入失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)

}

func DoAdvertiseOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoAdvertiseOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if request.Page == model.ShpAdvertiseOfPageHome {

		if !lo.Contains([]string{model.ShpAdvertiseOfBlockNewProduct}, request.Block) {
			http.Fail(ctx, "%s 不在页面 %s 下", request.Block, request.Page)
		}
	}

	var advertise model.ShpAdvertise

	fa := facades.Gorm.Scopes(scope.Platform(ctx)).First(&advertise, "`id`=?", request.ID)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该广告信息")
		return
	} else if fa.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fa.Error)
		return
	}

	updates := model.ShpAdvertise{
		Page:     request.Page,
		Block:    request.Block,
		Title:    request.Title,
		Target:   request.Target,
		URL:      request.URL,
		Thumb:    request.Thumb,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Model(&advertise).Select("Page", "Block", "Title", "Target", "URL", "Thumb", "Order", "IsEnable").Updates(&updates); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoAdvertiseOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoAdvertiseOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var advertise model.ShpAdvertise

	fa := facades.Gorm.Scopes(scope.Platform(ctx)).First(&advertise, "`id`=?", request.ID)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fa.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fa.Error)
		return
	}

	if result := facades.Gorm.Delete(&advertise); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoAdvertiseOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoAdvertiseOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var advertise model.ShpAdvertise

	fa := facades.Gorm.Scopes(scope.Platform(ctx)).First(&advertise, "`id`=?", request.ID)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fa.Error != nil {
		http.Fail(ctx, "查找失败：%v", fa.Error)
		return
	}

	if result := facades.Gorm.Model(&advertise).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}
