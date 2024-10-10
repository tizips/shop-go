package store

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
	req "project.io/shop/admin/http/request/site/store"
	res "project.io/shop/admin/http/response/site/store"
	"project.io/shop/model"
)

func DoSecretOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSecretOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	secret := model.SysSecret{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Secret:         lo.RandomString(32, lo.AlphanumericCharset),
	}

	if result := facades.Gorm.Create(&secret); result.Error != nil {
		http.Fail(ctx, "生成失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoSecretOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoSecretOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var secret model.SysSecret

	fs := facades.Gorm.Scopes(scope.Platform(ctx)).First(&secret, "`id`=?", request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fs.Error)
		return
	}

	if result := facades.Gorm.Delete(&secret); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToSecretOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToSecretOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToSecretOfSecret]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.Scopes(scope.Platform(ctx))

	tx.Model(&model.SysSecret{}).Count(&responses.Total)

	if responses.Total > 0 {

		var secrets []model.SysSecret

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` DESC").
			Find(&secrets)

		responses.Data = make([]res.ToSecretOfSecret, len(secrets))

		for idx, item := range secrets {

			responses.Data[idx] = res.ToSecretOfSecret{
				ID:        item.ID,
				Name:      item.Name,
				Secret:    item.Secret,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
