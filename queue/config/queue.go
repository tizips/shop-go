package config

import (
	"github.com/herhe-com/framework/facades"
	"project.io/shop/queue/biz/consumer/basic"
	"project.io/shop/queue/biz/consumer/shop"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("queue", map[string]any{
		"driver": cfg.Env("queue.driver"),
		"rabbitmq": map[string]any{
			"host":     cfg.Env("queue.rabbitmq.host", "127.0.0.1"),
			"port":     cfg.Env("queue.rabbitmq.port", 5672),
			"username": cfg.Env("queue.rabbitmq.username", "admin"),
			"password": cfg.Env("queue.rabbitmq.password", ""),
			"vhost":    cfg.Env("queue.rabbitmq.vhost", "/"),
		},
		"consumes": []func(){
			basic.Error,
			shop.PaymentSuccess,
			shop.PaymentRefund,
			shop.OrderPaid,
			shop.OrderRefund,
			shop.OrderClosed,
			shop.OrderReceived,
			shop.OrderCompleted,
			shop.OrderLog,
			shop.ServiceAgree,
			shop.ServiceFinish,
			shop.ServiceRefund,
		},
	})
}
