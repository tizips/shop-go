package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	mBasic "project.io/shop/admin/biz/shop/common/basic"
	pOrder "project.io/shop/admin/biz/shop/platform/order"
	sBasic "project.io/shop/admin/biz/shop/store/basic"
	sCommodity "project.io/shop/admin/biz/shop/store/commodity"
	sMember "project.io/shop/admin/biz/shop/store/member"
	sOrder "project.io/shop/admin/biz/shop/store/order"
	sPay "project.io/shop/admin/biz/shop/store/pay"
)

func ShopRouter(router *server.Hertz) {

	route := router.Group("shop")
	route.Use(middleware.Auth())
	{
		com := route.Group("common")
		{
			basic := com.Group("basic")
			{
				category := basic.Group("category")
				{
					category.GET("children", mBasic.ToCategoryOfChildren)
				}

				setting := basic.Group("setting")
				{
					setting.POST("", middleware.Permission("shop.basic.setting.update"), mBasic.DoSetting)
					setting.GET("", middleware.Permission("shop.basic.setting.list"), mBasic.ToSetting)
				}
			}
		}

		pla := route.Group("platform")
		{
			order := pla.Group("order")
			{
				ordinaries := order.Group("ordinaries")
				{
					ordinaries.GET(":id", pOrder.ToOrdinaryOfInformation)
					ordinaries.GET("", middleware.Permission("shop.order.ordinary.paginate"), pOrder.ToOrdinaryOfPaginate)
				}

				ordinary := order.Group("ordinary")
				{
					ordinary.GET("logs", pOrder.ToOrdinaryOfLogs)
				}

				services := order.Group("services")
				{
					services.GET("", middleware.Permission("shop.order.service.paginate"), pOrder.ToServiceOfPaginate)
				}

				service := order.Group("service")
				{
					service.GET("logs", pOrder.ToServiceOfLogs)
				}

				appraisals := order.Group("appraisals")
				{
					appraisals.GET("", middleware.Permission("shop.order.appraisal.paginate"), pOrder.ToAppraisalOfPaginate)
				}
			}
		}

		//clq := route.Group("clique")
		//{
		//}
		//
		//reg := route.Group("region")
		//{
		//}

		sto := route.Group("store")
		{
			order := sto.Group("order")
			{
				ordinaries := order.Group("ordinaries")
				{
					ordinaries.GET(":id", sOrder.ToOrdinaryOfInformation)
					ordinaries.GET("", middleware.Permission("shop.order.ordinary.paginate"), sOrder.ToOrdinaryOfPaginate)
				}

				ordinary := order.Group("ordinary")
				{
					ordinary.POST("shipment", middleware.Permission("shop.order.ordinary.shipment"), sOrder.DoOrdinaryOfShipment)
					ordinary.POST("remark", sOrder.DoOrdinaryOfRemark)
					ordinary.GET("address", sOrder.ToOrdinaryOfAddress)
					ordinary.GET("logs", sOrder.ToOrdinaryOfLogs)
				}

				services := order.Group("services")
				{
					services.GET("", middleware.Permission("shop.order.service.paginate"), sOrder.ToServiceOfPaginate)
				}

				service := order.Group("service")
				{
					service.GET("logs", sOrder.ToServiceOfLogs)
					service.POST("handle", middleware.Permission("shop.order.service.handle"), sOrder.DoServiceOfHandle)
					service.POST("shipment", middleware.Permission("shop.order.service.handle"), sOrder.DoServiceOfShipment)
					service.POST("finish", middleware.Permission("shop.order.service.handle"), sOrder.DoServiceOfFinish)
					service.POST("closed", middleware.Permission("shop.order.service.handle"), sOrder.DoServiceOfClosed)
				}

				appraisals := order.Group("appraisals")
				{
					appraisals.GET("", middleware.Permission("shop.order.appraisal.paginate"), sOrder.ToAppraisalOfPaginate)
				}
			}

			commodity := sto.Group("commodity")
			{

				products := commodity.Group("products")
				{
					products.GET("", middleware.Permission("shop.commodity.product.paginate"), sCommodity.ToProductOfPaginate)
					products.PUT(":id", middleware.Permission("shop.commodity.product.update"), sCommodity.DoProductOfUpdate)
					products.DELETE(":id", middleware.Permission("shop.commodity.product.delete"), sCommodity.DoProductOfDelete)
					products.GET(":id", sCommodity.ToProductOfInformation)
				}

				product := commodity.Group("product")
				{
					product.POST("", middleware.Permission("shop.commodity.product.create"), sCommodity.DoProductOfCreate)
					product.PUT("enable", middleware.Permission("shop.commodity.product.enable"), sCommodity.DoProductOfEnable)
					product.GET("specification", middleware.Permission("shop.commodity.product.specification"), sCommodity.ToProductOfSpecification)
					product.POST("specification", middleware.Permission("shop.commodity.product.specification"), sCommodity.DoProductOfSpecification)
				}

				categories := commodity.Group("categories")
				{
					categories.GET("", middleware.Permission("shop.commodity.category.tree"), sCommodity.ToCategories)
					categories.PUT(":id", middleware.Permission("shop.commodity.category.update"), sCommodity.DoCategoryOfUpdate)
					categories.DELETE(":id", middleware.Permission("shop.commodity.category.delete"), sCommodity.DoCategoryOfDelete)
				}

				category := commodity.Group("category")
				{
					category.POST("", middleware.Permission("shop.commodity.category.create"), sCommodity.DoCategoryOfCreate)
					category.GET("opening", sCommodity.ToCategoryOfOpening)
				}

				specifications := commodity.Group("specifications")
				{
					specifications.GET("", middleware.Permission("shop.commodity.specification.paginate"), sCommodity.ToSpecificationOfPaginate)
					specifications.PUT(":id", middleware.Permission("shop.commodity.specification.update"), sCommodity.DoSpecificationOfUpdate)
					specifications.DELETE(":id", middleware.Permission("shop.commodity.specification.delete"), sCommodity.DoSpecificationOfDelete)
				}

				specification := commodity.Group("specification")
				{
					specification.POST("", middleware.Permission("shop.commodity.specification.create"), sCommodity.DoSpecificationOfCreate)
					specification.PUT("enable", middleware.Permission("shop.commodity.specification.enable"), sCommodity.DoSpecificationOfEnable)
					specification.GET("opening", sCommodity.ToSpecificationOfOpening)
				}
			}

			member := sto.Group("member")
			{
				users := member.Group("users")
				{
					users.GET("", middleware.Permission("shop.member.user.paginate"), sMember.ToUserOfPaginate)
				}
			}

			pay := sto.Group("pay")
			{
				channels := pay.Group("channels")
				{
					channels.GET("", middleware.Permission("shop.pay.channel.paginate"), sPay.ToChannelOfPaginate)
					channels.GET(":id", middleware.Permission("shop.pay.channel.paginate"), sPay.ToChannelOfInformation)
					channels.PUT(":id", middleware.Permission("shop.pay.channel.update"), sPay.DoChannelOfUpdate)
					channels.DELETE(":id", middleware.Permission("shop.pay.channel.delete"), sPay.DoChannelOfDelete)
				}

				channel := pay.Group("channel")
				{
					channel.POST("", middleware.Permission("shop.pay.channel.create"), sPay.DoChannelOfCreate)
					channel.PUT("enable", middleware.Permission("shop.pay.channel.enable"), sPay.DoChannelOfEnable)
				}
			}

			basic := sto.Group("basic")
			{

				banners := basic.Group("banners")
				{
					banners.GET("", middleware.Permission("shop.basic.banner.paginate"), sBasic.ToBannerOfPaginate)
					banners.PUT(":id", middleware.Permission("shop.basic.banner.update"), sBasic.DoBannerOfUpdate)
					banners.DELETE(":id", middleware.Permission("shop.basic.banner.delete"), sBasic.DoBannerOfDelete)
				}

				banner := basic.Group("banner")
				{
					banner.POST("", middleware.Permission("shop.basic.banner.create"), sBasic.DoBannerOfCreate)
					banner.PUT("enable", middleware.Permission("shop.basic.banner.enable"), sBasic.DoBannerOfEnable)
				}

				advertises := basic.Group("advertises")
				{
					advertises.GET("", middleware.Permission("shop.basic.advertise.paginate"), sBasic.ToAdvertiseOfPaginate)
					advertises.PUT(":id", middleware.Permission("shop.basic.advertise.update"), sBasic.DoAdvertiseOfUpdate)
					advertises.DELETE(":id", middleware.Permission("shop.basic.advertise.delete"), sBasic.DoAdvertiseOfDelete)
				}

				advertise := basic.Group("advertise")
				{
					advertise.POST("", middleware.Permission("shop.basic.advertise.create"), sBasic.DoAdvertiseOfCreate)
					advertise.PUT("enable", middleware.Permission("shop.basic.advertise.enable"), sBasic.DoAdvertiseOfEnable)
				}

				pages := basic.Group("pages")
				{
					pages.GET("", middleware.Permission("shop.basic.page.paginate"), sBasic.ToPageOfPaginate)
					pages.PUT(":id", middleware.Permission("shop.basic.page.update"), sBasic.DoPageOfUpdate)
					pages.DELETE(":id", middleware.Permission("shop.basic.page.delete"), sBasic.DoPageOfDelete)
					pages.GET(":id", sBasic.ToPageOfInformation)
				}

				page := basic.Group("page")
				{
					page.POST("", middleware.Permission("shop.basic.page.create"), sBasic.DoPageOfCreate)
				}

				blogs := basic.Group("blogs")
				{
					blogs.GET("", middleware.Permission("shop.basic.blog.paginate"), sBasic.ToBlogOfPaginate)
					blogs.PUT(":id", middleware.Permission("shop.basic.blog.update"), sBasic.DoBlogOfUpdate)
					blogs.DELETE(":id", middleware.Permission("shop.basic.blog.delete"), sBasic.DoBlogOfDelete)
					blogs.GET(":id", sBasic.ToBlogOfInformation)
				}

				blog := basic.Group("blog")
				{
					blog.POST("", middleware.Permission("shop.basic.blog.create"), sBasic.DoBlogOfCreate)
				}

				shippings := basic.Group("shippings")
				{
					shippings.GET("", middleware.Permission("shop.basic.shipping.paginate"), sBasic.ToShippingOfPaginate)
					shippings.PUT(":id", middleware.Permission("shop.basic.shipping.update"), sBasic.DoShippingOfUpdate)
					shippings.DELETE(":id", middleware.Permission("shop.basic.shipping.delete"), sBasic.DoShippingOfDelete)
				}

				shipping := basic.Group("shipping")
				{
					shipping.POST("", middleware.Permission("shop.basic.shipping.create"), sBasic.DoShippingOfCreate)
					shipping.PUT("enable", middleware.Permission("shop.basic.shipping.enable"), sBasic.DoShippingOfEnable)
					shipping.GET("opening", sBasic.ToShippingOfOpening)
				}

				seos := basic.Group("seos")
				{
					seos.GET("", middleware.Permission("shop.basic.seo.paginate"), sBasic.ToSEOOfPaginate)
					seos.PUT(":id", middleware.Permission("shop.basic.seo.update"), sBasic.DoSEOOfUpdate)
					seos.DELETE(":id", middleware.Permission("shop.basic.seo.delete"), sBasic.DoSEOOfDelete)
				}

				seo := basic.Group("seo")
				{
					seo.POST("", middleware.Permission("shop.basic.seo.create"), sBasic.DoSEOOfCreate)
				}
			}
		}
	}
}
