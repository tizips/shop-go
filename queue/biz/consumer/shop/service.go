package shop

import (
	"github.com/gookit/color"
	"github.com/herhe-com/framework/queue"
	constant "project.io/shop/constants/queue"
)

func ServiceAgree() {

	q := queue.NewQueue()

	err := q.Consumer(doServiceAgree, constant.ShopServiceAgree, true)

	if err != nil {
		color.Errorf("[queue] shop service agree: %v", err)
		return
	}
}

func ServiceFinish() {

	q := queue.NewQueue()

	err := q.Consumer(doServiceFinish, constant.ShopServiceFinish, true)

	if err != nil {
		color.Errorf("[queue] shop service finish: %v", err)
		return
	}
}

func ServiceRefund() {

	q := queue.NewQueue()

	err := q.Consumer(doServiceRefund, constant.ShopServiceRefund)

	if err != nil {
		color.Errorf("[queue] shop service refund: %v", err)
		return
	}
}

func doServiceFinish(data []byte) {

	//id := string(data)

	//var service model.ShpService
	//
	//fs := facades.Gorm.First(&service, "`id`=?", id)
	//
	//if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
	//	return
	//} else if fs.Error != nil {
	//	basic.PublishError(data, constant.ShopServiceFinish, fs.Error)
	//	return
	//}
	//
	//if service.Status != model.ShpServiceOfStatusConfirmUser {
	//	return
	//}
	//
	//tx := facades.Gorm.Begin()
	//
	//log := model.ShpLog{
	//	Platform:       service.Platform,
	//	CliqueID:       service.CliqueID,
	//	OrganizationID: service.OrganizationID,
	//	UserID:         service.UserID,
	//	OrderID:        service.OrderID,
	//	DetailID:       service.DetailID,
	//	ServiceID:      &service.ID,
	//	Action:         "service_finish",
	//	Content:        "系统自动完结售后",
	//}
	//
	//if result := tx.Create(&log); result.Error != nil {
	//	tx.Rollback()
	//	basic.PublishError(data, constant.ShopServiceFinish, result.Error)
	//	return
	//}
	//
	//if result := tx.Model(&service).Update("status", model.ShpServiceOfStatusFinish); result.Error != nil {
	//	tx.Rollback()
	//	basic.PublishError(data, constant.ShopServiceFinish, result.Error)
	//	return
	//}
	//
	//tx.Commit()

}

func doServiceAgree(data []byte) {

	//id := string(data)
	//
	//var service model.ShpService
	//
	//fs := facades.Gorm.Preload("Products").First(&service, "`id`=?", id)
	//
	//if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
	//	return
	//} else if fs.Error != nil {
	//	basic.PublishError(data, constant.ShopServiceAgree, fs.Error)
	//	return
	//}
	//
	//if service.Status != model.ShpServiceOfStatusPending {
	//	return
	//}
	//
	//tx := facades.Gorm.Begin()
	//
	//log := model.ShpLog{
	//	Platform:       service.Platform,
	//	CliqueID:       service.CliqueID,
	//	OrganizationID: service.OrganizationID,
	//	UserID:         service.UserID,
	//	OrderID:        service.OrderID,
	//	DetailID:       service.DetailID,
	//	ServiceID:      &service.ID,
	//	Action:         "service_agree",
	//	Content:        "系统自动同意售后申请",
	//}
	//
	//if result := tx.Create(&log); result.Error != nil {
	//	tx.Rollback()
	//	basic.PublishError(data, constant.ShopServiceAgree, result.Error)
	//	return
	//}
	//
	//updates := map[string]any{
	//	"result": model.ShpServiceOfResultAgree,
	//	"status": model.ShpServiceOfStatusUser,
	//}
	//
	//if service.Type == model.ShpServiceOfTypeUnReceipt { // 未发货的订单，售后同意后直接退款
	//
	//	updates["status"] = model.ShpServiceOfStatusFinish
	//
	//	shop.PublishOrderRefund(constant.ShopOrderRefundMessage{
	//		ID:      service.OrderID,
	//		Detail:  service.DetailID,
	//		Service: &service.ID,
	//		Refund:  service.MoneyRefund,
	//		Reason:  "用户未收到货，售后退款",
	//		Details: lo.Map(service.Products, func(item model.ShpServiceProduct, index int) constant.ShopOrderRefundOfDetailMessage {
	//			return constant.ShopOrderRefundOfDetailMessage{
	//				ID:       item.DetailID,
	//				Quantity: item.Quantity,
	//				Refund:   item.Refund,
	//			}
	//		}),
	//	})
	//}
	//
	//if result := tx.Model(&model.ShpService{}).Where("`id`=?", service.ID).Updates(updates); result.Error != nil {
	//	tx.Rollback()
	//	basic.PublishError(data, constant.ShopServiceAgree, result.Error)
	//	return
	//}
	//
	//tx.Commit()

}

func doServiceRefund(data []byte) {

	//var body constant.ShopServiceRefundMessage
	//
	//if err := json.Unmarshal(data, &body); err != nil {
	//	return
	//}
	//
	//var service model.ShpService
	//
	//fs := facades.Gorm.First(&service, "`id`=?", body.ID)
	//
	//if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
	//	return
	//} else if fs.Error != nil {
	//	basic.PublishError(data, constant.ShopServiceRefund, fs.Error)
	//	return
	//}
	//
	//if service.Status != model.ShpServiceOfStatusFinish {
	//	return
	//}
	//
	//if body.Money != service.MoneyRefund { // 退款金额不一致，更新售后订单的退款总额
	//
	//	if result := facades.Gorm.Model(&service).Update("money_refund", body.Money); result.Error != nil {
	//		basic.PublishError(data, constant.ShopServiceRefund, result.Error)
	//		return
	//	}
	//
	//}

}
