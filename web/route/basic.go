package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"project.io/shop/web/biz/basic"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{

		login := route.Group("login").Use(middleware.Limiter(nil))
		{
			login.POST("account", basic.DoLoginOfAccount)
			login.POST("email", basic.DoLoginOfEMail)
		}

		register := route.Group("register").Use(middleware.Limiter(nil))
		{
			register.POST("email", basic.DoRegisterOfEMail)
		}

		upload := route.Group("upload").Use(middleware.Auth())
		{
			upload.POST("file", basic.DoUploadOfFile)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.GET("information", basic.ToAccountOfInformation)
			account.POST("logout", basic.DoLoginOfOut)
		}
	}
}
