package shop

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/plutov/paypal/v4"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
)

func DoNotifyOfPaypal(c context.Context, ctx *app.RequestContext) {

	var request req.DoNotifyOfPaypal

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var payment model.ShpPayment

	fp := facades.Gorm.Scopes(scope.Platform(ctx)).Preload("Order").First(&payment, "`no`=? and `channel`=?", request.Token, model.ShpPaymentOfChannelPaypal)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Order not found.")
		return
	} else if fp.Error != nil || payment.Order == nil {
		http.Fail(ctx, "Order query failed. Please try again later.")
		return
	}

	if payment.IsConfirmed == util.Yes {
		http.Fail(ctx, "The payment information has been confirmed; no duplicate action is needed.")
		return
	}

	if payment.Order.Status == model.ShpOrderOfStatusClosed {
		http.Fail(ctx, "This order has already been closed; no duplicate payment is required.")
		return
	} else if payment.Order.Status != model.ShpOrderOfStatusPay {
		http.Fail(ctx, "This order has already been paid; no duplicate payment is required.")
		return
	}

	now := carbon.Now()

	tx := facades.Gorm.Begin()

	updates := model.ShpPayment{
		IsConfirmed: util.Yes,
		PaidAt:      &now,
	}

	if result := tx.Model(&payment).Where("`is_confirmed`=?", util.No).Select("IsConfirmed", "PaidAt").Updates(updates); result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		http.Fail(ctx, "Payment information confirmation failed. Please try again later.")
		return
	}

	if result := tx.Model(&model.ShpOrder{}).
		Where("`id`=? and `user_id`=? and `status`=?", payment.OrderID, payment.UserID, model.ShpOrderOfStatusPay).
		Updates(map[string]any{
			"status":     model.ShpOrderOfStatusShipment,
			"payment_id": payment.ID,
			"is_paid":    util.Yes,
		}); result.Error != nil || result.RowsAffected <= 0 {
		tx.Rollback()
		http.Fail(ctx, "Payment information confirmation failed. Please try again later.")
		return
	}

	log := model.ShpLog{
		Platform:       payment.Platform,
		CliqueID:       payment.CliqueID,
		OrganizationID: payment.OrganizationID,
		UserID:         payment.UserID,
		OrderID:        payment.OrderID,
		Action:         "payment",
		Content:        "Payment Successful",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "Order placement failed. Please try again later.")
		return
	}

	base := lo.If(facades.Cfg.GetBool("payment.paypal.debug"), paypal.APIBaseSandBox).Else(paypal.APIBaseLive)

	client, err := paypal.NewClient(facades.Cfg.GetString("payment.paypal.client_id"), facades.Cfg.GetString("payment.paypal.secret_id"), base)

	if err != nil {
		http.Fail(ctx, "Payment initiation failed. Please try again later.")
		return
	}

	params := paypal.CaptureOrderRequest{
		PaymentSource: &paypal.PaymentSource{},
	}

	if result, err := client.CaptureOrder(c, request.Token, params); err != nil || result.Status != "COMPLETED" {
		tx.Rollback()
		http.Fail(ctx, "Payment information confirmation failed. Please try again later.")
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}
