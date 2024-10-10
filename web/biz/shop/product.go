package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToProductOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToProductOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToProductOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.WithContext(c).Where("`is_enable`=?", util.Yes)

	if request.Category > 0 {
		tx = tx.Where("`i1_category_id`=? or `i2_category_id`=? or `i3_category_id`=?", request.Category, request.Category, request.Category)
	}

	tx.Model(&model.ShpProduct{}).Count(&responses.Total)

	if responses.Total > 0 {

		var products []model.ShpProduct

		tx.
			Preload("Picture", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
			Preload("Price", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
			Order("`created_at` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&products)

		responses.Data = make([]res.ToProductOfPaginate, len(products))

		for idx, item := range products {

			responses.Data[idx] = res.ToProductOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.Price != nil {
				responses.Data[idx].Price = item.Price.Price
				responses.Data[idx].OriginPrice = item.Price.OriginPrice
			}

			if item.Picture != nil {
				responses.Data[idx].Picture = item.Picture.URL
			}
		}
	}

	http.Success(ctx, responses)
}

func ToProductOfHot(c context.Context, ctx *app.RequestContext) {

	var products []model.ShpProduct

	facades.Gorm.
		WithContext(c).
		Where("`is_hot`=?", util.Yes).
		Where("`is_enable`=?", util.Yes).
		Preload("Picture", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
		Preload("Price", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
		Order("RAND()").
		Limit(8).
		Find(&products)

	responses := make([]res.ToProductOfHot, len(products))

	for idx, item := range products {

		responses[idx] = res.ToProductOfHot{
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}

		if item.Price != nil {
			responses[idx].Price = item.Price.Price
			responses[idx].OriginPrice = item.Price.OriginPrice
		}

		if item.Picture != nil {
			responses[idx].Picture = item.Picture.URL
		}
	}

	http.Success(ctx, responses)
}

func ToProductOfRecommended(c context.Context, ctx *app.RequestContext) {

	var products []model.ShpProduct

	facades.Gorm.
		WithContext(c).
		Where("`is_recommend`=?", util.Yes).
		Where("`is_enable`=?", util.Yes).
		Preload("Picture", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
		Preload("Price", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
		Order("RAND()").
		Limit(10).
		Find(&products)

	responses := make([]res.ToProductOfRecommended, len(products))

	for idx, item := range products {

		responses[idx] = res.ToProductOfRecommended{
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}

		if item.Price != nil {
			responses[idx].Price = item.Price.Price
			responses[idx].OriginPrice = item.Price.OriginPrice
		}

		if item.Picture != nil {
			responses[idx].Picture = item.Picture.URL
		}
	}

	http.Success(ctx, responses)
}

func ToProductOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToProductOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("Price", func(t *gorm.DB) *gorm.DB { return t.Where("`is_default`=?", util.Yes) }).
		Preload("Pictures", func(t *gorm.DB) *gorm.DB { return t.Order("`order` asc, `id` asc") }).
		Preload("Information").
		First(&product, "`id`=? and `is_enable`=?", request.ID, util.Yes)

	responses := res.ToProductOfInformation{
		ID:      product.ID,
		Name:    product.Name,
		Summary: product.Summary,
		Pictures: lo.Map(product.Pictures, func(item model.ShpProductPicture, index int) string {
			return item.URL
		}),
		Price:       0,
		OriginPrice: 0,
		Information: product.Information.Description,
		Attributes: lo.Map(product.Information.Attributes, func(item model.ShpProductInformationOfAttribute, index int) res.Attribute {
			return res.Attribute{
				Label: item.Label,
				Value: item.Value,
			}
		}),
		IsMultiple:     product.IsMultiple,
		IsFreeShipping: product.IsFreeShipping,
		CreateAt:       product.CreatedAt.ToDateTimeString(),
	}

	if product.Price != nil {
		responses.Price = product.Price.Price
		responses.OriginPrice = product.Price.OriginPrice

		if responses.IsMultiple == util.No {
			responses.SKU = &res.SKU{
				ID:          product.Price.ID,
				Price:       product.Price.Price,
				OriginPrice: product.Price.OriginPrice,
				Stock:       product.Price.Stock,
			}
		}
	}

	http.Success(ctx, responses)
}

func ToProductOfSpecification(c context.Context, ctx *app.RequestContext) {

	var request req.ToProductOfSpecification

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		First(&product, "`id`=? and `is_enable`=?", request.ID, util.Yes)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该产品")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "产品查询失败：%v", fp.Error)
		return
	}

	if product.IsMultiple == util.No {
		http.Fail(ctx, "单规格产品不支持此接口")
		return
	}

	var Specifications []model.ShpSpecification

	facades.Gorm.
		Preload("Specifications", func(t *gorm.DB) *gorm.DB { return t.Order("`order` asc, `id` asc") }).
		Order("`order` asc, id asc").
		Find(&Specifications, "`product_id`=? and `parent_id`=?", product.ID, 0)

	var skus []model.ShpSku

	facades.Gorm.Find(&skus, "`product_id`=?", product.ID)

	responses := res.ToProductOfSpecification{
		ID: product.ID,
		Specifications: lo.Map(Specifications, func(item model.ShpSpecification, index int) res.Specification {
			return res.Specification{
				ID:   item.ID,
				Name: item.Name,
				Options: lo.Map(item.Specifications, func(value model.ShpSpecification, index int) res.Specification {
					return res.Specification{
						ID:   value.ID,
						Name: value.Name,
					}
				}),
			}
		}),
		SKUS: lo.Map(skus, func(item model.ShpSku, index int) res.SKU {
			return res.SKU{
				ID:          item.ID,
				Code:        item.Code,
				Price:       item.Price,
				OriginPrice: item.OriginPrice,
				Stock:       item.Stock,
				Picture:     item.Picture,
			}
		}),
	}

	http.Success(ctx, responses)
}
