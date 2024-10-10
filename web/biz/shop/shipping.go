package shop

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	"project.io/shop/web/http/response/shop"
)

func ToShippings(c context.Context, ctx *app.RequestContext) {

	var shippings []model.ShpShipping

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `id` asc").Find(&shippings, "`is_enable`=?", util.Yes)

	responses := make([]shop.ToShippings, len(shippings))

	for idx, item := range shippings {
		responses[idx] = shop.ToShippings{
			ID:    item.ID,
			Name:  item.Name,
			Money: item.Money,
		}
	}

	http.Success(ctx, responses)
}
