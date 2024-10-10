package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("database", map[string]any{
		"mysql": map[string]any{
			"username": cfg.Env("database.mysql.username", "root"),
			"password": cfg.Env("database.mysql.password", ""),
			"host":     cfg.Env("database.mysql.host", "127.0.0.1"),
			"port":     cfg.Env("database.mysql.port", "3306"),
			"db":       cfg.Env("database.mysql.db", "upper"),
			"charset":  cfg.Env("database.mysql.charset", "utf8mb4_unicode_ci"),
		},
		"redis": map[string]any{
			"host":     cfg.Env("database.redis.host", "127.0.0.1"),
			"password": cfg.Env("database.redis.password", ""),
			"port":     cfg.Env("database.redis.port", "6379"),
			"database": cfg.Env("database.redis.database", 0),
		},
		"migration": map[string]any{
			"dir":   cfg.Env("database.migration.dir", "/migration"),
			"table": cfg.Env("database.migration.table", "sys_migration"),
		},
	})
}
