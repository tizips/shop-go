package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	req "project.io/shop/admin/http/request/shop/store/basic"
	res "project.io/shop/admin/http/response/shop/store/basic"
	"project.io/shop/model"
)

func DoBlogOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBlogOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	tx := facades.Gorm.Begin()

	blog := model.ShpBlog{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Thumb:          request.Thumb,
		Summary:        request.Summary,
		IsTop:          request.IsTop,
		Content:        request.Content,
		PostedAt:       carbon.ParseByFormat(request.PostedAt, "Y-m-d"),
	}

	if result := tx.Create(&blog); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	seo := model.ShpSEO{
		Channel:     model.ShpSEOForChannelOfBlog,
		ChannelID:   blog.ID,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
	}

	if result := tx.Create(&seo); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoBlogOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBlogOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var blog model.ShpBlog

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfBlog) }).
		First(&blog, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	tx := facades.Gorm.Begin()

	blog.Name = request.Name
	blog.Thumb = request.Thumb
	blog.Summary = request.Summary
	blog.IsTop = request.IsTop
	blog.Content = request.Content
	blog.PostedAt = carbon.ParseByFormat(request.PostedAt, "Y-m-d")

	if result := tx.Omit(clause.Associations).Save(&blog); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	seo := model.ShpSEO{
		Channel:     model.ShpSEOForChannelOfBlog,
		ChannelID:   blog.ID,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
	}

	if blog.SEO != nil {
		seo = *blog.SEO

		seo.Title = request.Title
		seo.Keyword = request.Keyword
		seo.Description = request.Description
	}

	if result := tx.Save(&seo); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoBlogOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoBlogOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var blog model.ShpBlog

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&blog, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&blog); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToBlogOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToBlogOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToBlogOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	if request.IsTop > 0 {
		tx = tx.Where("`is_top` = ?", request.IsTop)
	}

	tx.Model(&model.ShpBlog{}).Count(&responses.Total)

	if responses.Total > 0 {

		var blogs []model.ShpBlog

		tx.
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&blogs)

		responses.Data = make([]res.ToBlogOfPaginate, len(blogs))

		for idx, item := range blogs {
			responses.Data[idx] = res.ToBlogOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Thumb:     item.Thumb,
				IsTop:     item.IsTop,
				PostedAt:  item.PostedAt.ToDateString(),
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToBlogOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToBlogOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var blog model.ShpBlog

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfBlog) }).
		First(&blog, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	responses := res.ToBlogOfInformation{
		ID:        blog.ID,
		Name:      blog.Name,
		Thumb:     blog.Thumb,
		Summary:   blog.Summary,
		IsTop:     blog.IsTop,
		Content:   blog.Content,
		PostedAt:  blog.PostedAt.ToDateString(),
		CreatedAt: blog.CreatedAt.ToDateTimeString(),
	}

	if blog.SEO != nil {
		responses.Title = blog.SEO.Title
		responses.Keyword = blog.SEO.Keyword
		responses.Description = blog.SEO.Description
	}

	http.Success(ctx, responses)
}
