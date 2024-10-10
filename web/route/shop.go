package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"project.io/shop/web/biz/basic"
	"project.io/shop/web/biz/shop"
	middle "project.io/shop/web/http/middleware"
)

func ShopRouter(router *server.Hertz) {

	route := router.Group("shop")
	{
		banners := route.Group("banners")
		{
			banners.GET("", shop.ToBanners)
		}

		page := route.Group("page")
		{
			page.GET("code", shop.ToPage)
		}

		products := route.Group("products")
		{
			products.GET("", shop.ToProductOfPaginate)
			products.GET(":id", shop.ToProductOfInformation)
		}

		product := route.Group("product")
		{
			product.GET("hot", shop.ToProductOfHot)
			product.GET("recommended", shop.ToProductOfRecommended)
			product.GET("specification", shop.ToProductOfSpecification)
		}

		carts := route.Group("carts").Use(middleware.Auth())
		{
			carts.GET("", shop.ToCarts)
			carts.PUT(":id", shop.DoCartOfUpdate)
			carts.DELETE(":id", shop.DoCartOfDelete)
		}

		cart := route.Group("cart").Use(middleware.Auth())
		{
			cart.POST("", shop.DoCartOfCreate)
			cart.DELETE("batch", shop.DoCartOfDeletes)
		}

		orders := route.Group("orders").Use(middleware.Auth())
		{
			orders.GET("", shop.ToOrderOfPaginate)
			orders.GET(":id", shop.ToOrderOfInformation)
		}

		order := route.Group("order").Use(middleware.Auth())
		{
			order.POST("", shop.DoOrder)
		}

		payment := route.Group("payment").Use(middleware.Auth())
		{
			payment.POST("paypal", shop.DoPaymentOfPaypal)
		}

		notify := route.Group("notify").Use(middle.NotAuthorize())
		{
			notify.POST("paypal", shop.DoNotifyOfPaypal)
		}

		advertises := route.Group("advertises")
		{
			advertises.GET("", shop.ToAdvertises)
		}

		categories := route.Group("categories")
		{
			categories.GET("", shop.ToCategories)
		}

		blogs := route.Group("blogs")
		{
			blogs.GET("", shop.ToBlogOfPaginate)
			blogs.GET(":id", shop.ToBlogOfInformation)
		}

		shippings := route.Group("shippings")
		{
			shippings.GET("", shop.ToShippings)
		}

		wishlists := route.Group("wishlists")
		{
			wishlists.GET("", shop.ToWishlistOfPaginate)
			wishlists.DELETE(":id", shop.DoWishlistOfDelete)
		}

		wishlist := route.Group("wishlist")
		{
			wishlist.POST("", shop.DoWishlistOfCreate)
		}

		seo := route.Group("seo")
		{
			seo.GET("", shop.ToSEO)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.POST("logout", basic.DoLoginOfOut)
		}
	}
}
