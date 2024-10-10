package config

import (
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/filesystem"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("filesystem", map[string]any{
		"driver": cfg.Env("filesystem.driver", filesystem.DriverLocal),
		"qiniu": map[string]any{
			"access": cfg.Env("filesystem.qiniu.access"),
			"secret": cfg.Env("filesystem.qiniu.secret"),
			"bucket": cfg.Env("filesystem.qiniu.bucket"),
			"domain": cfg.Env("filesystem.qiniu.domain"),
			"prefix": cfg.Env("filesystem.qiniu.prefix"),
		},
	})
}
