package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/http"
	"project.io/shop/admin/http/request/shop/store/pay"
)

func ToPaymentOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request pay.ToPaymentOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

}
