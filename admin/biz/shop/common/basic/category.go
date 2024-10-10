package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "project.io/shop/admin/http/request/shop/common/basic"
	res "project.io/shop/admin/http/response/shop/common/basic"
	"project.io/shop/model"
)

func ToCategoryOfChildren(c context.Context, ctx *app.RequestContext) {

	var request req.ToCategoryOfChildren

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var categories []model.ShpCategory

	tx := facades.Gorm.WithContext(c)

	if request.All <= 0 {
		tx = tx.Where("`is_enable` = ?", util.Yes)
	}

	tx.
		Order("`order` asc").
		Find(&categories, "`parent_id`=?", request.Parent)

	responses := make([]res.ToCategoryOfChildren, len(categories))

	for idx, item := range categories {
		responses[idx] = res.ToCategoryOfChildren{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}
