package shop

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToAdvertises(c context.Context, ctx *app.RequestContext) {

	var request req.ToAdvertises

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var advertises []model.ShpAdvertise

	facades.Gorm.
		WithContext(c).
		Scopes(scope.Platform(ctx)).
		Order("`order` asc, `id` asc").
		Find(&advertises, "`page`=? and `block`=?", request.Page, request.Block)

	responses := make([]res.ToAdvertises, len(advertises))

	for idx, item := range advertises {
		responses[idx] = res.ToAdvertises{
			ID:        item.ID,
			Title:     item.Title,
			Target:    item.Target,
			URL:       item.URL,
			Thumb:     item.Thumb,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}
