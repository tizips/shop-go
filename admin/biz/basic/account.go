package basic

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/basic"
	res "project.io/shop/admin/http/response/basic"
	"project.io/shop/model"
)

func ToAccountOfInformation(c context.Context, ctx *app.RequestContext) {

	var user model.SysUser

	fu := facades.Gorm.First(&user, "`id`=?", auth.ID(ctx))

	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfInformation{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	if user.Username != nil {
		responses.Username = *user.Username
	}

	if user.Mobile != nil {
		responses.Mobile = *user.Mobile
	}

	if user.Email != nil {
		responses.Username = *user.Email
	}

	http.Success(ctx, responses)
}

func ToAccountOfPlatform(c context.Context, ctx *app.RequestContext) {

	var user model.SysUser

	fu := facades.Gorm.First(&user, "`id`=?", auth.ID(ctx))

	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfPlatform{
		Platform:     auth.Platform(ctx),
		Organization: facades.Cfg.GetString("app.title"),
	}

	if temporary, _ := auth.Temporary(c, ctx); temporary != nil {

		responses.Back = responses.Organization

		if temporary.Bak != nil {
			responses.Back = temporary.Bak.Organization
		}
	}

	if auth.Organization(ctx) != nil {
		// 查询组织名称

		responses.Org = *auth.Organization(ctx)

		var organization model.HROrganization

		fo := facades.Gorm.First(&organization, "`platform`=? and `id`=?", responses.Platform, responses.Org)

		if fo.Error == nil {
			responses.Organization = organization.Name
		}
	}

	http.Success(ctx, responses)
}

func ToAccountOfModules(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToAccountOfModules, 0)

	modules := auth.Modules(auth.Platform(ctx))

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		for _, item := range modules {
			responses = append(responses, res.ToAccountOfModules{
				Code: item.Code,
				Name: item.Name,
			})
		}

	} else if temporary, _ := auth.Temporary(c, ctx); temporary != nil {

		for _, item := range modules {
			responses = append(responses, res.ToAccountOfModules{
				Code: item.Code,
				Name: item.Name,
			})
		}

	} else {

		var moduleCodes []string

		facades.Gorm.
			Model(&model.SysRoleBindPermission{}).
			Distinct("module").
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("module", &moduleCodes)

		for _, item := range modules {

			if lo.Contains(moduleCodes, item.Code) {

				responses = append(responses, res.ToAccountOfModules{
					Code: item.Code,
					Name: item.Name,
				})
			}
		}
	}

	http.Success(ctx, responses)
}

func ToAccountOfPermissions(c context.Context, ctx *app.RequestContext) {

	var request req.ToAccountOfPermissions

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := make([]string, 0)

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		modules := auth.Modules(auth.Platform(ctx))

		for _, item := range modules {
			if item.Code == request.Module {
				responses = item.Permissions
				break
			}
		}

	} else if temporary, _ := auth.Temporary(c, ctx); temporary != nil {

		modules := auth.Modules(auth.Platform(ctx))

		for _, item := range modules {
			if item.Code == request.Module {
				responses = item.Permissions
				break
			}
		}
	} else {

		facades.Gorm.
			Model(&model.SysRoleBindPermission{}).
			Where("`module` = ?", request.Module).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("permission", &responses)
	}

	http.Success(ctx, responses)
}

func DoAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	updates := make(map[string]any)

	if request.Mobile != "" {
		updates["mobile"] = request.Mobile
	}

	if request.Email != "" {
		updates["email"] = request.Email
	}

	if request.Password != "" {
		updates["password"] = auth.Password(request.Password)
	}

	if len(updates) > 0 {

		if result := facades.Gorm.Model(&model.SysUser{}).Where("`id` = ?", auth.ID(ctx)).Updates(updates); result.Error != nil {
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	http.Success[any](ctx)
}

func DoAccountOfBack(c context.Context, ctx *app.RequestContext) {

	temporary, _ := auth.Temporary(c, ctx)

	if temporary == nil {
		http.Fail(ctx, "未找到可回退的商户信息")
		return
	}

	if temporary.HasBak() {

		if err := auth.DoTemporary(c, ctx, temporary.Bak.Platform, temporary.Bak.Org, temporary.Bak.Organization, temporary.Bak.Clique, temporary.Bak.Bak); err != nil {
			http.Fail(ctx, "商户切换失败：%v", err)
			return
		}
	} else {

		if err := auth.DoTemporaryOfDelete(c, ctx); err != nil {
			http.Fail(ctx, "商户切换失败：%v", err)
			return
		}
	}

	http.Success[any](ctx)
}
