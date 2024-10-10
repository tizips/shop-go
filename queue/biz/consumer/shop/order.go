package shop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/queue"
	"github.com/samber/lo"
	"gorm.io/gorm"
	constant "project.io/shop/constants/queue"
	"project.io/shop/model"
	"project.io/shop/queue/biz/producer/basic"
	"project.io/shop/queue/biz/producer/shop"
)

func OrderPaid() {

	q := queue.NewQueue()

	err := q.Consumer(doOrderPaid, constant.ShopOrderPaid)

	if err != nil {
		color.Errorf("[queue] shop order paid: %v", err)
		return
	}
}

func OrderRefund() {

	q := queue.NewQueue()

	err := q.Consumer(doOrderRefund, constant.ShopOrderRefund)

	if err != nil {
		color.Errorf("[queue] shop order refund: %v", err)
		return
	}
}

func OrderClosed() {

	q := queue.NewQueue()

	err := q.Consumer(doOrderClosed, constant.ShopOrderClosed, true)

	if err != nil {
		color.Errorf("[queue] shop order closed: %v", err)
		return
	}
}

func OrderCompleted() {

	q := queue.NewQueue()

	err := q.Consumer(doOrderCompleted, constant.ShopOrderCompleted, true)

	if err != nil {
		color.Errorf("[queue] shop order completed: %v", err)
		return
	}
}

func OrderLog() {

	q := queue.NewQueue()

	err := q.Consumer(DoOrderLog, constant.ShopOrderLog)

	if err != nil {
		color.Errorf("[queue] shop order log: %v", err)
		return
	}
}

func doOrderPaid(data []byte) {

	var body constant.ShopOrderPaidMessage

	var err error

	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	if body.Order == "" {

		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
			ID:     body.ID,
			Reason: "No matching order found.",
		})

		return
	}

	c := context.Background()

	var total int64 = 0

	if result := facades.Gorm.
		WithContext(c).
		Model(&model.ShpOrder{}).
		Where("`id`=? and `is_paid`=?", body.Order, util.Yes).
		Count(&total); result.Error != nil {
		basic.PublishError(data, constant.ShopOrderPaid, result.Error)
		return
	}

	if total > 0 {

		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
			ID:     body.ID,
			Reason: "The order has been paid for; no payment is required.",
		})

		return
	}

	if result := facades.Gorm.
		WithContext(c).
		Model(&model.ShpOrder{}).
		Where("`id`=? and `status`=?", body.Order, model.ShpOrderOfStatusClosed).
		Count(&total); result.Error != nil {
		basic.PublishError(data, constant.ShopOrderPaid, result.Error)
		return
	}

	if total > 0 {

		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
			ID:     body.ID,
			Reason: "The order has been closed and cannot be completed for payment.",
		})

		return
	}

	result := facades.Gorm.
		WithContext(c).
		Model(&model.ShpOrder{}).
		Where("id=? and `is_paid`=? and `status`=?", body.Order, util.No, model.ShpOrderOfStatusPay).
		Updates(map[string]any{
			"status":     model.ShpOrderOfStatusShipment,
			"payment_id": body.ID,
			"is_paid":    util.Yes,
		})

	if result.Error != nil {
		basic.PublishError(data, constant.ShopOrderPaid, result.Error)
		return
	} else if result.RowsAffected <= 0 {
		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
			ID:     body.ID,
			Reason: "No orders have been successfully paid.",
		})
		return
	}

	shop.PublishOrderLog(constant.ShopOrderLogMessage{
		Order:     body.Order,
		Action:    "confirmed",
		Content:   "Order confirmed, awaiting shipment.",
		CreatedAt: carbon.Now(),
	})
}

func doOrderRefund(data []byte) {

	var body constant.ShopOrderRefundMessage

	if err := json.Unmarshal(data, &body); err != nil {
		return
	}

	var order model.ShpOrder

	fo := facades.Gorm.First(&order, "`id`=?", body.ID)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		return
	} else if fo.Error != nil {
		basic.PublishError(data, constant.ShopOrderRefund, fo.Error)
		return
	}

	var refund uint = 0

	tx := facades.Gorm.Begin()

	for _, item := range body.Details {

		up := map[string]any{
			"returned": gorm.Expr("`returned`+?", item.Quantity),
			"refund":   item.Refund,
		}

		refund += item.Refund

		if result := tx.Model(&model.ShpOrderDetail{}).Where("`id`=? and `order_id`=?", item.ID, body.ID).Updates(up); result.Error != nil {
			tx.Rollback()
			basic.PublishError(data, constant.ShopOrderRefund, result.Error)
			return
		}
	}

	if result := tx.Model(&model.ShpOrder{}).Where("`id`=?", body.ID).Update("`refund`", gorm.Expr("`refund`+?", refund)); result.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopOrderRefund, result.Error)
		return
	}

	var total int64 = 0

	tx.Model(&model.ShpOrderDetail{}).Where("`order_id`=? and `quantity`>`returned`", order.ID).Count(&total)

	if total == 0 && order.Status != model.ShpOrderOfStatusClosed && order.Status != model.ShpOrderOfStatusCompleted { // 全额退款后，将进行中的订单手动关闭

		if result := tx.Model(&model.ShpOrder{}).Where("`id`=? and `prices`<=`refund`", body.ID).Updates(map[string]any{"status": model.ShpOrderOfStatusClosed}); result.Error != nil {
			tx.Rollback()
			basic.PublishError(data, constant.ShopOrderRefund, result.Error)
			return
		}
	}

	if order.PaymentID != nil {

		shop.PublishPaymentRefund(constant.ShopPaymentRefundMessage{
			ID:      *order.PaymentID,
			Order:   body.ID,
			Detail:  body.Detail,
			Service: body.Service,
			Money:   lo.If(body.Refund > 0, body.Refund).Else(refund),
			Reason:  body.Reason,
		})
	}

	tx.Commit()

}

