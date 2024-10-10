package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("app", map[string]any{
		"title":  cfg.Env("app.title", "UPER"),
		"domain": cfg.Env("app.domain", "http://127.0.0.1:9600"),
	})
}

func Boot() {

}
