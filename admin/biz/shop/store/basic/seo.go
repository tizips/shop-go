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

func DoSEOOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSEOOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64 = 0

	facades.Gorm.Model(&model.ShpSEO{}).Where("`channel`=? and `channel_id`=?", model.ShpSEOForChannelOfCategory, request.Code).Count(&total)

	if total > 0 {
		http.Fail(ctx, "该 Code 已被使用")
		return
	}

	seo := model.ShpSEO{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Channel:        model.ShpSEOForChannelOfCategory,
		ChannelID:      request.Code,
		Title:          request.Title,
		Keyword:        request.Keyword,
		Description:    request.Description,
	}

	if result := facades.Gorm.Create(&seo); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSEOOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSEOOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var seo model.ShpSEO

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&seo, "channel=? and `id`=?", model.ShpSEOForChannelOfCategory, request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fs.Error)
		return
	}

	seo.Title = request.Title
	seo.Keyword = request.Keyword
	seo.Description = request.Description

	if result := facades.Gorm.Save(&seo); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSEOOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoSEOOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var seo model.ShpSEO

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&seo, "`channel`=? and `id`=?", model.ShpSEOForChannelOfCategory, request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fs.Error)
		return
	}

	if result := facades.Gorm.Delete(&seo); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToSEOOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToSEOOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToSEOOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx)).Where("`channel`=?", model.ShpSEOForChannelOfCategory)

	tx.Model(&model.ShpSEO{}).Count(&responses.Total)

	if responses.Total > 0 {

		var seos []model.ShpSEO

		tx.
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&seos)

		responses.Data = make([]res.ToSEOOfPaginate, len(seos))

		for idx, item := range seos {
			responses.Data[idx] = res.ToSEOOfPaginate{
				ID:          item.ID,
				Channel:     item.Channel,
				Code:        item.ChannelID,
				Title:       item.Title,
				Keyword:     item.Keyword,
				Description: item.Description,
				CreatedAt:   item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
