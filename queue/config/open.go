package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("open", map[string]any{
		"wechat": map[string]any{
			"mini_program": map[string]any{
				"app_id":     cfg.Env("open.wechat.mini_program.app_id", ""),
				"app_secret": cfg.Env("open.wechat.mini_program.app_secret", ""),
				"cache":      cfg.Env("open.wechat.mini_program.cache", ""),
			},
			"payment": map[string]any{
				"domain":      cfg.Env("open.wechat.payment.domain", ""),
				"description": cfg.Env("open.wechat.payment.description", "购买商品"),
				"app_id":      cfg.Env("open.wechat.payment.app_id", ""),
				"mch_id":      cfg.Env("open.wechat.payment.mch_id", ""),
				"key":         cfg.Env("open.wechat.payment.key", ""),
				"cert_path":   cfg.Env("open.wechat.payment.cert_path", ""),
				"key_path":    cfg.Env("open.wechat.payment.key_path", ""),
				"serial_no":   cfg.Env("open.wechat.payment.serial_no", ""),
			},
		},
	})
}
