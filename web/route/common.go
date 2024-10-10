package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	com "project.io/shop/web/biz/common"
)

func CommonRouter(router *server.Hertz) {

	route := router.Group("common")
	{

		setting := route.Group("setting")
		{
			setting.GET("", com.ToSetting)
		}
	}
}
