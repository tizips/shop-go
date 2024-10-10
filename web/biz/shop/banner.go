package shop

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	res "project.io/shop/web/http/response/shop"
)

func ToBanners(c context.Context, ctx *app.RequestContext) {

	var banners []model.ShpBanner

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `created_at` desc").Find(&banners, "`is_enable`=?", util.Yes)

	responses := make([]res.ToBanners, len(banners))

	for idx, item := range banners {
		responses[idx] = res.ToBanners{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Button:      item.Button,
			Picture:     item.Picture,
			URL:         item.URL,
			Target:      item.Target,
			CreatedAt:   item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}
