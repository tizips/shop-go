package common

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
)

func ToPermissions(c context.Context, ctx *app.RequestContext) {

	var responses []authConstants.Tree

	var ep = true
	var codes []string

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

		facades.Gorm.
			Scopes(scope.Platform(ctx)).
			Model(&model.SysRoleBindPermission{}).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("permission", &codes)

		ep = false
	}

	responses = auth.Trees(auth.Platform(ctx), ep, codes)

	http.Success(ctx, responses)
}
