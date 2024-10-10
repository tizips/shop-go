package shop

import (
	"encoding/json"
	"github.com/herhe-com/framework/facades"
	"project.io/shop/constants/queue"
)

func PublishPaymentSuccess(data queue.ShopPaymentSuccessMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopPaymentSuccess, []string{queue.ShopPaymentSuccess})
}

func PublishPaymentRefund(data queue.ShopPaymentRefundMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopPaymentRefund, []string{queue.ShopPaymentRefund})
}
