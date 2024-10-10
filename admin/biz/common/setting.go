package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-playground/validator/v10"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/support/util"
	"github.com/samber/lo"
	"project.io/shop/admin/http/response/common"
	"project.io/shop/model"
)

func ToSetting(c context.Context, ctx *app.RequestContext, module string) {

	var templates []model.ComSettingTemplate

	facades.Gorm.Where("`organization_id`=? or `organization_id` is null", auth.Organization(ctx)).Order("`order` asc, `id` asc").Find(&templates, "`module`=?", module)

	responses := make([]common.ToSetting, len(templates))

	for index, item := range templates {
		responses[index] = common.ToSetting{
			ID:         item.ID,
			Type:       item.Type,
			Label:      item.Label,
			Key:        item.Key,
			IsRequired: item.IsRequired,
			CreatedAt:  item.CreatedAt.ToDateTimeString(),
		}
	}

	var settings []model.ComSetting

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&settings, "`module`=?", module)

	for idx, item := range responses {
		for _, value := range settings {
			if item.Key == value.Key {
				responses[idx].Val = value.Val
				break
			}
		}
	}

	http.Success(ctx, responses)
}

func DoSetting(c context.Context, ctx *app.RequestContext, module string) {

	var request map[string]string

	if err := ctx.Bind(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var templates []model.ComSettingTemplate

	facades.Gorm.Find(&templates, "`module`=?", module)

	if facades.Validator == nil {
		http.Fail(ctx, "请先开启验证器")
		return
	}

	valid := facades.Validator.Engine().(*validator.Validate)

	for _, item := range templates {

		req, ok := request[item.Key]

		if item.IsRequired == model.ComSettingForIsRequiredOfYes && (!ok || lo.IsEmpty(req)) {
			http.BadRequest(ctx, item.Label+"不能为空")
			return
		}

		if !lo.IsEmpty(req) {

			var err error

			if item.Type == model.ComSettingForTypeOfEnable {
				err = valid.Var(req, "oneof=1 2")
			} else if item.Type == model.ComSettingForTypeOfURL || item.Type == model.ComSettingForTypeOfPicture {
				err = valid.Var(req, "url")
			} else if item.Type == model.ComSettingForTypeOfEmail {
				err = valid.Var(req, "email")
			}

			if err != nil {
				http.BadRequest(ctx, err)
				return
			}
		}
	}

	key := util.Keys("setting", module, *(auth.Organization(ctx)))

	if _, err := facades.Redis.Del(c, key).Result(); err != nil {
		http.Fail(ctx, "缓存清空失败")
		return
	}

	var settings []model.ComSetting

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&settings, "`module`=?", module)

	creates := make([]model.ComSetting, 0)
	updates := make([]model.ComSetting, 0)

	for idx, item := range request {

		cm := true

		for _, val := range settings {

			if idx == val.Key {
				cm = false

				if item != val.Val {
					updates = append(updates, model.ComSetting{
						ID:  val.ID,
						Val: item,
					})
				}

				break
			}
		}

		if cm {
			creates = append(creates, model.ComSetting{
				Platform:       auth.Platform(ctx),
				OrganizationID: auth.Organization(ctx),
				Module:         module,
				Key:            idx,
				Val:            item,
			})
		}
	}

	tx := facades.Gorm.Begin()

	if len(creates) > 0 {

		if result := tx.Create(&creates); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "保存失败：%v", result.Error)
			return
		}
	}

	if len(updates) > 0 {

		for _, item := range updates {
			if result := tx.Model(model.ComSetting{}).Where("`id`=?", item.ID).Update("val", item.Val); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "保存失败：%v", result.Error)
				return
			}
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}
