package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	cOrg "project.io/shop/admin/biz/hr/clique/org"
	pOrg "project.io/shop/admin/biz/hr/platform/org"
	rOrg "project.io/shop/admin/biz/hr/region/org"
	sOrg "project.io/shop/admin/biz/hr/store/org"
)

func HRRouter(router *server.Hertz) {

	route := router.Group("hr")
	route.Use(middleware.Auth())
	{
		//com := route.Group("common")
		//{
		//}

		pla := route.Group("platform")
		{
			org := pla.Group("org")
			{
				organizations := org.Group("organizations")
				{
					organizations.GET(":id", pOrg.DoOrganizationOfInformation)
					organizations.GET("", middleware.Permission("hr.org.organization.paginate"), pOrg.ToOrganizationOfPaginate)
					organizations.PUT(":id", middleware.Permission("hr.org.organization.update"), pOrg.DoOrganizationOfUpdate)
					organizations.DELETE(":id", middleware.Permission("hr.org.organization.delete"), pOrg.DoOrganizationOfDelete)
				}

				organization := org.Group("organization")
				{
					organization.POST("", middleware.Permission("hr.org.organization.create"), pOrg.DoOrganizationOfCreate)
					organization.PUT("enable", middleware.Permission("hr.org.organization.enable"), pOrg.DoOrganizationOfEnable)
					organization.POST("enter", middleware.Permission("hr.org.organization.manage"), pOrg.DoOrganizationOfEnter)
				}
			}
		}

		clq := route.Group("clique")
		{
			org := clq.Group("org")
			{
				brands := org.Group("brands")
				{
					brands.GET("", middleware.Permission("hr.org.brand.list"), cOrg.ToBrands)
					brands.PUT(":id", middleware.Permission("hr.org.brand.update"), cOrg.DoBrandOfUpdate)
					brands.DELETE(":id", middleware.Permission("hr.org.brand.delete"), cOrg.DoBrandOfDelete)
				}

				brand := org.Group("brand")
				{
					brand.POST("", middleware.Permission("hr.org.brand.create"), cOrg.DoBrandOfCreate)
					brand.GET("opening", cOrg.ToBrandOfOpening)
				}

				organizations := org.Group("organizations")
				{
					organizations.GET(":id", cOrg.ToOrganizationOfInformation)
					organizations.GET("", middleware.Permission("hr.org.organization.paginate"), cOrg.ToOrganizationOfPaginate)
					organizations.PUT(":id", middleware.Permission("hr.org.organization.update"), cOrg.DoOrganizationOfUpdate)
					organizations.DELETE(":id", middleware.Permission("hr.org.organization.delete"), cOrg.DoOrganizationOfDelete)
				}

				organization := org.Group("organization")
				{
					organization.POST("", middleware.Permission("hr.org.organization.create"), cOrg.DoOrganizationOfCreate)
					organization.PUT("enable", middleware.Permission("hr.org.organization.enable"), cOrg.DoOrganizationOfEnable)
					organization.POST("enter", middleware.Permission("hr.org.organization.manage"), cOrg.DoOrganizationOfEnter)
				}
			}
		}

		reg := route.Group("region")
		{
			org := reg.Group("org")
			{
				organizations := org.Group("organizations")
				{
					organizations.GET(":id", rOrg.ToOrganizationOfInformation)
					organizations.GET("", middleware.Permission("hr.org.organization.paginate"), rOrg.ToOrganizationOfPaginate)
					organizations.PUT(":id", middleware.Permission("hr.org.organization.update"), rOrg.DoOrganizationOfUpdate)
					organizations.DELETE(":id", middleware.Permission("hr.org.organization.delete"), rOrg.DoOrganizationOfDelete)
				}

				organization := org.Group("organization")
				{
					organization.POST("", middleware.Permission("hr.org.organization.create"), rOrg.DoOrganizationOfCreate)
					organization.PUT("enable", middleware.Permission("hr.org.organization.enable"), rOrg.DoOrganizationOfEnable)
					organization.POST("enter", middleware.Permission("hr.org.organization.manage"), rOrg.DoOrganizationOfEnter)
				}
			}
		}

		sto := route.Group("store")
		{
			org := sto.Group("org")
			{
				organizations := org.Group("organization")
				{
					organizations.GET("", middleware.Permission("site.organization.information"), sOrg.ToOrganization)
					organizations.PUT("", middleware.Permission("site.organization.update"), sOrg.DoOrganization)
				}
			}
		}
	}
}
