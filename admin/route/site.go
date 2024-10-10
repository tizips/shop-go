package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"project.io/shop/admin/biz/site/common"
	"project.io/shop/admin/biz/site/platform"
	"project.io/shop/admin/biz/site/store"
)

func SiteRouter(router *server.Hertz) {

	route := router.Group("site")
	route.Use(middleware.Auth())
	{
		com := route.Group("common")
		{
			permissions := com.Group("permissions")
			{
				permissions.GET("", common.ToPermissions)
			}

			area := com.Group("area")
			{
				area.GET("opening", common.ToAreaOfOpening)
			}

			roles := com.Group("roles")
			{
				roles.GET(":id", common.ToRoleByInformation)
				roles.GET("", middleware.Permission("site.role.paginate"), common.ToRoleByPaginate)
				roles.PUT(":id", middleware.Permission("site.role.update"), common.DoRoleByUpdate)
				roles.DELETE(":id", middleware.Permission("site.role.delete"), common.DoRoleByDelete)
			}

			role := com.Group("role")
			{
				role.POST("", middleware.Permission("site.role.create"), common.DoRoleByCreate)
				role.GET("opening", common.ToRoleByOpening)
			}

			users := com.Group("users")
			{
				users.GET(":id", common.ToRoleByInformation)
				users.GET("", middleware.Permission("site.user.paginate"), common.ToUserByPaginate)
				users.PUT(":id", middleware.Permission("site.user.update"), common.DoUserByUpdate)
				users.DELETE(":id", middleware.Permission("site.user.delete"), common.DoUserByDelete)
			}

			user := com.Group("user")
			{
				user.POST("", middleware.Permission("site.user.create"), common.DoUserByCreate)
				user.PUT("enable", middleware.Permission("site.user.enable"), common.DoUserByEnable)
			}
		}

		pla := route.Group("platform")
		{
			areas := pla.Group("areas")
			{
				areas.GET("", middleware.Permission("site.area.paginate"), platform.ToAreaOfPaginate)
				areas.PUT(":id", middleware.Permission("site.area.update"), platform.DoAreaOfUpdate)
				areas.DELETE(":id", middleware.Permission("site.area.delete"), platform.DoAreaOfDelete)
			}

			area := pla.Group("area")
			{
				area.POST("", middleware.Permission("site.area.create"), platform.DoAreaOfCreate)
			}
		}

		sto := route.Group("store")
		{

			secrets := sto.Group("secrets")
			{
				secrets.GET("", middleware.Permission("site.secret.paginate"), store.ToSecretOfPaginate)
				secrets.DELETE(":id", middleware.Permission("site.secret.delete"), store.DoSecretOfDelete)
			}

			secret := sto.Group("secret")
			{
				secret.POST("", middleware.Permission("site.secret.create"), store.DoSecretOfCreate)
			}
		}

	}
}
