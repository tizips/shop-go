package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "project.io/shop/admin/http/request/site/common"
	"project.io/shop/admin/http/response/site/common"
	"project.io/shop/model"
)

func ToAreaOfOpening(ctx context.Context, c *app.RequestContext) {

	var request req.ToAreaOfOpening

	if err := c.BindAndValidate(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var areas []model.SysArea

	facades.Gorm.
		Where("`parent_id`=?", request.Parent).
		Order("`id` asc").
		Find(&areas)

	responses := make([]common.ToAreaOfOpening, len(areas))

	for index, item := range areas {
		responses[index] = common.ToAreaOfOpening{
			ID:    item.ID,
			Name:  item.Name,
			Level: item.Level,
			Code:  item.Code,
		}
	}

	http.Success(c, responses)
}
