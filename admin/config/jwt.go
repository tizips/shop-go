package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("jwt", map[string]any{
		"secret":   cfg.Env("jwt.secret", ""),
		"leeway":   cfg.Env("jwt.leeway", 3),
		"lifetime": cfg.Env("jwt.lifetime", 60*24),
	})
}
