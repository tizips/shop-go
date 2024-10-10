package member

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "project.io/shop/admin/http/request/shop/store/member"
	res "project.io/shop/admin/http/response/shop/store/member"
	"project.io/shop/model"
)

func ToUserOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToUserOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToUserOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(&model.ShpUser{}).Count(&responses.Total)

	if responses.Total > 0 {

		var users []model.ShpUser

		tx.
			Order("`created_at` desc, `id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&users)

		responses.Data = make([]res.ToUserOfPaginate, len(users))

		for idx, item := range users {
			responses.Data[idx] = res.ToUserOfPaginate{
				ID:        item.ID,
				Email:     item.Email,
				FirstName: item.FirstName,
				LastName:  item.LastName,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
