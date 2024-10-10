package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	req "project.io/shop/admin/http/request/site/common"
	res "project.io/shop/admin/http/response/site/common"
	"project.io/shop/model"
	"strings"
)

func ToRoleByPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToRoleByPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToRoleByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).
		Where("`id`<>?", authConstants.CodeOfDeveloper)

	tx.Model(&model.SysRole{}).Count(&responses.Total)

	if responses.Total > 0 {

		var roles []model.SysRole

		tx.
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Order("`id` desc").
			Find(&roles)

		responses.Data = make([]res.ToRoleByPaginate, len(roles))

		for index, item := range roles {
			responses.Data[index] = res.ToRoleByPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Summary:   item.Summary,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToRoleByInformation(c context.Context, ctx *app.RequestContext) {

	id := ctx.Param("id")

	var role model.SysRole

	fr := facades.Gorm.Preload("BindPermissions").First(&role, "`id`=?", id)
	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(ctx, "查询失败：%v", fr.Error)
		return
	}

	responses := res.ToRoleByInformation{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: make([]string, len(role.BindPermissions)),
		Summary:     role.Summary,
		CreatedAt:   role.CreatedAt.ToDateTimeString(),
	}

	for index, item := range role.BindPermissions {
		responses.Permissions[index] = item.Permission
	}

	http.Success(ctx, responses)
}

func DoRoleByCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoRoleByCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	modules := auth.Modules(auth.Platform(ctx))

	//	检测用户是否非开发者

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

		// 查询该用户所有权限

		var userPermissions []model.SysRoleBindPermission

		facades.Gorm.
			Scopes(scope.Platform(ctx)).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Find(&userPermissions)

		modules = lo.FilterMap(modules, func(item authConstants.Module, index int) (authConstants.Module, bool) {

			module := item

			permissions := lo.Filter(item.Permissions, func(value string, index int) bool {

				if lo.ContainsBy(userPermissions, func(val model.SysRoleBindPermission) bool {
					return value == val.Permission && item.Code == val.Module
				}) {
					return true
				}

				return false
			})

			if len(permissions) > 0 {

				module.Permissions = permissions

				return module, true
			}

			return module, false
		})
	}

	//	检测提交的权限中实际可被授权的权限

	bindings := make([]model.SysRoleBindPermission, 0)
	permissions := make([][]string, 0)

	for _, item := range request.Permissions {

		for _, value := range modules {

			for _, val := range value.Permissions {

				if strings.HasPrefix(val, item) {

					permissions = append(permissions, []string{val, auth.SPlatform(ctx)})

					bindings = append(bindings, model.SysRoleBindPermission{
						Platform:       auth.Platform(ctx),
						OrganizationID: auth.Organization(ctx),
						Module:         value.Code,
						Permission:     val,
					})
				}
			}
		}
	}

	tx := facades.Gorm.Begin()

	role := model.SysRole{
		Platform:       auth.Platform(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Summary:        request.Summary,
	}

	if result := tx.Create(&role); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	for idx, _ := range bindings {
		bindings[idx].RoleID = role.ID
	}

	if result := tx.CreateInBatches(&bindings, 50); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	if ok, _ := facades.Casbin.AddPermissionsForUser(auth.NameOfRole(role.ID), permissions...); !ok {
		tx.Rollback()
		http.Fail(ctx, "创建失败")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoRoleByUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoRoleByUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var role model.SysRole

	fr := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("BindPermissions").
		First(&role, "`id`=?", request.ID)

	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(ctx, "查询失败：%v", fr.Error)
		return
	}

	if role.ID == authConstants.CodeOfDeveloper {
		http.Fail(ctx, "内置角色，无法修改")
		return
	}

	modules := auth.Modules(auth.Platform(ctx))

	//	检测用户是否非开发者

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

		// 查询该用户所有权限

		var userPermissions []model.SysRoleBindPermission

		facades.Gorm.
			Scopes(scope.Platform(ctx)).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Find(&userPermissions)

		modules = lo.FilterMap(modules, func(item authConstants.Module, index int) (authConstants.Module, bool) {

			module := item

			permissions := lo.Filter(item.Permissions, func(value string, index int) bool {

				if lo.ContainsBy(userPermissions, func(val model.SysRoleBindPermission) bool {
					return value == val.Permission && item.Code == val.Module
				}) {
					return true
				}

				return false
			})

			if len(permissions) > 0 {

				module.Permissions = permissions

				return module, true
			}

			return module, false
		})
	}

	//	提取提交的权限中实际可被授权的权限

	bindings := make([]model.SysRoleBindPermission, 0)
	permissions := make([][]string, 0)

	for _, item := range request.Permissions {

		for _, value := range modules {

			for _, val := range value.Permissions {

				if strings.HasPrefix(val, item) {

					permissions = append(permissions, []string{val, auth.SPlatform(ctx)})

					bindings = append(bindings, model.SysRoleBindPermission{
						Platform:       auth.Platform(ctx),
						OrganizationID: auth.Organization(ctx),
						Module:         value.Code,
						RoleID:         role.ID,
						Permission:     val,
					})
				}
			}
		}
	}

	//	提取需要创建的权限
	creates := lo.Filter(bindings, func(item model.SysRoleBindPermission, index int) bool {

		//	如果该权限不包含在已添加的权限中，视为添加权限
		if !lo.ContainsBy(role.BindPermissions, func(value model.SysRoleBindPermission) bool {
			return value.Permission == item.Permission
		}) {
			return true
		}

		return false
	})

	//	提取需要删除的权限
	deletes := lo.FilterMap(role.BindPermissions, func(item model.SysRoleBindPermission, index int) (uint, bool) {

		if !lo.ContainsBy(bindings, func(value model.SysRoleBindPermission) bool {
			return value.Permission == item.Permission
		}) {
			return item.ID, true
		}

		return 0, false
	})

	tx := facades.Gorm.Begin()

	role.Name = request.Name
	role.Summary = request.Summary

	if result := tx.Scopes(scope.Platform(ctx)).Omit(clause.Associations).Save(&role); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	if len(creates) > 0 {
		if result := tx.CreateInBatches(&creates, 50); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	if len(deletes) > 0 {
		if result := tx.Scopes(scope.Platform(ctx)).Delete(&model.SysRoleBindPermission{}, "`role_id`=? and `id` IN (?)", role.ID, deletes); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	if len(creates) > 0 || len(deletes) > 0 {

		//	清除旧权限
		if ok, _ := facades.Casbin.DeletePermissionsForUser(auth.NameOfRole(role.ID)); !ok {
			tx.Rollback()
			http.Fail(ctx, "修改失败")
			return
		}

		if ok, _ := facades.Casbin.AddPermissionsForUser(auth.NameOfRole(role.ID), permissions...); !ok {
			tx.Rollback()
			http.Fail(ctx, "修改失败")
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoRoleByDelete(c context.Context, ctx *app.RequestContext) {

	id := ctx.Param("id")

	var role model.SysRole

	fr := facades.Gorm.Preload("BindPermissions").First(&role, "`id`=?", id)
	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fr.Error != nil {
		http.Fail(ctx, "查询失败：%v", fr.Error)
		return
	}

	if role.ID == authConstants.CodeOfDeveloper {
		http.Fail(ctx, "内置角色，无法删除")
		return
	}

	tx := facades.Gorm.Begin()

	if dr := tx.Delete(&role, "`id`=?", role.ID); dr.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", dr.Error)
		return
	}

	//	清除旧角色
	if ok, _ := facades.Casbin.DeleteRole(auth.NameOfRole(role.ID)); !ok {
		tx.Rollback()
		http.Fail(ctx, "删除失败")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func ToRoleByOpening(c context.Context, ctx *app.RequestContext) {

	var roles []model.SysRole

	tx := facades.Gorm.WithContext(c)

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
		tx.Where("`id`<>?", authConstants.CodeOfDeveloper)
	}

	tx.Find(&roles)

	responses := make([]response.Opening[uint], len(roles))

	for index, item := range roles {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}
