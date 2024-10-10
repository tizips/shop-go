package commodity

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"project.io/shop/admin/helper/functions"
	req "project.io/shop/admin/http/request/shop/store/commodity"
	res "project.io/shop/admin/http/response/shop/store/commodity"
	"project.io/shop/model"
	"sort"
	"strconv"
	"strings"
)

func DoProductOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoProductOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var category model.ShpCategory

	fc := facades.Gorm.Scopes(scope.Platform(ctx)).First(&category, "`id`=? and `is_enable`=?", request.Category, util.Yes)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.Fail(ctx, "未找到该栏目")
		return
	}

	var count int64 = 0

	facades.Gorm.Scopes(scope.Platform(ctx)).Model(&model.ShpCategory{}).Where("`parent_id` = ?", request.Category).Count(&count)

	if count > 0 {
		http.Fail(ctx, "请使用子栏目挂靠产品")
		return
	}

	if category.Level == model.ShpCategoryOfLevel3 {

		facades.Gorm.Model(&category).Association("Parent").Find(&category.Parent, "`parent_id`=? and `is_enable`=?", category.ID, util.Yes)

		if category.Parent == nil {
			http.Fail(ctx, "未找到该栏目的父级栏目")
			return
		}
	}

	tx := facades.Gorm.Begin()

	product := model.ShpProduct{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Summary:        request.Summary,
		IsHot:          request.IsHot,
		IsRecommend:    request.IsRecommend,
		IsMultiple:     request.IsMultiple,
		IsFreeShipping: request.IsFreeShipping,
		IsFreeze:       request.IsFreeze,
		IsEnable:       request.IsEnable,
	}

	if category.Level == model.ShpCategoryOfLevel1 {
		product.I1CategoryID = category.ID
	} else if category.Level == model.ShpCategoryOfLevel2 {
		product.I2CategoryID = category.ID
		product.I1CategoryID = category.ParentID
	} else if category.Level == model.ShpCategoryOfLevel3 {
		product.I3CategoryID = category.ID
		product.I2CategoryID = category.ParentID
		product.I1CategoryID = category.Parent.ParentID
	}

	if result := tx.Create(&product); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	pictures := make([]model.ShpProductPicture, len(request.Pictures))

	for idx, item := range request.Pictures {

		pictures[idx] = model.ShpProductPicture{
			Platform:       product.Platform,
			CliqueID:       product.CliqueID,
			OrganizationID: product.OrganizationID,
			ProductID:      product.ID,
			URL:            item,
			Order:          uint8(idx + 1),
			IsDefault:      util.No,
		}

		if idx == 0 {
			pictures[idx].IsDefault = util.Yes
		}
	}

	if result := tx.Create(&pictures); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	information := model.ShpProductInformation{
		Platform:       product.Platform,
		CliqueID:       product.CliqueID,
		OrganizationID: product.OrganizationID,
		ProductID:      product.ID,
		Description:    request.Information,
		Attributes: lo.Map(request.Attributes, func(item req.DoProductOfAttribute, index int) model.ShpProductInformationOfAttribute {
			return model.ShpProductInformationOfAttribute{
				Label: item.Label,
				Value: item.Value,
			}
		}),
	}

	if result := tx.Create(&information); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	seo := model.ShpSEO{
		Platform:       product.Platform,
		CliqueID:       product.CliqueID,
		OrganizationID: product.OrganizationID,
		Channel:        model.ShpSEOForChannelOfProduct,
		ChannelID:      product.ID,
		Title:          request.Title,
		Keyword:        request.Keyword,
		Description:    request.Description,
	}

	if result := tx.Create(&seo); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	if request.IsMultiple != util.Yes {

		sku := model.ShpSku{
			ID:             facades.Snowflake.Generate().String(),
			Platform:       product.Platform,
			CliqueID:       product.CliqueID,
			OrganizationID: product.OrganizationID,
			ProductID:      product.ID,
			Code:           product.ID,
			Price:          request.Price,
			OriginPrice:    request.OriginPrice,
			CostPrice:      request.CostPrice,
			Stock:          request.Stock,
			Warn:           request.Warn,
			IsDefault:      util.Yes,
		}

		if result := tx.Create(&sku); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "创建失败：%v", result.Error)
			return
		}
	}

	//var client *elastic.Client
	//var err error
	//
	//if client, err = elastic.NewClient(
	//	elastic.SetURL(facades.Cfg.GetString("elasticsearch.host")),
	//	elastic.SetBasicAuth(facades.Cfg.GetString("elasticsearch.username"), facades.Cfg.GetString("elasticsearch.password")),
	//	elastic.SetSniff(false),
	//); err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "创建失败：%v", err)
	//	return
	//}
	//
	//_, err = client.Index().
	//	Index(shop.SearchIndexOfProduct).
	//	Id(product.ID).
	//	BodyJson(shop.SearchIndexOfProductBody{
	//		ID:          product.ID,
	//		Name:        product.Name,
	//		Picture:     product.Picture,
	//		Title:       seo.Title,
	//		Keyword:     seo.Keyword,
	//		Description: seo.Description,
	//		Text:        html.UnescapeString(strip.StripTags(information.Concept)),
	//		PriceMin:    0,
	//		PriceMax:    0,
	//		IsEnable:    product.IsEnable,
	//		CreatedAt:   product.CreatedAt.ToDateTimeString(),
	//	}).
	//	Do(c)
	//
	//if err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "创建失败：%v", err)
	//	return
	//}

	tx.Commit()

	http.Success[any](ctx)
}

func DoProductOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoProductOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("Pictures", func(t *gorm.DB) *gorm.DB { return t.Order("`order` asc") }).
		Preload("Information").
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfProduct) }).
		First(&product, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	categories := lo.Filter([]uint{product.I1CategoryID, product.I2CategoryID, product.I3CategoryID}, func(item uint, index int) bool {
		return item > 0
	})

	if request.Category != lo.Min(categories) {

		var category model.ShpCategory

		fc := facades.Gorm.Scopes(scope.Platform(ctx)).First(&category, "`id`=? and `is_enable`=?", request.Category, util.Yes)

		if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
			http.Fail(ctx, "未找到该栏目")
			return
		}

		var count int64 = 0

		facades.Gorm.Scopes(scope.Platform(ctx)).Model(&model.ShpCategory{}).Where("`parent_id` = ?", request.Category).Count(&count)

		if count > 0 {
			http.Fail(ctx, "请使用子栏目挂靠产品")
			return
		}

		if category.Level == model.ShpCategoryOfLevel3 {

			facades.Gorm.Model(&category).Association("Parent").Find(&category.Parent, "`parent_id`=? and `is_enable`=?", category.ID, util.Yes)

			if category.Parent == nil {
				http.Fail(ctx, "未找到该栏目的父级栏目")
				return
			}
		}

		product.I2CategoryID = 0
		product.I3CategoryID = 0

		if category.Level == model.ShpCategoryOfLevel1 {
			product.I1CategoryID = category.ID
		} else if category.Level == model.ShpCategoryOfLevel2 {
			product.I2CategoryID = category.ID
			product.I1CategoryID = category.ParentID
		} else if category.Level == model.ShpCategoryOfLevel3 {
			product.I3CategoryID = category.ID
			product.I2CategoryID = category.ParentID
			product.I1CategoryID = category.Parent.ParentID
		}
	}

	tx := facades.Gorm.Begin()

	// 规格切换
	if product.IsMultiple != request.IsMultiple {

		// 规格切换，清空原有 SKU 数据
		if result := tx.Delete(&model.ShpSku{}, "product_id=?", product.ID); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}

		// 从多规格切换成单规格，增加清空原有的属性数据
		if product.IsMultiple == util.Yes {

			if result := tx.Delete(&model.ShpSpecification{}, "product_id=?", product.ID); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	} else if request.IsMultiple == util.No {

		var price model.ShpSku

		_ = facades.Gorm.Model(&product).Association("Price").Find(&price, "`product_id`=?", product.ID)

		price.Price = request.Price
		price.OriginPrice = request.OriginPrice
		price.CostPrice = request.CostPrice
		price.Stock = request.Stock
		price.Warn = request.Warn

		if result := tx.Save(&price); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	product.Name = request.Name
	product.Summary = request.Summary
	product.IsHot = request.IsHot
	product.IsRecommend = request.IsRecommend
	product.IsMultiple = request.IsMultiple
	product.IsFreeShipping = request.IsFreeShipping
	product.IsFreeze = request.IsFreeze
	product.IsEnable = request.IsEnable

	product.Information.Description = request.Information
	product.Information.Attributes = lo.Map(request.Attributes, func(item req.DoProductOfAttribute, index int) model.ShpProductInformationOfAttribute {
		return model.ShpProductInformationOfAttribute{
			Label: item.Label,
			Value: item.Value,
		}
	})

	product.SEO.Title = request.Title
	product.SEO.Keyword = request.Keyword
	product.SEO.Description = request.Description

	if result := tx.Omit(clause.Associations).Save(&product); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	if result := tx.Omit(clause.Associations).Save(&product.SEO); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	if result := tx.Omit(clause.Associations).Save(&product.Information); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	creates := make([]model.ShpProductPicture, 0)
	updates := make([]model.ShpProductPicture, 0)
	deletes := make([]uint, 0)

	var idx uint8 = 1

	for _, item := range request.Pictures {

		mark := true

		for _, val := range product.Pictures {

			if item == val.URL {

				mark = false

				if val.Order != idx {

					val.Order = idx

					updates = append(updates, val)
				}

				break
			}
		}

		if mark {
			creates = append(creates, model.ShpProductPicture{
				Platform:       product.Platform,
				CliqueID:       product.CliqueID,
				OrganizationID: product.OrganizationID,
				ProductID:      product.ID,
				URL:            item,
				Order:          idx,
				IsDefault:      util.No,
			})
		}

		idx += 1
	}

	for _, item := range product.Pictures {

		mark := true

		for _, val := range request.Pictures {

			if item.URL == val {
				mark = false
			}
		}

		if mark {
			deletes = append(deletes, item.ID)
		}
	}

	if len(creates) > 0 {

		if result := tx.Create(&creates); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	if len(updates) > 0 {

		for _, item := range updates {
			if result := tx.Model(&item).Update("order", item.Order); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	}

	if len(deletes) > 0 {

		if result := tx.Delete(&model.ShpProductPicture{}, "`product_id`=? and `id` IN (?)", product.ID, deletes); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	_ = tx.Model(&product).Order("`order` asc").Limit(1).Association("Picture").Find(&product.Picture, "`product_id`=?", product.ID)

	if product.Picture != nil {

		if product.Picture.IsDefault != util.Yes {

			if result := tx.Model(&model.ShpProductPicture{}).Where("`product_id`=? and id=?", product.ID, product.Picture.ID).Update("is_default", util.Yes); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}

			if result := tx.Model(&model.ShpProductPicture{}).Where("`product_id`=? and id!=?", product.ID, product.Picture.ID).Update("is_default", util.No); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}

	} else {
		tx.Rollback()
		http.Fail(ctx, "轮播图查询失败，请稍后重试")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoProductOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoProductOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&product, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if result := tx.Delete(&product); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	//var client *elastic.Client
	//var err error
	//
	//if client, err = elastic.NewClient(
	//	elastic.SetURL(facades.Cfg.GetString("elasticsearch.host")),
	//	elastic.SetBasicAuth(facades.Cfg.GetString("elasticsearch.username"), facades.Cfg.GetString("elasticsearch.password")),
	//	elastic.SetSniff(false),
	//); err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "创建失败：%v", err)
	//	return
	//}
	//
	//_, err = client.Delete().
	//	Index(shop.SearchIndexOfProduct).
	//	Id(product.ID).
	//	Do(c)
	//
	//if err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "删除失败：%v", err)
	//	return
	//}

	tx.Commit()

	http.Success[any](ctx)
}

func DoProductOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoProductOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).First(&product, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查找失败：%v", fp.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if result := tx.Model(&product).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	//var client *elastic.Client
	//var err error
	//
	//if client, err = elastic.NewClient(
	//	elastic.SetURL(facades.Cfg.GetString("elasticsearch.host")),
	//	elastic.SetBasicAuth(facades.Cfg.GetString("elasticsearch.username"), facades.Cfg.GetString("elasticsearch.password")),
	//	elastic.SetSniff(false),
	//); err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "创建失败：%v", err)
	//	return
	//}
	//
	//_, err = client.Update().
	//	Index(shop.SearchIndexOfProduct).
	//	Id(product.ID).
	//	Doc(map[string]any{"is_enable": request.IsEnable}).
	//	Do(c)

	//if err != nil {
	//	tx.Rollback()
	//	http.Fail(ctx, "删除失败：%v", err)
	//	return
	//}

	tx.Commit()

	http.Success[any](ctx)
}

func ToProductOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToProductOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToProductOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c)

	if request.IsHot > 0 {
		tx = tx.Where("`is_hot` = ?", request.IsHot)
	}

	if request.IsRecommend > 0 {
		tx = tx.Where("`is_recommend` = ?", request.IsRecommend)
	}

	if request.IsMultiple > 0 {
		tx = tx.Where("`is_multiple` = ?", request.IsMultiple)
	}

	if request.IsFreeShipping > 0 {
		tx = tx.Where("`is_free_shipping` = ?", request.IsFreeShipping)
	}

	if request.IsFreeze > 0 {
		tx = tx.Where("`is_freeze` = ?", request.IsFreeze)
	}

	if request.IsEnable > 0 {
		tx = tx.Where("`is_enable = ?` = ?", request.IsEnable)
	}

	if request.Keyword != "" {
		tx = tx.Where("`name` LIKE ?", "%"+request.Keyword+"%")
	}

	tx.Model(&model.ShpProduct{}).Count(&responses.Total)

	if responses.Total > 0 {

		var products []model.ShpProduct

		tx.
			Scopes(scope.Platform(ctx)).
			Preload("Picture", "`is_default`=?", util.Yes).
			Preload("Price", "`is_default`=?", util.Yes).
			Order("`created_at` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&products)

		responses.Data = make([]res.ToProductOfPaginate, len(products))

		for idx, item := range products {

			responses.Data[idx] = res.ToProductOfPaginate{
				ID:             item.ID,
				Name:           item.Name,
				IsHot:          item.IsHot,
				IsRecommend:    item.IsRecommend,
				IsMultiple:     item.IsMultiple,
				IsFreeShipping: item.IsFreeShipping,
				IsFreeze:       item.IsFreeze,
				IsEnable:       item.IsEnable,
				CreatedAt:      item.CreatedAt.ToDateTimeString(),
			}

			if item.Picture != nil {
				responses.Data[idx].Picture = item.Picture.URL
			}

			if item.Price != nil {
				responses.Data[idx].Price = item.Price.Price
				responses.Data[idx].OriginPrice = item.Price.OriginPrice
			}
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

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("Pictures", func(t *gorm.DB) *gorm.DB { return t.Order("`order` asc") }).
		Preload("Information").
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", model.ShpSEOForChannelOfProduct) }).
		First(&product, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	responses := res.ToProductOfInformation{
		ID: product.ID,
		Categories: lo.Filter([]uint{product.I1CategoryID, product.I2CategoryID, product.I3CategoryID}, func(item uint, index int) bool {
			return item > 0
		}),
		Name:    product.Name,
		Summary: product.Summary,
		Pictures: lo.Map(product.Pictures, func(item model.ShpProductPicture, index int) string {
			return item.URL
		}),
		IsHot:          product.IsHot,
		IsRecommend:    product.IsRecommend,
		IsMultiple:     product.IsMultiple,
		IsFreeShipping: product.IsFreeShipping,
		IsFreeze:       product.IsFreeze,
		IsEnable:       product.IsEnable,
		Title:          product.SEO.Title,
		Keyword:        product.SEO.Keyword,
		Description:    product.SEO.Description,
		Information:    product.Information.Description,
		Attributes: lo.Map(product.Information.Attributes, func(item model.ShpProductInformationOfAttribute, index int) res.ToProductOfInformationWithAttribute {
			return res.ToProductOfInformationWithAttribute{
				Label: item.Label,
				Value: item.Value,
			}
		}),
		CreatedAt: product.CreatedAt.ToDateTimeString(),
	}

	if product.IsMultiple == util.No {

		facades.Gorm.Model(&product).Association("Price").Find(&product.Price, "`product_id`=?", product.ID)

		if product.Price != nil {
			responses.SKU = &res.ToProductOfSKU{
				ID:          product.Price.ID,
				Price:       product.Price.Price,
				OriginPrice: product.Price.OriginPrice,
				CostPrice:   product.Price.CostPrice,
				Stock:       product.Price.Stock,
				Warn:        product.Price.Warn,
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

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).First(&product, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	if product.IsMultiple == util.No {
		http.Fail(ctx, "该产品非多规格，无法进行此操作")
		return
	}

	responses := res.ToProductOfSpecification{
		ID:             product.ID,
		Specifications: nil,
		SKUS:           nil,
	}

	var specifications []model.ShpSpecification

	facades.Gorm.Preload("Specifications").Order("`order` asc").Find(&specifications, "`product_id`=? and parent_id=?", product.ID, 0)

	responses.Specifications = make([]res.ToProductOfSpecificationWithGroup, len(specifications))

	for idx, item := range specifications {
		responses.Specifications[idx] = res.ToProductOfSpecificationWithGroup{
			ID:   item.ID,
			Name: item.Name,
			Children: lo.Map(item.Specifications, func(value model.ShpSpecification, index int) res.ToProductOfSpecificationWithChildren {
				return res.ToProductOfSpecificationWithChildren{
					ID:        value.ID,
					Name:      value.Name,
					CreatedAt: value.CreatedAt.ToDateTimeString(),
				}
			}),
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	var skus []model.ShpSku

	facades.Gorm.Find(&skus, "`product_id`=?", product.ID)

	responses.SKUS = make([]res.ToProductOfSKU, len(skus))

	for idx, item := range skus {

		responses.SKUS[idx] = res.ToProductOfSKU{
			ID:          item.ID,
			Key:         item.Code,
			Price:       item.Price,
			OriginPrice: item.OriginPrice,
			CostPrice:   item.CostPrice,
			Stock:       item.Stock,
			Warn:        item.Warn,
			Picture:     item.Picture,
			IsDefault:   item.IsDefault,
		}
	}

	http.Success(ctx, responses)
}

func DoProductOfSpecification(c context.Context, ctx *app.RequestContext) {

	var request req.DoProductOfSpecification

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var product model.ShpProduct

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).First(&product, "`id`=?", request.ID)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fp.Error)
		return
	}

	if product.IsMultiple == util.No {
		http.Fail(ctx, "该产品非多规格，无法进行此操作")
		return
	}

	defaultCount := 0

	for _, item := range request.SKUS {
		if item.IsDefault == util.Yes {
			defaultCount += 1
		}
	}

	if defaultCount > 1 {
		http.Fail(ctx, "无法设置多个默认价格，请检查后重试")
		return
	}

	// 生成 规格 组合
	codes := functions.Descartes(lo.Map(request.Specification, func(item req.DoProductOfSpecificationWithGroup, index int) []uint {
		return lo.Map(item.Children, func(value req.DoProductOfSpecificationWithChildren, index int) uint {
			return value.ID
		})
	}))

	keys := make([]string, 0)

	// 检查是否有规格未设置 SKU
	for _, item := range codes {

		sort.Slice(item, func(i, j int) bool { return item[i] < item[j] })

		key := strings.Join(lo.Map(item, func(value uint, index int) string { return strconv.Itoa(int(value)) }), ":")

		keys = append(keys, key)

		if _, _, ok := lo.FindIndexOf(request.SKUS, func(value req.DoProductWithSKU) bool { return key == value.Key }); !ok {
			http.Fail(ctx, "部分规格未设置 SKU，请检查后重试1")
			return
		}
	}

	// 过滤多余 SKU
	request.SKUS = lo.Filter(request.SKUS, func(item req.DoProductWithSKU, index int) bool {
		_, _, ok := lo.FindIndexOf(keys, func(value string) bool { return item.Key == value })
		return ok
	})

	ids := make([]uint, 0)

	for _, item := range request.Specification {

		if !item.Virtual {
			ids = append(ids, item.ID)
		}

		for _, value := range item.Children {

			if !item.Virtual {
				ids = append(ids, value.ID)
			}
		}
	}

	var specifications []model.ShpSpecification
	var skus []model.ShpSku

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&specifications, "`product_id`=?", product.ID)
	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&skus, "`product_id`=?", product.ID)

	if len(ids) > 0 && len(specifications) == 0 {
		http.Fail(ctx, "部分规格未找到，请检查后重试")
		return
	}

	for _, item := range ids {

		if _, _, ok := lo.FindIndexOf(specifications, func(value model.ShpSpecification) bool { return item == value.ID }); !ok {
			http.Fail(ctx, "部分规格未找到，请检查后重试")
			return
		}
	}

	mapping := make(map[int]int)

	// 筛选出 Group 中需要执行操作的元素
	gc := make([]model.ShpSpecification, 0)
	gu := make([]model.ShpSpecification, 0)
	gd := make([]uint, 0)

	for index, item := range request.Specification {

		if item.Virtual { // 虚拟 ID，需要系统生成

			create := model.ShpSpecification{
				Platform:       product.Platform,
				CliqueID:       product.CliqueID,
				OrganizationID: product.OrganizationID,
				ProductID:      product.ID,
				Name:           item.Name,
				Order:          uint8(index + 1),
				Specifications: lo.Map(item.Children, func(value req.DoProductOfSpecificationWithChildren, idx int) model.ShpSpecification {

					mapping[((index+1)*10 + idx + 1)] = int(value.ID)

					return model.ShpSpecification{
						Platform:       product.Platform,
						CliqueID:       product.CliqueID,
						OrganizationID: product.OrganizationID,
						ProductID:      product.ID,
						Name:           value.Name,
						Order:          uint8(idx + 1),
					}
				}),
			}

			gc = append(gc, create)
		} else {

			// 判断 Group 是否需要求改
			for _, value := range specifications {

				if item.ID == value.ID && (item.Name != value.Name || value.Order != uint8(index+1)) {
					gu = append(gu, model.ShpSpecification{
						ID:    value.ID,
						Name:  item.Name,
						Order: uint8(index + 1),
					})
				}
			}

			for idx, value := range item.Children {

				if value.Virtual {

					create := model.ShpSpecification{
						Platform:       product.Platform,
						CliqueID:       product.CliqueID,
						OrganizationID: product.OrganizationID,
						ProductID:      product.ID,
						ParentID:       item.ID,
						Name:           value.Name,
						Order:          uint8(idx + 1),
					}

					mapping[(int(item.ID)*10 + idx + 1)] = int(value.ID)

					gc = append(gc, create)
				} else {

					for _, val := range specifications {

						if value.ID == val.ID && (value.Name != val.Name || val.Order != uint8(idx+1)) {
							gu = append(gu, model.ShpSpecification{
								ID:    val.ID,
								Name:  item.Name,
								Order: uint8(idx + 1),
							})
						}
					}
				}
			}
		}
	}

	gd = lo.FilterMap(specifications, func(value model.ShpSpecification, idx int) (uint, bool) {

		if _, _, ok := lo.FindIndexOf(ids, func(val uint) bool { return value.ID == val }); !ok {
			return value.ID, true
		}

		return 0, false
	})

	tx := facades.Gorm.Begin()

	if len(gc) > 0 {

		if result := tx.Omit(clause.Associations).Create(&gc); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}

		for _, item := range gc {

			if len(item.Specifications) > 0 {

				for key, _ := range item.Specifications {
					item.Specifications[key].ParentID = item.ID
				}

				if result := tx.Omit(clause.Associations).Create(&item.Specifications); result.Error != nil {
					tx.Rollback()
					http.Fail(ctx, "修改失败：%v", result.Error)
					return
				}
			}
		}

		for _, item := range gc {

			if item.ParentID > 0 {

				index := int(item.ID)*10 + int(item.Order)

				if val, ok := mapping[index]; ok {
					mapping[val] = int(item.ID)
				}

			} else {

				for _, value := range item.Specifications {

					idx := int(item.Order*10 + value.Order)

					if v, ok := mapping[idx]; ok {
						mapping[v] = int(value.ID)
					}
				}
			}
		}
	}

	if len(gu) > 0 {

		for _, item := range gu {

			if result := tx.Model(&item).Where("`id`=?", item.ID).Updates(map[string]any{"name": item.Name, "order": item.Order}); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	}

	if len(gd) > 0 {

		if result := tx.Delete(&model.ShpSpecification{}, "`id` IN (?)", gd); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	sc := make([]model.ShpSku, 0)
	su := make([]model.ShpSku, 0)
	sd := make([]string, 0)

	if len(mapping) > 0 { // 存在映射关系，将 SKU 的 Key 进行映射

		ck := make([]string, 0)
		caches := make(map[string]int, len(mapping))

		for key, val := range mapping {

			k := strconv.Itoa(key)

			ck = append(ck, k)
			caches[k] = val
		}

		for index, item := range request.SKUS {

			mark := false

			keys = strings.Split(item.Key, ":")

			for k, key := range keys {

				if idx, _, ok := lo.FindIndexOf(ck, func(value string) bool { return key == value }); ok {
					keys[k] = strconv.Itoa(caches[idx])
					mark = true
				}
			}

			if mark {

				sort.Strings(keys)

				request.SKUS[index].Key = strings.Join(keys, ":")
			}
		}
	}

	for _, item := range request.SKUS {

		mark := true

		for _, value := range skus {

			if item.Key == value.Code {
				mark = false
			}
		}

		if mark {
			sc = append(sc, model.ShpSku{
				ID:             facades.Snowflake.Generate().String(),
				Platform:       product.Platform,
				CliqueID:       product.CliqueID,
				OrganizationID: product.OrganizationID,
				ProductID:      product.ID,
				Code:           item.Key,
				Price:          item.Price,
				OriginPrice:    item.OriginPrice,
				CostPrice:      item.CostPrice,
				Stock:          item.Stock,
				Warn:           item.Warn,
				Picture:        item.Picture,
				IsDefault:      item.IsDefault,
			})
		}
	}

	for _, item := range skus {

		mark := true

		for _, value := range request.SKUS {

			if item.Code == value.Key && value.ID != "" {

				mark = false

				if item.Price != value.Price || item.OriginPrice != value.OriginPrice || item.CostPrice != value.CostPrice || item.Stock != value.Stock || item.Warn != value.Warn || item.Picture != value.Picture || item.IsDefault != value.IsDefault {
					su = append(su, model.ShpSku{
						ID:          item.ID,
						Price:       value.Price,
						OriginPrice: value.OriginPrice,
						CostPrice:   value.CostPrice,
						Stock:       value.Stock,
						Warn:        value.Warn,
						Picture:     value.Picture,
						IsDefault:   value.IsDefault,
					})
				}
			}
		}

		if mark {
			sd = append(sd, item.ID)
		}
	}

	if len(sc) > 0 {

		if result := tx.Omit(clause.Associations).Create(&sc); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	if len(su) > 0 {

		for _, item := range su {

			updates := map[string]any{
				"price":        item.Price,
				"origin_price": item.OriginPrice,
				"cost_price":   item.CostPrice,
				"stock":        item.Stock,
				"warn":         item.Warn,
				"picture":      item.Picture,
				"is_default":   item.IsDefault,
			}

			if result := tx.Model(&item).Where("`id`=?", item.ID).Updates(updates); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	}

	if len(sd) > 0 {

		if result := tx.Delete(&model.ShpSku{}, "`id` IN (?)", sd); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}
