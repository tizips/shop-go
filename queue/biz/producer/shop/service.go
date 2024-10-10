package shop

import (
	"encoding/json"
	"github.com/herhe-com/framework/facades"
	"project.io/shop/constants/queue"
)

func PublishServiceAgree(id string) error {
	return facades.Queue.Producer([]byte(id), queue.ShopServiceAgree, []string{queue.ShopServiceAgree}, 60*60*24*3)
}

func PublishServiceFinish(id string) error {
	return facades.Queue.Producer([]byte(id), queue.ShopServiceFinish, []string{queue.ShopServiceFinish}, 60*60*24*7)
}

func PublishServiceRefund(data queue.ShopServiceRefundMessage) error {

	body, _ := json.Marshal(data)

	return facades.Queue.Producer(body, queue.ShopServiceRefund, []string{queue.ShopServiceRefund})
}
