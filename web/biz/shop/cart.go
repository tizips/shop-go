package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
	"strconv"
	"strings"
)

func DoCartOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoCartOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var sku model.ShpSku

	fs := facades.Gorm.First(&sku, "`id`=?", request.SKU)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Product not found.")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "Operation failed, please try again later.")
		return
	}

	create := false

	var cart model.ShpCart

	fc := facades.Gorm.First(&cart, "`sku_id`=? and `product_id`=?", sku.ID, sku.ProductID)

	if fc.Error != nil && !errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.Fail(ctx, "Operation failed, please try again later.")
		return
	}

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {

		var product model.ShpProduct
		var picture model.ShpProductPicture

		fp := facades.Gorm.Model(&sku).Association("Product").Find(&product, "`id`=? and `is_enable`=?", sku.ProductID, util.Yes)

		if fp != nil {
			http.NotFound(ctx, "Product not found.")
			return
		}

		_ = facades.Gorm.First(&picture, "`product_id`=? and `is_default`=?", sku.ProductID, util.Yes)

		create = true

		cart = model.ShpCart{
			Platform:       sku.Platform,
			CliqueID:       sku.CliqueID,
			OrganizationID: sku.OrganizationID,
			UserID:         auth.ID(ctx),
			ProductID:      sku.ProductID,
			SkuID:          sku.ID,
			Code:           sku.Code,
			Name:           product.Name,
			Picture:        picture.URL,
			Specifications: make([]string, 0),
			Price:          sku.Price,
			Quantity:       request.Quantity,
			IsInvalid:      util.No,
		}

		if cart.Picture == "" {

		}

		if sku.Code != sku.ProductID && sku.Code != "" {

			codes := strings.Split(sku.Code, ":")

			var specification []model.ShpSpecification

			facades.Gorm.Order("`order` asc, `id` asc").Find(&specification, "`product_id`=? and `id` IN (?) and parent_id>?", sku.ProductID, codes, 0)

			if len(specification) != len(codes) {
				http.Success(ctx, "Failed to add to cart: some product specifications not found.")
				return
			}

			cart.Specifications = lo.Map(specification, func(item model.ShpSpecification, index int) string {
				return item.Name
			})
		}
	}

	if create {
		if err := facades.Gorm.Create(&cart).Error; err != nil {
			http.Fail(ctx, "Operation failed, please try again later.")
			return
		}
	} else {
		if result := facades.Gorm.Model(&cart).Update("quantity", gorm.Expr("`quantity`+?", 1)); result.Error != nil {
			http.Fail(ctx, "Operation failed, please try again later.")
			return
		}
	}

	http.Success[any](ctx)
}

func DoCartOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoCartOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var cart model.ShpCart

	fc := facades.Gorm.First(&cart, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx))

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Product not found in the cart.")
		return
	} else if fc.Error != nil {
		http.Fail(ctx, "Operation failed, please try again later.")
		return
	}

	if result := facades.Gorm.Model(&cart).Update("quantity", request.Quantity); result.Error != nil {
		http.Fail(ctx, "Failed to increase the quantity of products in the cart.")
		return
	}

	//if request.Action == "inc" {
	//
	//	if result := facades.Gorm.Model(&cart).Update("quantity", gorm.Expr("`quantity`+?", request.Quantity)); result.Error != nil {
	//		http.Fail(ctx, "Failed to increase the quantity of products in the cart.")
	//		return
	//	}
	//} else if request.Action == "dec" {
	//
	//	if cart.Quantity > request.Quantity {
	//
	//		if result := facades.Gorm.Model(&cart).Update("quantity", gorm.Expr("`quantity`-?", request.Quantity)); result.Error != nil {
	//			http.Fail(ctx, "Failed to decrease the quantity of products in the cart.")
	//			return
	//		}
	//	} else {
	//
	//		if result := facades.Gorm.Delete(&cart, "`id`=?", cart.ID); result.Error != nil {
	//			http.Fail(ctx, "Failed to remove the product from the cart.")
	//			return
	//		}
	//	}
	//}

	http.Success[any](ctx)
}

func DoCartOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoCartOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if result := facades.Gorm.Delete(&model.ShpCart{}, "`id`=? and `user_id`=?", request.ID, auth.ID(ctx)); result.Error != nil {
		http.Fail(ctx, "Failed to remove the product from the cart.")
		return
	}

	http.Success[any](ctx)
}

func DoCartOfDeletes(c context.Context, ctx *app.RequestContext) {

	var request req.DoCartOfDeletes

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if result := facades.Gorm.Delete(&model.ShpCart{}, "`id` IN (?) and `user_id`=?", request.IDS, auth.ID(ctx)); result.Error != nil {
		http.Fail(ctx, "Failed to remove the product from the cart.")
		return
	}

	http.Success[any](ctx)
}

func ToCarts(c context.Context, ctx *app.RequestContext) {

	var carts []model.ShpCart

	facades.Gorm.
		Preload("SKU").
		Order("`id` desc").
		Find(&carts, "`user_id`=?", auth.ID(ctx))

	responses := make([]res.ToCarts, len(carts))

	codes := make([]string, 0)

	for idx, item := range carts {

		responses[idx] = res.ToCarts{
			ID:             item.ID,
			Product:        item.ProductID,
			Name:           item.Name,
			Picture:        item.Picture,
			Code:           item.Code,
			Price:          item.Price,
			Quantity:       item.Quantity,
			Specifications: item.Specifications,
			IsInvalid:      item.IsInvalid,
			CreateAt:       item.CreatedAt.ToDateTimeString(),
		}

		if item.SKU == nil || item.SKU.Code != item.Code {
			responses[idx].IsInvalid = util.Yes
		}

		if responses[idx].IsInvalid != util.Yes {

			responses[idx].Price = item.SKU.Price

			if item.Code != item.ProductID && item.Code != "" {
				codes = append(codes, strings.Split(item.Code, ":")...)
			}
		}
	}

	if len(codes) > 0 {

		var specifications []model.ShpSpecification

		codes = lo.Uniq(codes)

		facades.Gorm.Find(&specifications, "`id` IN (?) and `parent_id`>?", codes, 0)

		for idx, item := range responses {

			code := strings.Split(item.Code, ":")

			spec := make([]string, 0)

			for _, val := range code {

				for _, specification := range specifications {

					if strconv.Itoa(int(specification.ID)) == val {
						spec = append(spec, specification.Name)
					}
				}
			}

			if len(code) == len(spec) {
				responses[idx].Specifications = spec
			}
		}
	}

	for idx, _ := range responses {
		responses[idx].Code = ""
	}

	http.Success(ctx, responses)
}
