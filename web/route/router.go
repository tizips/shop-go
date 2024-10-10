package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"project.io/shop/web/constants"
	middle "project.io/shop/web/http/middleware"
)

func Router(router *server.Hertz) {

	router.Use(middleware.Jwt(constants.JwtOfIssuerWithWeb))

	router.Use(middle.Authorize())

	CommonRouter(router)

	BasicRouter(router)

	ShopRouter(router)

}
