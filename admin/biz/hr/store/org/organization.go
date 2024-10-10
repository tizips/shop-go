package org

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/hr/store/org"
	res "project.io/shop/admin/http/response/hr/store/org"
	"project.io/shop/model"
)

func DoOrganization(c context.Context, ctx *app.RequestContext) {

	var request req.DoOrganization

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	start := carbon.ParseByFormat(request.ValidStart, "Y-m-d")
	end := carbon.ParseByFormat(request.ValidEnd, "Y-m-d")

	if start.Gt(end) {
		http.Fail(ctx, "「有效期：开始」不能大于「有效期：结束」")
		return
	}

	var organization model.HROrganization

	fo := facades.Gorm.First(&organization, "`id`=? and `platform`=?", auth.Organization(ctx), authConstants.CodeOfStore)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fo.Error)
		return
	}

	var total int64 = 0

	if request.Province != organization.ProvinceId {

		facades.Gorm.Model(&model.SysArea{}).Where("`id`=? and `level`=?", request.Province, model.SysAreaOfLevelLv1).Count(&total)

		if total <= 0 {
			http.NotFound(ctx, "未找到该省份")
			return
		}
	}

	if request.City != organization.CityID {

		facades.Gorm.Model(&model.SysArea{}).Where("`id`=? and parent_id=? and `level`=?", request.City, request.Province, model.SysAreaOfLevelLv2).Count(&total)

		if total <= 0 {
			http.NotFound(ctx, "未找到该城市")
			return
		}
	}

	if request.Area != organization.AreaID {

		facades.Gorm.Model(&model.SysArea{}).Where("`id`=? and parent_id=? and `level`=?", request.Area, request.City, model.SysAreaOfLevelLv3).Count(&total)

		if total <= 0 {
			http.NotFound(ctx, "未找到该区县")
			return
		}
	}

	_ = facades.Gorm.Model(&organization).Association("Thumb").Find(&organization.Thumb)
	_ = facades.Gorm.Model(&organization).Association("Pictures").Find(&organization.Pictures)

	creates := lo.FilterMap(request.Pictures, func(item string, index int) (model.HROrganizationPicture, bool) {

		create := model.HROrganizationPicture{
			ID:             0,
			OrganizationID: organization.ID,
			CliqueID:       organization.CliqueID,
			URL:            item,
			Order:          uint8(index + 1),
		}

		if !lo.ContainsBy(organization.Pictures, func(value model.HROrganizationPicture) bool {
			return value.URL == item
		}) {
			return create, true
		}

		return create, false
	})

	updates := lo.FilterMap(organization.Pictures, func(item model.HROrganizationPicture, index int) (model.HROrganizationPicture, bool) {

		for key, value := range request.Pictures {

			order := uint8(key)

			if item.URL == value && item.Order != order {
				item.Order = order
				return item, true
			}
		}

		return item, false
	})

	deletes := lo.FilterMap(organization.Pictures, func(item model.HROrganizationPicture, index int) (uint, bool) {

		if !lo.Contains(request.Pictures, item.URL) {
			return item.ID, true
		}

		return 0, false
	})

	organization.Name = request.Name
	organization.ValidStart = start
	organization.ValidEnd = end
	organization.User = request.User
	organization.Telephone = request.Telephone
	organization.ProvinceId = request.Province
	organization.CityID = request.City
	organization.AreaID = request.Area
	organization.Address = request.Address
	organization.Longitude = request.Longitude
	organization.Latitude = request.Latitude

	tx := facades.Gorm.Begin()

	if result := tx.Save(&organization); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	if organization.Thumb != nil {

		if request.Thumb == "" {

			if result := tx.Delete(&organization.Thumb, "id=?", organization.Thumb.ID); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		} else if organization.Thumb.URL != request.Thumb {

			if result := tx.Model(&organization.Thumb).Update("url", request.Thumb); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	} else if request.Thumb != "" {

		picture := model.HROrganizationThumb{
			OrganizationID: organization.ID,
			CliqueID:       organization.CliqueID,
			URL:            request.Thumb,
		}

		if result := tx.Create(&picture); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
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

			if result := tx.Model(&model.HROrganizationPicture{}).Where("`id`=?", item.ID).Update("order", item.Order); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", result.Error)
				return
			}
		}
	}

	if len(deletes) > 0 {

		if result := tx.Delete(&model.HROrganizationPicture{}, "`id` IN (?)", deletes); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func ToOrganization(c context.Context, ctx *app.RequestContext) {

	var organization model.HROrganization

	fo := facades.Gorm.
		Preload("Thumb").
		Preload("Pictures", func(t *gorm.DB) *gorm.DB { return t.Order("`order` asc") }).
		First(&organization, "`id`=? and `platform`=?", auth.Organization(ctx), authConstants.CodeOfStore)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fo.Error != nil {
		http.Fail(ctx, "查询失败：%v", fo.Error)
		return
	}

	responses := res.ToOrganization{
		ID:          organization.ID,
		Brand:       organization.BrandID,
		Name:        organization.Name,
		ValidStart:  organization.ValidStart.ToDateString(),
		ValidEnd:    organization.ValidEnd.ToDateString(),
		User:        organization.User,
		Telephone:   organization.Telephone,
		Province:    organization.ProvinceId,
		City:        organization.CityID,
		Area:        organization.AreaID,
		Address:     organization.Address,
		Longitude:   organization.Longitude,
		Latitude:    organization.Latitude,
		Description: organization.Description,
		Pictures: lo.Map(organization.Pictures, func(item model.HROrganizationPicture, index int) string {
			return item.URL
		}),
		CreatedAt: organization.CreatedAt.ToDateTimeString(),
	}

	if organization.Thumb != nil {
		responses.Thumb = organization.Thumb.URL
	}

	http.Success(ctx, responses)
}
