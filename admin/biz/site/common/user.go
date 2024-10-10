package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	req "project.io/shop/admin/http/request/site/common"
	res "project.io/shop/admin/http/response/site/common"
	"project.io/shop/model"
)

func ToUserByPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToUserByPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToUserByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	ter := facades.Gorm.WithContext(context.Background()).
		Select("1").
		Model(&model.SysUserBindRole{}).
		Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole))

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
		ter = ter.Where("`role_id`<>?", authConstants.CodeOfDeveloper)
	}

	tx := facades.Gorm.WithContext(c).Where("exists (?)", ter)

	tx.Model(&model.SysUser{}).Count(&responses.Total)

	if responses.Total > 0 {

		var roles []model.SysUser

		tx.
			Preload("BindRoles.Role").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Order("`id` desc").
			Find(&roles)

		responses.Data = make([]res.ToUserByPaginate, len(roles))

		for index, item := range roles {

			responses.Data[index] = res.ToUserByPaginate{
				ID:        item.ID,
				Nickname:  item.Nickname,
				Roles:     make([]res.ToUserByPaginateOfRoles, 0),
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.Username != nil {
				responses.Data[index].Username = *item.Username
			}

			if item.Mobile != nil {
				responses.Data[index].Mobile = *item.Mobile
			}

			if item.Email != nil {
				responses.Data[index].Email = *item.Email
			}

			for _, value := range item.BindRoles {
				if value.Role != nil {
					responses.Data[index].Roles = append(responses.Data[index].Roles, res.ToUserByPaginateOfRoles{
						ID:   value.Role.ID,
						Name: value.Role.Name,
					})
				}
			}
		}
	}

	http.Success(ctx, responses)

}

func DoUserByCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64 = 0

	vr := facades.Gorm.WithContext(context.Background()).Where("`id` IN (?)", request.Roles)

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
		vr = vr.Where("`id`<>?", authConstants.CodeOfDeveloper)
	}

	vr.Model(&model.SysRole{}).Count(&total)

	if int(total) != len(request.Roles) {
		http.Fail(ctx, "部分角色未找到")
		return
	}

	tx := facades.Gorm.Begin()

	user := model.SysUser{
		ID:       facades.Snowflake.Generate().String(),
		Nickname: request.Nickname,
		Username: &request.Username,
		Password: auth.Password(request.Password),
		IsEnable: request.IsEnable,
	}

	if request.Mobile != "" {
		user.Mobile = &request.Mobile
	}

	if request.Email != "" {
		user.Email = &request.Email
	}

	if result := tx.Create(&user); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	bindings := make([]model.SysUserBindRole, len(request.Roles))

	for index, item := range request.Roles {
		bindings[index] = model.SysUserBindRole{
			Platform:       auth.Platform(ctx),
			OrganizationID: auth.Organization(ctx),
			UserID:         user.ID,
			RoleID:         item,
		}
	}

	if result := tx.Create(&bindings); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	items := make([]string, len(bindings))

	for index, item := range bindings {
		items[index] = auth.NameOfRole(item.RoleID)
	}

	if ok, _ := facades.Casbin.AddRolesForUser(auth.NameOfUser(user.ID), items, auth.SPlatform(ctx)); !ok {
		tx.Rollback()
		http.Fail(ctx, "创建失败")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoUserByUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Gorm.
		Preload("BindRoles.Role").
		Where("exists (?)", facades.Gorm.
			Select("1").
			Model(model.SysUserBindRole{}).
			Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole)),
		).
		First(&user, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	//	现可用的角色
	rids := make([]uint, 0)

	for _, item := range user.BindRoles {
		if item.Role != nil {
			rids = append(rids, item.RoleID)
		}
	}

	creates := make([]model.SysUserBindRole, 0)
	deletes := make([]uint, 0)

	for _, item := range request.Roles {
		mark := true
		for _, value := range rids {
			if item == value {
				mark = false
			}
		}
		if mark {
			creates = append(creates, model.SysUserBindRole{
				UserID: user.ID,
				RoleID: item,
			})
		}
	}

	for _, item := range rids {
		mark := true
		for _, value := range request.Roles {
			if item == value {
				mark = false
			}
		}
		if mark {
			deletes = append(deletes, item)
		}
	}

	if len(creates) > 0 {

		var total int64 = 0

		vr := facades.Gorm.WithContext(context.Background()).
			Where("`id` IN (?)", request.Roles)

		if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
			vr = vr.Where("`id`<>?", authConstants.CodeOfDeveloper)
		}

		vr.Model(&model.SysRole{}).Count(&total)

		if int(total) != len(request.Roles) {
			http.Fail(ctx, "部分角色未找到")
			return
		}
	}

	tx := facades.Gorm.Begin()

	user.Nickname = request.Nickname
	user.IsEnable = request.IsEnable

	if request.Password != "" {
		user.Password = auth.Password(request.Password)
	}

	if request.Mobile != "" {
		user.Mobile = &request.Mobile
	}

	if request.Email != "" {
		user.Email = &request.Email
	}

	if uu := tx.Omit(clause.Associations).Save(&user); uu.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", uu.Error)
		return
	}

	if len(creates) > 0 {

		if cb := tx.Create(&creates); cb.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", cb.Error)
			return
		}

		items := make([]string, len(creates))

		for index, item := range creates {
			items[index] = auth.NameOfRole(item.RoleID)
		}

		if ok, _ := facades.Casbin.AddRolesForUser(auth.NameOfUser(user.ID), items); !ok {
			tx.Rollback()
			http.Fail(ctx, "修改失败")
			return
		}
	}

	if len(deletes) > 0 {

		if db := tx.Delete(&model.SysUserBindRole{}, "`user_id`=? and `role_id` IN (?)", user.ID, deletes); db.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", db.Error)
			return
		}

		for _, item := range deletes {

			if ok, _ := facades.Casbin.DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item)); !ok {
				tx.Rollback()
				http.Fail(ctx, "修改失败")
				return
			}
		}
	}

	if request.IsEnable != user.IsEnable {

		var bindings []model.SysUserBindRole

		_ = facades.Gorm.Model(&user).Association("BindRoles").Find(&bindings)

		if request.IsEnable == util.Yes && user.IsEnable == util.No {
			//	启用用户角色

			items := make([]string, len(bindings))

			for index, item := range bindings {
				items[index] = auth.NameOfRole(item.RoleID)
			}

			if ok, _ := facades.Casbin.AddRolesForUser(auth.NameOfUser(user.ID), items); !ok {
				tx.Rollback()
				http.Fail(ctx, "修改失败")
				return
			}

		} else if request.IsEnable == util.No && user.IsEnable == util.Yes {
			//	禁用用户角色

			for _, item := range bindings {

				if ok, _ := facades.Casbin.DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); !ok {
					tx.Rollback()
					http.Fail(ctx, "修改失败")
					return
				}
			}
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoUserByDelete(c context.Context, ctx *app.RequestContext) {

	id := ctx.Param("id")

	var user model.SysUser

	fu := facades.Gorm.
		Preload("BindRoles", func(t *gorm.DB) *gorm.DB {

			return t.
				Where("exists (?)", facades.Gorm.
					Select("1").
					Model(&model.SysRole{}).
					Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysUserBindRole, model.TableSysRole)),
				)
		}).
		First(&user, "`id`=?", id)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if du := tx.Delete(&user); du.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", du.Error)
		return
	}

	for _, item := range user.BindRoles {

		if ok, _ := facades.Casbin.DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); !ok {
			tx.Rollback()
			http.Fail(ctx, "修改失败")
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoUserByEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Gorm.
		Preload("BindRoles").
		Where("exists (?)", facades.Gorm.
			Select("1").
			Model(model.SysUserBindRole{}).
			Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole)),
		).
		First(&user, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if uu := tx.Omit(clause.Associations).Model(user).Update("is_enable", request.IsEnable); uu.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", uu.Error)
		return
	}

	if request.IsEnable != user.IsEnable {

		if request.IsEnable == util.Yes && user.IsEnable == util.No {
			//	启用用户角色

			items := make([]string, len(user.BindRoles))

			for index, item := range user.BindRoles {
				items[index] = auth.NameOfRole(item.RoleID)
			}

			if ok, _ := facades.Casbin.AddRolesForUser(auth.NameOfUser(user.ID), items); !ok {
				tx.Rollback()
				http.Fail(ctx, "修改失败")
				return
			}

		} else if request.IsEnable == util.No && user.IsEnable == util.Yes {
			//	禁用用户角色

			for _, item := range user.BindRoles {

				if ok, _ := facades.Casbin.DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); !ok {
					tx.Rollback()
					http.Fail(ctx, "修改失败")
					return
				}
			}
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}
