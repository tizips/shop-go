package shop

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToBlogOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToBlogOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToBlogOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	tx.Model(&model.ShpBlog{}).Count(&responses.Total)

	if responses.Total > 0 {

		var blogs []model.ShpBlog

		tx.
			Omit("Content").
			Order("`is_top` asc, `posted_at` desc, `id` DESC").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&blogs)

		responses.Data = make([]res.ToBlogOfPaginate, len(blogs))

		for idx, item := range blogs {
			responses.Data[idx] = res.ToBlogOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Summary:   item.Summary,
				Thumb:     item.Thumb,
				PostedAt:  item.PostedAt.ToDateString(),
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToBlogOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToBlogOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var blog model.ShpBlog

	facades.Gorm.Scopes(scope.Platform(ctx)).First(&blog, "`id`=?", request.ID)

	responses := res.ToBlogOfInformation{
		ID:        blog.ID,
		Name:      blog.Name,
		Thumb:     blog.Thumb,
		Summary:   blog.Summary,
		PostedAt:  blog.PostedAt.ToDateString(),
		Content:   blog.Content,
		CreatedAt: blog.CreatedAt.ToDateTimeString(),
	}

	http.Success(ctx, responses)
}
