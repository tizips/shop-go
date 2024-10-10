package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("payment", map[string]any{
		"paypal": map[string]any{
			"live":      cfg.Env("payment.paypal.live", true),
			"client_id": cfg.Env("payment.paypal.client_id", ""),
			"secret_id": cfg.Env("payment.paypal.secret_id", ""),
			"url": map[string]any{
				"return": cfg.Env("payment.paypal.url.return", ""),
				"cancel": cfg.Env("payment.paypal.url.cancel", ""),
			},
		},
	})
}
