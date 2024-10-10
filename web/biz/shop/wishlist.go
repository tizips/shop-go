package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func DoWishlistOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoWishlistOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fp := facades.Gorm.
		First(&product, "`id`=? and `is_enable`=?", request.ID, util.Yes)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Product not found.")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "Product query failed. Please try again later.")
		return
	}

	var total int64 = 0

	facades.Gorm.Model(&model.ShpWishlist{}).Where("`user_id`=? and `product_id`=?", auth.ID(ctx), request.ID).Count(&total)

	if total > 0 {
		http.Fail(ctx, "This product is already in your wishlist; no further action is required.")
		return
	}

	wishlist := model.ShpWishlist{
		ID:             0,
		Platform:       product.Platform,
		CliqueID:       product.CliqueID,
		OrganizationID: product.OrganizationID,
		UserID:         auth.ID(ctx),
		ProductID:      product.ID,
	}

	if result := facades.Gorm.Create(&wishlist); result.Error != nil {
		http.Fail(ctx, "Failed to add to wishlist. Please try again later.")
		return
	}

	http.Success[any](ctx)
}

func DoWishlistOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoWishlistOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if result := facades.Gorm.Delete(&model.ShpWishlist{}, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx)); result.Error != nil {
		http.Fail(ctx, "Failed to remove the product from the wishlist. Please try again later.")
		return
	}

	http.Success[any](ctx)
}

func ToWishlistOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToWishlistOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToWishlistOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c).Where("`user_id`=?", auth.ID(ctx))

	tx.Model(&model.ShpWishlist{}).Count(&responses.Total)

	if responses.Total > 0 {

		var wishlists []model.ShpWishlist

		tx.
			Preload("Product", func(t *gorm.DB) *gorm.DB {
				return t.
					Unscoped().
					Preload("Picture", func(j *gorm.DB) *gorm.DB {
						return j.Where("`is_default`=?", util.Yes)
					}).
					Where("`is_enable`=?", util.Yes)
			}).
			Preload("SKU", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&wishlists)

		responses.Data = make([]res.ToWishlistOfPaginate, len(wishlists))

		for idx, item := range wishlists {

			responses.Data[idx] = res.ToWishlistOfPaginate{
				ID:        item.ID,
				ProductID: item.ProductID,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.Product != nil {
				responses.Data[idx].Name = item.Product.Name
				responses.Data[idx].Price = item.SKU.Price

				if item.Product.Picture != nil {
					responses.Data[idx].Picture = item.Product.Picture.URL
				}
			}

			if item.SKU != nil {
				responses.Data[idx].Price = item.SKU.Price
			}
		}
	}

	http.Success(ctx, responses)
}