func doOrderClosed(data []byte) {

	var body constant.ShopOrderClosedMessage

	if err := json.Unmarshal(data, &body); err != nil {
		return
	}

	if body.Order == "" {
		return
	}

	tx := facades.Gorm.Begin()

	result := tx.
		Model(&model.ShpOrder{}).
		Where("`id`=? and `status`=? and `user_id`=?", body.Order, model.ShpOrderOfStatusPay, body.UserID).
		Update("status", model.ShpOrderOfStatusClosed)

	if result.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopOrderPaid, result.Error)
		return
	} else if result.RowsAffected <= 0 {
		tx.Rollback()
		return
	}

	var details []model.ShpOrderDetail

	fd := tx.
		Where("EXISTS (?)",
			facades.Gorm.
				Select("1").
				Model(&model.ShpOrder{}).
				Where(fmt.Sprintf("`%s`.`order_id`=`%s`.`id` and `%s`.`id`=? and `%s`.`status`=? and `%s`.`user_id`=?", model.TableShpOrderDetail, model.TableShpOrder, model.TableShpOrder, model.TableShpOrder, model.TableShpOrder), body.Order, model.ShpOrderOfStatusClosed, body.UserID),
		).
		Find(&details)

	if fd.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopOrderPaid, result.Error)
		return
	}

	if len(details) > 0 {

		for _, item := range details {

			if rest := tx.Model(&model.ShpSku{}).Where("`id`=?", item.SkuID).Update("stock", gorm.Expr("`stock`+?", item.Quantity)); rest.Error != nil {
				tx.Rollback()
				basic.PublishError(data, constant.ShopOrderPaid, result.Error)
				return
			}
		}
	}

	tx.Commit()
}

func doOrderCompleted(data []byte) {

	id := string(data)

	var order model.ShpOrder

	fo := facades.Gorm.First(&order, "`id`=?", id)

	if errors.Is(fo.Error, gorm.ErrRecordNotFound) {
		return
	} else if fo.Error != nil {
		basic.PublishError(data, constant.ShopOrderCompleted, fo.Error)
		return
	}

	if order.Status != model.ShpOrderOfStatusShipment {
		return
	}

	tx := facades.Gorm.Begin()

	log := model.ShpLog{
		Platform:       order.Platform,
		CliqueID:       order.CliqueID,
		OrganizationID: order.OrganizationID,
		UserID:         order.UserID,
		OrderID:        order.ID,
		Action:         "completed",
		Content:        "The system has automatically confirmed receipt.",
	}

	if result := tx.Create(&log); result.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopOrderCompleted, result.Error)
		return
	}

	now := carbon.Now()

	updates := model.ShpOrder{
		Status:      model.ShpOrderOfStatusCompleted,
		CompletedAt: &now,
	}

	if result := tx.Model(&order).Select("Status", "CompletedAt").Updates(updates); result.Error != nil {
		tx.Rollback()
		basic.PublishError(data, constant.ShopOrderCompleted, result.Error)
		return
	}

	tx.Commit()

}

func DoOrderLog(data []byte) {

	var body constant.ShopOrderLogMessage

	if err := json.Unmarshal(data, &body); err != nil {
		return
	}

	if body.Order == "" && body.CreatedAt.IsZero() {
		return
	}

	var orders []model.ShpOrder

	facades.Gorm.Find(&orders, "`id`=?", body.Order)

	if len(orders) > 0 {

		logs := make([]model.ShpLog, len(orders))

		for idx, item := range orders {

			logs[idx] = model.ShpLog{
				Platform:       item.Platform,
				CliqueID:       item.CliqueID,
				OrganizationID: item.OrganizationID,
				UserID:         item.UserID,
				OrderID:        item.ID,
				Action:         body.Action,
				Content:        body.Content,
				CreatedAt:      body.CreatedAt,
			}
		}

		if result := facades.Gorm.Create(&logs); result.Error != nil {
			basic.PublishError(data, constant.ShopOrderLog, result.Error)
		}
	}
}
