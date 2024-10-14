package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("payment", map[string]any{
		"paypal": map[string]any{
			"debug":     cfg.Env("payment.paypal.debug", false),
			"client_id": cfg.Env("payment.paypal.client_id", ""),
			"secret_id": cfg.Env("payment.paypal.secret_id", ""),
		},
	})
}
