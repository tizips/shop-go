package org

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/hr/platform/org"
	res "project.io/shop/admin/http/response/hr/platform/org"
	"project.io/shop/model"
)

func ToOrganizationOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToOrganizationOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToOrganizationOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Where("`parent_id` is null")

	if request.Platform > 0 {
		tx = tx.Where("`platform` = ?", request.Platform)
	}

	if request.Keyword != "" {
		tx = tx.Where("name like ?", "%"+request.Keyword+"%")
	}

	tx.Model(&model.HROrganization{}).Count(&responses.Total)

	if responses.Total > 0 {

		var organizations []model.HROrganization

		tx.
			Order("`created_at` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&organizations)

		responses.Data = make([]res.ToOrganizationOfPaginate, len(organizations))

		for index, item := range organizations {
			responses.Data[index] = res.ToOrganizationOfPaginate{
				ID:         item.ID,
				Platform:   item.Platform,
				Name:       item.Name,
				ValidStart: item.ValidStart.ToDateString(),
				ValidEnd:   item.ValidEnd.ToDateString(),
				IsEnable:   item.IsEnable,
				CreatedAt:  item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func DoOrganizationOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganizationOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id` is null", request.ID)

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

	organization := model.HROrganization{
		ID:          facades.Snowflake.Generate().String(),
		Platform:    request.Platform,
		Name:        request.Name,
		ValidStart:  start,
		ValidEnd:    end,
		User:        request.User,
		Telephone:   request.Telephone,
		Description: request.Description,
		IsEnable:    request.IsEnable,
	}

	if request.Platform == authConstants.CodeOfClique {
		organization.CliqueID = &organization.ID
	}

	if result := facades.Gorm.Create(&organization); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

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

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id` is null", request.ID)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	organization.Name = request.Name
	organization.ValidStart = start
	organization.ValidEnd = end
	organization.User = request.User
	organization.Telephone = request.Telephone
	organization.Description = request.Description
	organization.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&organization); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
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

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id` is null", request.ID)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	if do := facades.Gorm.Delete(&organization, "id=?", request.ID); do.Error != nil {
		http.Fail(ctx, "删除失败：%v", do.Error)
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

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id` is null", request.ID)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	if result := facades.Gorm.Model(organization).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
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

	fo := facades.Gorm.First(&organization, "`id`=? and `parent_id` is null", request.ID)

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

	if err := auth.DoTemporary(c, ctx, organization.Platform, organization.ID, organization.Name, organization.CliqueID); err != nil {
		http.Fail(ctx, "商户切换失败：%v", err)
		return
	}

	http.Success[any](ctx)
}
