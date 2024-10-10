package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"project.io/shop/admin/biz/common"
)

func ToSetting(c context.Context, ctx *app.RequestContext) {
	common.ToSetting(c, ctx, "shop")
}

func DoSetting(c context.Context, ctx *app.RequestContext) {
	common.DoSetting(c, ctx, "shop")
}
