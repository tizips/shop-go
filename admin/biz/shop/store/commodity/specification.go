package commodity

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
	req "project.io/shop/admin/http/request/shop/store/commodity"
	res "project.io/shop/admin/http/response/shop/store/commodity"
	"project.io/shop/model"
)

func DoSpecificationOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSpecificationOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	template := model.ShpTemplateSpecification{
		Name:           request.Name,
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Label:          request.Label,
		Options:        request.Options,
		IsEnable:       request.IsEnable,
	}

	if ct := facades.Gorm.Create(&template); ct.Error != nil {
		http.Fail(ctx, "创建失败：%v", ct.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSpecificationOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSpecificationOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var template model.ShpTemplateSpecification

	ft := facades.Gorm.Scopes(scope.Platform(ctx)).First(&template, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该信息")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", ft.Error)
		return
	}

	template.Name = request.Name
	template.Label = request.Label
	template.Options = request.Options
	template.IsEnable = request.IsEnable

	if ut := facades.Gorm.Save(&template); ut.Error != nil {
		http.Fail(ctx, "修改失败：%v", ut.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSpecificationOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoSpecificationOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var template model.ShpTemplateSpecification

	ft := facades.Gorm.Scopes(scope.Platform(ctx)).First(&template, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该信息")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", ft.Error)
		return
	}

	if dt := facades.Gorm.Delete(&template, "`id`=?", request.ID); dt.Error != nil {
		http.Fail(ctx, "删除失败：%v", dt.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSpecificationOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoSpecificationOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var template model.ShpTemplateSpecification

	ft := facades.Gorm.Scopes(scope.Platform(ctx)).First(&template, "`id`=?", request.ID)
	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该信息")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", ft.Error)
		return
	}

	if ut := facades.Gorm.Model(template).Update("is_enable", request.IsEnable); ut.Error != nil {
		http.Fail(ctx, "修改失败：%v", ut.Error)
		return
	}

	http.Success[any](ctx)
}

func ToSpecificationOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToSpecificationOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToSpecificationOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	tx.Model(&model.ShpTemplateSpecification{}).Count(&responses.Total)

	if responses.Total > 0 {

		var templates []model.ShpTemplateSpecification

		tx.
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&templates)

		responses.Data = make([]res.ToSpecificationOfPaginate, len(templates))

		for idx, item := range templates {
			responses.Data[idx] = res.ToSpecificationOfPaginate{
				ID:       item.ID,
				Name:     item.Name,
				Label:    item.Label,
				Options:  item.Options,
				IsEnable: item.IsEnable,
			}
		}
	}

	http.Success(ctx, responses)
}

func ToSpecificationOfOpening(c context.Context, ctx *app.RequestContext) {

	var templates []model.ShpTemplateSpecification

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&templates, "`is_enable`=?", util.Yes)

	responses := make([]res.ToSpecificationOfOpening, len(templates))

	for idx, item := range templates {

		responses[idx] = res.ToSpecificationOfOpening{
			ID:      item.ID,
			Name:    item.Name,
			Label:   item.Label,
			Options: item.Options,
		}
	}

	http.Success(ctx, responses)
}
