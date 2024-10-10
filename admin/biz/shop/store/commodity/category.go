package commodity

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/commodity"
	res "project.io/shop/admin/http/response/shop/store/commodity"
	"project.io/shop/model"
)

func ToCategories(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToCategories, 0)

	var categories []model.ShpCategory

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `id` asc").Find(&categories)

	children := make(map[uint][]res.ToCategories)

	for _, item := range categories {

		resp := res.ToCategories{
			ID:       item.ID,
			Level:    item.Level,
			Name:     item.Name,
			Order:    item.Order,
			IsEnable: item.IsEnable,
			Children: make([]res.ToCategories, 0),
		}

		if item.Level == model.ShpCategoryOfLevel1 {
			responses = append(responses, resp)
		} else {

			child, ok := children[item.ParentID]

			if !ok {
				children[item.ParentID] = make([]res.ToCategories, 0)
			}

			child = append(child, resp)

			children[item.ParentID] = child
		}
	}

	for idx, item := range responses {

		if child, ok := children[item.ID]; ok {

			for key, value := range child {

				if v, o := children[value.ID]; o {
					child[key].Children = v
				}
			}

			responses[idx].Children = child
		}
	}

	http.Success(ctx, responses)
}

func DoCategoryOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoCategoryOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	level := model.ShpCategoryOfLevel1

	if request.Parent > 0 {

		var parent model.ShpCategory

		fp := facades.Gorm.First(&parent, "`id`=?", request.Parent)

		if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
			http.NotFound(ctx, "未找到父级信息")
			return
		} else if fp.Error != nil {
			http.Fail(ctx, "父级查询失败：%v", fp.Error)
			return
		}

		if parent.Level == model.ShpCategoryOfLevel1 {
			level = model.ShpCategoryOfLevel2
		} else if parent.Level == model.ShpCategoryOfLevel2 {
			level = model.ShpCategoryOfLevel3
		} else {
			http.Fail(ctx, "父级查询失败")
			return
		}
	}

	category := model.ShpCategory{
		Level:          level,
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		ParentID:       request.Parent,
		Name:           request.Name,
		Order:          request.Order.Order,
		IsEnable:       request.IsEnable,
	}

	if result := facades.Gorm.Create(&category); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoCategoryOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoCategoryOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var category model.ShpCategory

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).First(&category, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到信息")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	category.Name = request.Name
	category.IsEnable = request.IsEnable
	category.Order = request.Order.Order

	if result := facades.Gorm.Save(&category); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoCategoryOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoCategoryOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var category model.ShpCategory

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).First(&category, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到信息")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	if result := facades.Gorm.Delete(&category, "`id`=?", request.ID); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToCategoryOfOpening(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToCategoryOfOpening, 0)

	var categories []model.ShpCategory

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `id` asc").Find(&categories, "`is_enable`=?", util.Yes)

	children := make(map[uint][]res.ToCategoryOfOpening)

	for _, item := range categories {

		resp := res.ToCategoryOfOpening{
			ID:       item.ID,
			Level:    item.Level,
			Name:     item.Name,
			Children: make([]res.ToCategoryOfOpening, 0),
		}

		if item.Level == model.ShpCategoryOfLevel1 {
			responses = append(responses, resp)
		} else {

			child, ok := children[item.ParentID]

			if !ok {
				children[item.ParentID] = make([]res.ToCategoryOfOpening, 0)
			}

			child = append(child, resp)

			children[item.ParentID] = child
		}
	}

	for idx, item := range responses {

		if child, ok := children[item.ID]; ok {

			for key, value := range child {

				if v, o := children[value.ID]; o {
					child[key].Children = v
				}
			}

			responses[idx].Children = child
		}
	}

	http.Success(ctx, responses)
}
