package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToPage(c context.Context, ctx *app.RequestContext) {

	var request req.ToPage

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShpPage

	fp := facades.Gorm.First(&page, "`code`=?", request.Code)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该页面信息")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败")
		return
	}

	responses := res.ToPage{
		Name:    page.Name,
		Content: page.Content,
	}

	http.Success(ctx, responses)
}
