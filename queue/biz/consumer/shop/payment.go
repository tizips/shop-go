package shop

import (
	"github.com/gookit/color"
	"github.com/herhe-com/framework/queue"
	constant "project.io/shop/constants/queue"
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

	//var body constant.ShopPaymentRefundMessage
	//
	//if err := json.Unmarshal(data, &body); err != nil {
	//	return
	//}
	//
	//var payment model.ShpPayment
	//
	//fp := facades.Gorm.First(&payment, "`id`=? and `is_confirmed`=?", body.ID, util.Yes)
	//
	//if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
	//	return
	//} else if fp.Error != nil {
	//	basic.PublishError(data, constant.ShopPaymentRefund, fp.Error)
	//	return
	//}
	//
	//if payment.OpenID == nil {
	//	basic.PublishError(data, constant.ShopPaymentRefund, errors.New("未找到可被退款的用户"))
	//	return
	//}
	//
	//var refunded uint = 0 // 已退款金额
	//
	//facades.Gorm.Model(&model.ShpRefund{}).Select("COALESCE(SUM(`money`), 0)").Where("`payment_id`=?", payment.ID).Scan(&refunded)
	//
	//balance := payment.PayerPrice - refunded // 可退款余额
	//
	//if balance <= 0 { // 可退款的余额不足
	//	return
	//}
	//
	//if body.OpenID == nil {
	//	body.OpenID = payment.OpenID
	//}
	//
	//now := carbon.Now()
	//
	//refund := model.ShpRefund{
	//	ID:             facades.Snowflake.Generate().String(),
	//	Platform:       payment.Platform,
	//	CliqueID:       payment.CliqueID,
	//	OrganizationID: payment.OrganizationID,
	//	UserID:         payment.UserID,
	//	OrderID:        lo.If(payment.OrderID != nil, payment.OrderID).Else(nil),
	//	DetailID:       body.Detail,
	//	PaymentID:      &payment.ID,
	//	ServiceID:      body.Service,
	//	No:             *payment.No,
	//	Channel:        payment.Channel,
	//	Money:          lo.If(body.Money > payment.PayerPrice, payment.PayerPrice).Else(body.Money),
	//	Currency:       payment.Currency,
	//	OpenID:         body.OpenID,
	//	IsConfirmed:    util.Yes,
	//	Remark:         body.Reason,
	//	RefundedAt:     &now,
	//	CreatedAt:      carbon.Carbon{},
	//	UpdatedAt:      carbon.Carbon{},
	//	DeletedAt:      gorm.DeletedAt{},
	//}
	//
	//if body.Money <= 0 && body.Detail == nil { // 无需查询子订单，直接退全部
	//	refund.Money = payment.PayerPrice
	//}
	//
	//if body.Detail != nil && (payment.Batch != nil || body.Money <= 0) { // 子订单 ID 存在时，如果为批量订单或退款金额没有，需要查询子订单完善相应信息
	//
	//	var detail model.ShpOrderDetail
	//
	//	fd := facades.Gorm.First(&detail, "`id`=? and `user_id`=?", *body.Detail, payment.UserID)
	//
	//	if errors.Is(fd.Error, gorm.ErrRecordNotFound) {
	//		basic.PublishError(data, constant.ShopPaymentRefund, errors.New("未找到可被退款的订单详情"))
	//		return
	//	} else if fd.Error != nil {
	//		basic.PublishError(data, constant.ShopPaymentRefund, fd.Error)
	//		return
	//	}
	//
	//	if payment.Batch != nil {
	//		refund.OrderID = &detail.OrderID
	//		refund.DetailID = &detail.ID
	//		refund.Platform = detail.Platform
	//		refund.CliqueID = detail.CliqueID
	//		refund.OrganizationID = detail.OrganizationID
	//	}
	//
	//	if refund.Money <= 0 {
	//		refund.Money = detail.Prices
	//	}
	//} else if body.Order != nil && payment.Batch != nil { // 批量订单时，需要查询订单详情完善退款信息
	//
	//	var order model.ShpOrder
	//
	//	fo := facades.Gorm.First(&order, "`id`=? and `user_id`=?", *body.Order, payment.UserID)
	//
	//	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
	//		basic.PublishError(data, constant.ShopPaymentRefund, errors.New("未找到可被退款的订单详情"))
	//		return
	//	} else if fo.Error != nil {
	//		basic.PublishError(data, constant.ShopPaymentRefund, fo.Error)
	//		return
	//	}
	//
	//	refund.OrderID = &order.ID
	//	refund.Platform = order.Platform
	//	refund.CliqueID = order.CliqueID
	//	refund.OrganizationID = order.OrganizationID
	//}
	//
	//if refund.Money > balance { // 如果订单的退款金额大于可退款余额，则取可退款余额
	//	refund.Money = balance
	//}
	//
	//ctx := context.Background()
	//
	//tx := facades.Gorm.Begin()
	//
	//if result := tx.Create(&refund); result.Error != nil {
	//	tx.Rollback()
	//	basic.PublishError(data, constant.ShopPaymentRefund, result.Error)
	//	return
	//}
	//
	//if refund.Money > 0 && payment.Channel == model.ShpPaymentOfChannelWeChat {
	//
	//	pay, err := wechat.NewPayment(&wechat.UserConfig{
	//		AppID:       facades.Cfg.GetString("open.wechat.payment.app_id"),
	//		MchID:       facades.Cfg.GetString("open.wechat.payment.mch_id"),
	//		MchApiV3Key: facades.Cfg.GetString("open.wechat.payment.key"),
	//		CertPath:    facades.Cfg.GetString("open.wechat.payment.cert_path"),
	//		KeyPath:     facades.Cfg.GetString("open.wechat.payment.key_path"),
	//		SerialNo:    facades.Cfg.GetString("open.wechat.payment.serial_no"),
	//		Debug:       false,
	//	})
	//
	//	if err != nil {
	//		tx.Rollback()
	//		basic.PublishError(data, constant.ShopPaymentRefund, err)
	//		return
	//	}
	//
	//	resp, err := pay.Refund.Refund(ctx, &request.RequestRefund{
	//		TransactionID: *payment.No,
	//		OutRefundNo:   refund.ID,
	//		Reason:        body.Reason,
	//		Amount: &request.RefundAmount{
	//			Refund:   int(refund.Money),
	//			Total:    int(payment.Price),
	//			Currency: payment.Currency,
	//		},
	//	})
	//
	//	if err != nil {
	//		tx.Rollback()
	//		basic.PublishError(data, constant.ShopPaymentRefund, err)
	//		return
	//	} else if resp.Code != "" {
	//		tx.Rollback()
	//		basic.PublishError(data, constant.ShopPaymentRefund, errors.New(resp.Message))
	//		return
	//	} else if resp.ErrCode != "" {
	//		tx.Rollback()
	//		basic.PublishError(data, constant.ShopPaymentRefund, errors.New(resp.ErrMsg))
	//		return
	//	}
	//} else {
	//	tx.Rollback()
	//	return
	//}
	//
	//if body.Service != nil && body.Money != refund.Money { // 售后订单存在，且退款金额不一致，更新售后订单数据
	//	shop.PublishServiceRefund(constant.ShopServiceRefundMessage{
	//		ID:    *body.Service,
	//		Money: body.Money,
	//	})
	//}
	//
	//tx.Commit()
}
