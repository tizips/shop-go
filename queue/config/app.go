package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("app", map[string]any{
		"name":  cfg.Env("app.name", "UPER"),
		"title": cfg.Env("app.title", "UPER"),
	})
}

func Boot() {

}
