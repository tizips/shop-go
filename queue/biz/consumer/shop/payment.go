package shop

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/queue"
	"github.com/plutov/paypal/v4"
	"github.com/samber/lo"
	"gorm.io/gorm"
	constant "project.io/shop/constants/queue"
	"project.io/shop/model"
	"project.io/shop/queue/biz/producer/basic"
	"strconv"
)

func PaymentSuccess() {

	q := queue.NewQueue()

	err := q.Consumer(doPaymentSuccess, constant.ShopPaymentSuccess)

	if err != nil {
		color.Errorf("[queue] shop payment success: %v", err)
		return
	}
}

func PaymentRefund() {

	q := queue.NewQueue()

	err := q.Consumer(doPaymentRefund, constant.ShopPaymentRefund)

	if err != nil {
		color.Errorf("[queue] shop payment refund: %v", err)
		return
	}
}

func doPaymentSuccess(data []byte) {

	//var body constant.ShopPaymentSuccessMessage
	//
	//if err := json.Unmarshal(data, &body); err != nil {
	//	return
	//}
	//
	//var payment model.ShpPayment
	//
	//fp := facades.Gorm.First(&payment, "`id`=? and `channel`=?", body.ID, body.Channel)
	//
	//if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
	//	return
	//} else if fp.Error != nil {
	//	basic.PublishError(data, constant.ShopPaymentSuccess, fp.Error)
	//	return
	//}
	//
	//if payment.IsConfirmed == util.Yes {
	//
	//	if payment.No != nil && *payment.No != body.No { // 已被支付，退款该笔支付
	//		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
	//			ID:     payment.ID,
	//			Money:  body.Payer,
	//			Reason: "该订单已被支付，无需重复支付",
	//			Ext:    body.Ext,
	//		})
	//		return
	//	}
	//
	//	return
	//} else if payment.Price != body.Total {
	//	shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
	//		ID:     payment.ID,
	//		Money:  body.Payer,
	//		Reason: "该订单价格不匹配，无法完成支付",
	//		Ext:    body.Ext,
	//	})
	//	return
	//}
	//
	//updates := map[string]any{
	//	"no":           body.No,
	//	"payer_price":  body.Payer,
	//	"currency":     body.Currency,
	//	"openid":       body.OpenID,
	//	"is_confirmed": util.Yes,
	//	"paid_at":      body.PaidAt,
	//}
	//
	//if body.Ext != nil {
	//
	//	ext, _ := json.Marshal(body.Ext)
	//
	//	updates["ext"] = string(ext)
	//}
	//
	//if result := facades.Gorm.Model(&payment).Updates(updates); result.Error != nil {
	//	basic.PublishError(data, constant.ShopPaymentSuccess, result.Error)
	//	return
	//}
	//
	//shop.PublishOrderPaid(constant.ShopOrderPaidMessage{
	//	ID:    body.ID,
	//	No:    body.No,
	//	Order: payment.OrderID,
	//	Batch: payment.Batch,
	//})
	//
	//shop.PublishOrderLog(constant.ShopOrderLogMessage{
	//	Order:     payment.OrderID,
	//	Batch:     payment.Batch,
	//	Action:    "pay",
	//	Content:   "用户支付成功",
	//	CreatedAt: body.PaidAt,
	//})

}

func doPaymentRefund(data []byte) {

	var body constant.ShopPaymentRefundMessage

	if err := json.Unmarshal(data, &body); err != nil {
		return
	}

	var payment model.ShpPayment

	fp := facades.Gorm.First(&payment, "`id`=? and `is_confirmed`=?", body.ID, util.Yes)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		return
	} else if fp.Error != nil {
		basic.PublishError(data, constant.ShopPaymentRefund, fp.Error)
		return
	}

	//var refunded uint = 0 // 已退款金额
	//
	//facades.Gorm.Model(&model.ShpRefund{}).Select("COALESCE(SUM(`money`), 0)").Where("`payment_id`=?", payment.ID).Scan(&refunded)
	//
	//balance := payment.Money - refunded // 可退款余额
	//
	//if balance <= 0 { // 可退款的余额不足
	//	return
	//}

	now := carbon.Now()

	refund := model.ShpRefund{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       payment.Platform,
		CliqueID:       payment.CliqueID,
		OrganizationID: payment.OrganizationID,
		UserID:         payment.UserID,
		OrderID:        &body.Order,
		DetailID:       body.Detail,
		PaymentID:      &payment.ID,
		ServiceID:      body.Service,
		No:             *payment.No,
		Channel:        payment.Channel,
		Money:          body.Money,
		Currency:       payment.Currency,
		IsConfirmed:    util.Yes,
		Remark:         body.Reason,
		RefundedAt:     &now,
		CreatedAt:      carbon.Carbon{},
		UpdatedAt:      carbon.Carbon{},
		DeletedAt:      gorm.DeletedAt{},
	}

	tx := facades.Gorm.Begin()

	if result := tx.Create(&refund); result.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopPaymentRefund, result.Error)
		return
	}

	if refund.Money > 0 && payment.Channel == model.ShpPaymentOfChannelPaypal {

		base := lo.If(facades.Cfg.GetBool("payment.paypal.debug"), paypal.APIBaseSandBox).Else(paypal.APIBaseLive)

		client, err := paypal.NewClient(facades.Cfg.GetString("payment.paypal.client_id"), facades.Cfg.GetString("payment.paypal.secret_id"), base)

		if err != nil {
			tx.Rollback()
			basic.PublishError(data, constant.ShopPaymentRefund, err)
			return
		}

		ctx := context.Background()

		order, err := client.GetOrder(ctx, *payment.No)

		if err != nil {
			tx.Rollback()
			basic.PublishError(data, constant.ShopPaymentRefund, err)
			return
		}

		captureID := ""

		for _, item := range order.PurchaseUnits {
			for _, value := range item.Payments.Captures {
				captureID = value.ID
				break
			}
		}

		params := paypal.RefundCaptureRequest{
			Amount: &paypal.Money{
				Currency: payment.Currency,
				Value:    strconv.FormatFloat(float64(refund.Money)/100, 'f', 2, 64),
			},
			InvoiceID:   payment.ID,
			NoteToPayer: "",
		}

		if result, err := client.RefundCapture(ctx, captureID, params); err != nil || result.Status != "COMPLETED" {
			tx.Rollback()
			basic.PublishError(data, constant.ShopPaymentRefund, err)
			return
		}
	} else {
		tx.Rollback()
		return
	}

	//if body.Service != nil && body.Money != refund.Money { // 售后订单存在，且退款金额不一致，更新售后订单数据
	//	shop.PublishServiceRefund(constant.ShopServiceRefundMessage{
	//		ID:    *body.Service,
	//		Money: body.Money,
	//	})
	//}

	tx.Commit()
}
