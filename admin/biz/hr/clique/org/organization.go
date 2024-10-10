package org

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/hr/clique/org"
	res "project.io/shop/admin/http/response/hr/clique/org"
	"project.io/shop/model"
)

func ToOrganizationOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrganizationOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToOrganizationOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.WithContext(c).Where("`clique_id`=?", auth.Clique(ctx))

	if request.Parent == "" {
		tx = tx.Where("`parent_id`=?", auth.Organization(ctx))
	} else {
		tx = tx.Where("`parent_id`=?", request.Parent)
	}

	if request.Keyword != "" {
		tx = tx.Where("`name` like ?", "%"+request.Keyword+"%")
	}

	if request.Platform > 0 {
		tx = tx.Where("`platform` = ?", request.Platform)
	}

	tx.Model(&model.HROrganization{}).Count(&responses.Total)

	if responses.Total > 0 {

		var organizations []model.HROrganization

		tx.
			Preload("Brand", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Order("`created_at` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&organizations)

		responses.Data = make([]res.ToOrganizationOfPaginate, len(organizations))

		for index, item := range organizations {

			responses.Data[index] = res.ToOrganizationOfPaginate{
				ID:         item.ID,
				Name:       item.Name,
				Platform:   item.Platform,
				ValidStart: item.ValidStart.ToDateString(),
				ValidEnd:   item.ValidEnd.ToDateString(),
				IsEnable:   item.IsEnable,
				CreatedAt:  item.CreatedAt.ToDateTimeString(),
			}

			if item.Brand != nil {
				responses.Data[index].Brand = item.Brand.Name
			}
		}
	}

	http.Success(ctx, responses)
}

func DoOrganizationOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	start := carbon.ParseByFormat(request.ValidStart, "Y-m-d")
	end := carbon.ParseByFormat(request.ValidEnd, "Y-m-d")

	if start.Gt(end) {
		http.Fail(ctx, "「有效期：开始」不能大于「有效期：结束」")
		return
	}

	if request.Brand > 0 {

		var count int64 = 0

		facades.Gorm.Model(&model.HRBrand{}).Where("id=?", request.Brand).Count(&count)

		if count <= 0 {
			http.NotFound(ctx, "未找到该品牌")
			return
		}
	}

	tx := facades.Gorm.Begin()

	organization := model.HROrganization{
		ID:          facades.Snowflake.Generate().String(),
		Platform:    request.Platform,
		CliqueID:    auth.Clique(ctx),
		Name:        request.Name,
		BrandID:     request.Brand,
		ParentID:    auth.Organization(ctx),
		ValidStart:  start,
		ValidEnd:    end,
		User:        request.User,
		Telephone:   request.Telephone,
		Description: request.Description,
		IsEnable:    request.IsEnable,
	}

	if result := tx.Create(&organization); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	children := make([]model.HRChildren, 0)

	if request.Parent == "" {
		children = append(children, model.HRChildren{
			OrganizationID: *auth.Organization(ctx),
			ChildID:        organization.ID,
		})
	}

	if len(children) > 0 {
		if result := tx.Create(&children); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "创建失败：%v", result.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoOrganizationOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	start := carbon.ParseByFormat(request.ValidStart, "Y-m-d")
	end := carbon.ParseByFormat(request.ValidEnd, "Y-m-d")

	if start.Gt(end) {
		http.Fail(ctx, "「有效期：开始」不能大于「有效期：结束」")
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `clique_id`=?", request.ID, auth.Clique(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fo.Error)
		return
	}

	if organization.BrandID > 0 && request.Brand != organization.BrandID {

		var count int64 = 0

		facades.Gorm.Model(&model.HRBrand{}).Where("`id`=?", request.Brand).Count(&count)

		if count <= 0 {
			http.NotFound(ctx, "未找到该品牌")
			return
		}
	}

	organization.Name = request.Name
	organization.ValidStart = start
	organization.ValidEnd = end
	organization.User = request.User
	organization.Telephone = request.Telephone
	organization.IsEnable = request.IsEnable

	if organization.BrandID > 0 {
		organization.BrandID = request.Brand
	}

	if result := facades.Gorm.Save(&organization); result.Error != nil {
		http.Fail(ctx, "修改失败")
		return
	}

	http.Success[any](ctx)
}

func DoOrganizationOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `clique_id`=?", request.ID, auth.Clique(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fo.Error)
		return
	}

	if result := facades.Gorm.Delete(&organization, "`id`=?", organization.ID); result.Error != nil {
		http.Fail(ctx, "删除失败")
		return
	}

	http.Success[any](ctx)

}

func DoOrganizationOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `clique_id`=?", request.ID, auth.Clique(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fo.Error)
		return
	}

	if result := facades.Gorm.Model(&organization).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "修改失败")
		return
	}

	http.Success[any](ctx)
}

func DoOrganizationOfEnter(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfEnter

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id`=?", request.ID, auth.Organization(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	if organization.IsEnable != util.Yes {
		http.Fail(ctx, "该商户暂未启用")
		return
	}

	temporary, _ := auth.Temporary(c, ctx)

	if err := auth.DoTemporary(c, ctx, organization.Platform, organization.ID, organization.Name, organization.CliqueID, temporary); err != nil {
		http.Fail(ctx, "商户切换失败：%v", err)
		return
	}

	http.Success[any](ctx)
}

func ToOrganizationOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrganizationOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `clique_id`=?", request.ID, auth.Clique(ctx))

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	responses := res.ToOrganizationOfInformation{
		ID:          organization.ID,
		Platform:    organization.Platform,
		Brand:       organization.BrandID,
		Name:        organization.Name,
		ValidStart:  organization.ValidStart.ToDateString(),
		ValidEnd:    organization.ValidEnd.ToDateString(),
		Description: organization.Description,
		User:        organization.User,
		Telephone:   organization.Telephone,
		IsEnable:    organization.IsEnable,
		CreatedAt:   organization.CreatedAt.ToDateTimeString(),
	}

	http.Success(ctx, responses)
}
