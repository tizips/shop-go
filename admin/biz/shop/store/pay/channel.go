package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/shop/store/pay"
	res "project.io/shop/admin/http/response/shop/store/pay"
	"project.io/shop/model"
)

func DoChannelOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoChannelOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	channel := model.ShpPaymentChannel{
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Name:           request.Name,
		Channel:        request.Channel,
		Key:            request.Key,
		Secret:         request.Secret,
		IsDebug:        request.IsDebug,
		Ext:            map[string]any{},
		Order:          request.Order.Order,
		IsEnable:       request.IsEnable,
	}

	if request.Channel == model.ShpPaymentOfChannelPaypal {

		ext := model.ShpPaymentChannelOfExtPayPal{}

		if request.PayPal != nil {
			ext.URL.Return = request.PayPal.ReturnURL
			ext.URL.Cancel = request.PayPal.CancelURL
		}

		if err := mapstructure.Decode(ext, &channel.Ext); err != nil {
			http.Fail(ctx, "创建失败：%v", err)
			return
		}
	}

	if result := facades.Gorm.Create(&channel); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoChannelOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoChannelOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var channel model.ShpPaymentChannel

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&channel, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	channel.Name = request.Name
	channel.Channel = request.Channel
	channel.Key = request.Key
	channel.Secret = request.Secret
	channel.Order = request.Order.Order
	channel.IsDebug = request.IsDebug
	channel.IsEnable = request.IsEnable

	if request.Channel == model.ShpPaymentOfChannelPaypal {

		ext := model.ShpPaymentChannelOfExtPayPal{}

		if request.PayPal != nil {
			ext.URL.Return = request.PayPal.ReturnURL
			ext.URL.Cancel = request.PayPal.CancelURL
		}

		if err := mapstructure.Decode(ext, &channel.Ext); err != nil {
			http.Fail(ctx, "创建失败：%v", err)
			return
		}
	}

	if result := facades.Gorm.Save(&channel); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoChannelOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoChannelOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var channel model.ShpPaymentChannel

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&channel, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&channel); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoChannelOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoChannelOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var channel model.ShpPaymentChannel

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&channel, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&channel).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToChannelOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToChannelOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToChannelOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c).Scopes(scope.Platform(ctx))

	tx.Model(&model.ShpPaymentChannel{}).Count(&responses.Total)

	if responses.Total > 0 {

		var banners []model.ShpPaymentChannel

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&banners)

		responses.Data = make([]res.ToChannelOfPaginate, len(banners))

		for idx, item := range banners {
			responses.Data[idx] = res.ToChannelOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Key:       item.Key,
				IsDebug:   item.IsDebug,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToChannelOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToChannelOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var channel model.ShpPaymentChannel

	fm := facades.Gorm.Scopes(scope.Platform(ctx)).First(&channel, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	responses := res.ToChannelOfInformation{
		ID:        channel.ID,
		Name:      channel.Name,
		Channel:   channel.Channel,
		Key:       channel.Key,
		Secret:    channel.Secret,
		IsDebug:   channel.IsDebug,
		PayPal:    nil,
		Order:     channel.Order,
		IsEnable:  channel.IsEnable,
		CreatedAt: channel.CreatedAt.ToDateTimeString(),
	}

	if channel.Channel == model.ShpPaymentOfChannelPaypal {

		responses.PayPal = &res.ToChannelOfPayPal{}

		var ext model.ShpPaymentChannelOfExtPayPal

		if err := mapstructure.Decode(channel.Ext, &ext); err != nil {
			http.Fail(ctx, "查找失败：%v", err)
			return
		}

		responses.PayPal.URL.Return = ext.URL.Return
		responses.PayPal.URL.Cancel = ext.URL.Cancel
	}

	http.Success(ctx, responses)
}
