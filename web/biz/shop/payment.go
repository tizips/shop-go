package shop

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/davecgh/go-spew/spew"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/mitchellh/mapstructure"
	"github.com/plutov/paypal/v4"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
	"strconv"
	"strings"
)

func ToPaymentOfChannel(c context.Context, ctx *app.RequestContext) {

	var channels []model.ShpPaymentChannel

	facades.Gorm.Scopes(scope.Platform(ctx)).Order("`order` asc, `id` asc").Group("channel").Find(&channels)

	responses := make([]res.ToPaymentOfChannel, len(channels))

	for idx, item := range channels {

		responses[idx] = res.ToPaymentOfChannel{
			ID:      item.ID,
			Channel: item.Channel,
		}
	}

	http.Success(ctx, responses)
}

func DoPaymentOfPaypal(c context.Context, ctx *app.RequestContext) {

	var request req.DoPaymentOfPaypal

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var payment model.ShpPayment

	fp := facades.Gorm.
		Scopes(scope.Platform(ctx)).
		Preload("Order").
		Preload("Channels").
		First(&payment, "`id`=? and `channel`=? and `user_id`=?", request.ID, model.ShpPaymentOfChannelPaypal, auth.ID(ctx))

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Order not found.")
		return
	} else if fp.Error != nil || payment.Order == nil {
		http.Fail(ctx, "Order query failed. Please try again later.")
		return
	} else if payment.Channels == nil {
		http.Fail(ctx, "No payment information found for this order.")
		return
	}

	var address model.ShpOrderAddress

	facades.Gorm.Scopes(scope.Platform(ctx)).First(&address, "`order_id`=?", payment.OrderID)

	var details []model.ShpOrderDetail

	facades.Gorm.Scopes(scope.Platform(ctx)).Find(&details, "`order_id`=?", payment.OrderID)

	base := lo.If(payment.Channels.IsDebug == util.Yes, paypal.APIBaseSandBox).Else(paypal.APIBaseLive)

	client, err := paypal.NewClient(payment.Channels.Key, payment.Channels.Secret, base)

	if err != nil {
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	var ext model.ShpPaymentChannelOfExtPayPal

	if err = mapstructure.Decode(payment.Channels.Ext, &ext); err != nil {
		spew.Dump(err)
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	unit := paypal.PurchaseUnitRequest{
		Amount: &paypal.PurchaseUnitAmount{
			Currency: "USD",
			Value:    strconv.FormatFloat(float64(payment.Money)/100, 'f', 2, 64),
			Breakdown: &paypal.PurchaseUnitAmountBreakdown{
				ItemTotal: &paypal.Money{
					Currency: "USD",
					Value:    strconv.FormatFloat(float64(payment.Order.TotalPrice)/100, 'f', 2, 64),
				},
				Shipping: &paypal.Money{
					Currency: "USD",
					Value:    strconv.FormatFloat(float64(payment.Order.CostShipping)/100, 'f', 2, 64),
				},
			},
		},
		InvoiceID: payment.ID,
		Items:     make([]paypal.Item, len(details)),
	}

	for idx, item := range details {
		unit.Items[idx] = paypal.Item{
			Name: lo.If(lo.RuneLength(item.Name) > 120, lo.Substring(item.Name, 0, 117)+"...").Else(item.Name),
			UnitAmount: &paypal.Money{
				Currency: "USD",
				Value:    strconv.FormatFloat(float64(item.Price)/100, 'f', 2, 64),
			},
			Quantity: fmt.Sprintf("%v", item.Quantity),
			SKU:      strings.Join(item.Specifications, ";"),
		}
	}

	units := []paypal.PurchaseUnitRequest{unit}
	source := &paypal.PaymentSource{
		Paypal: &paypal.PaymentSourcePaypal{
			ExperienceContext: paypal.PaymentSourcePaypalExperienceContext{
				PaymentMethodPreference: "UNRESTRICTED",
				BrandName:               payment.Channels.Name,
				Locale:                  "en-US",
				LandingPage:             "LOGIN",
				ShippingPreference:      "NO_SHIPPING",
				UserAction:              "PAY_NOW",
				ReturnURL:               ext.URL.Return,
				CancelURL:               ext.URL.Cancel,
			},
		},
	}
	appCtx := &paypal.ApplicationContext{}

	order, err := client.CreateOrder(c, paypal.OrderIntentCapture, units, source, appCtx)

	if err != nil {
		spew.Dump(err)
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	updates := map[string]any{
		"no":       order.ID,
		"currency": "USD",
	}

	if result := facades.Gorm.Model(&payment).Updates(updates); result.Error != nil {
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	responses := res.DoPaymentOfPaypal{}

	for _, item := range order.Links {
		if item.Rel == "payer-action" {
			responses.Link = item.Href
			break
		}
	}

	http.Success(ctx, responses)
}
