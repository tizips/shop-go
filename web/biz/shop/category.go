package shop

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	res "project.io/shop/web/http/response/shop"
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
