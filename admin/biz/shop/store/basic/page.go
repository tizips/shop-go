package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	req "project.io/shop/admin/http/request/shop/store/basic"
	res "project.io/shop/admin/http/response/shop/store/basic"
	"project.io/shop/model"
	"strconv"
)

func DoPageOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64 = 0

	facades.Gorm.Model(&model.ShpPage{}).Scopes(scope.Platform(ctx)).Where("`code`=?", request.Code).Count(&total)

	if total > 0 {
		http.Fail(ctx, "该 Code 已被使用")
		return
	}

	if request.IsSystem == util.Yes {

		if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

			permissions := []any{auth.NameOfUser(auth.ID(ctx)), "shop.basic.page.system"}

			platform := auth.Platform(ctx)

			if platform > 0 {
				permissions = append(permissions, auth.SPlatform(ctx))
			}

			if ok, _ = facades.Casbin.Enforce(permissions...); !ok {
				http.Fail(ctx, "没有权限写入系统级页面")
				return
			}
		}
	}

	tx := facades.Gorm.Begin()

	page := model.ShpPage{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Code:           request.Code,
		IsSystem:       request.IsSystem,
		Content:        request.Content,
	}

	if result := tx.Create(&page); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	seo := model.ShpSEO{
		Channel:     model.ShpSEOForChannelOfPage,
		ChannelID:   strconv.Itoa(int(page.ID)),
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

func DoPageOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShpPage

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfPage) }).
		First(&page, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	if request.IsSystem > 0 && request.IsSystem != page.IsSystem {

		if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

			permissions := []any{auth.NameOfUser(auth.ID(ctx)), "shop.basic.page.system"}

			platform := auth.Platform(ctx)

			if platform > 0 {
				permissions = append(permissions, auth.SPlatform(ctx))
			}

			if ok, _ = facades.Casbin.Enforce(permissions...); !ok {
				http.Fail(ctx, "没有权限更改系统级页面")
				return
			}
		}
	}

	tx := facades.Gorm.Begin()

	page.Name = request.Name
	page.Content = request.Content

	if request.IsSystem > 0 {
		page.IsSystem = request.IsSystem
	}

	if result := tx.Omit(clause.Associations).Save(&page); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	seo := model.ShpSEO{
		Channel:     model.ShpSEOForChannelOfPage,
		ChannelID:   strconv.Itoa(int(page.ID)),
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
	}

	if page.SEO != nil {
		seo = *page.SEO

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

func DoPageOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShpPage

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&page, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if page.IsSystem == util.Yes {

		if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

			permissions := []any{auth.NameOfUser(auth.ID(ctx)), "shop.basic.page.system"}

			platform := auth.Platform(ctx)

			if platform > 0 {
				permissions = append(permissions, auth.SPlatform(ctx))
			}

			if ok, _ = facades.Casbin.Enforce(permissions...); !ok {
				http.Fail(ctx, "系统内置页面，无法删除")
				return
			}
		}
	}

	if result := facades.Gorm.Delete(&page); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToPageOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToPageOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	if request.IsSystem > 0 {
		tx = tx.Where("`is_system` = ?", request.IsSystem)
	}

	tx.Model(&model.ShpPage{}).Count(&responses.Total)

	if responses.Total > 0 {

		var pages []model.ShpPage

		tx.
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&pages)

		responses.Data = make([]res.ToPageOfPaginate, len(pages))

		for idx, item := range pages {
			responses.Data[idx] = res.ToPageOfPaginate{
				ID:        item.ID,
				Code:      item.Code,
				Name:      item.Name,
				IsSystem:  item.IsSystem,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToPageOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShpPage

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfPage) }).
		First(&page, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	responses := res.ToPageOfInformation{
		ID:        page.ID,
		Name:      page.Name,
		Code:      page.Code,
		Content:   page.Content,
		IsSystem:  page.IsSystem,
		CreatedAt: page.CreatedAt.ToDateTimeString(),
	}

	if page.SEO != nil {
		responses.Title = page.SEO.Title
		responses.Keyword = page.SEO.Keyword
		responses.Description = page.SEO.Description
	}

	http.Success(ctx, responses)
}
