package shop

import (
	"encoding/json"
	"github.com/herhe-com/framework/facades"
	"project.io/shop/constants/queue"
)

func PublishOrderPaid(data queue.ShopOrderPaidMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopOrderPaid, []string{queue.ShopOrderPaid})
}

func PublishOrderRefund(data queue.ShopOrderRefundMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopOrderRefund, []string{queue.ShopOrderRefund})
}

func PublishOrderClosed(data queue.ShopOrderClosedMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopOrderClosed, []string{queue.ShopOrderClosed}, 60*10)
}

func PublishOrderReceived(order string) error {
	return facades.Queue.Producer([]byte(order), queue.ShopOrderReceived, []string{queue.ShopOrderReceived}, 60*60*24*15)
}

func PublishOrderCompleted(order string) error {
	return facades.Queue.Producer([]byte(order), queue.ShopOrderCompleted, []string{queue.ShopOrderCompleted}, 60*60*24*7)
}

func PublishOrderLog(data queue.ShopOrderLogMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopOrderLog, []string{queue.ShopOrderLog})
}
