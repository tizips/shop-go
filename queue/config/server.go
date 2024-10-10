package config

import (
	"github.com/herhe-com/framework/console"
	"github.com/herhe-com/framework/console/consoles"
	cons "github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/service"
	"github.com/herhe-com/framework/database/gorm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/herhe-com/framework/microservice/snowflake"
	"github.com/herhe-com/framework/queue"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("server", map[string]any{
		"name":     cfg.Env("app.name", "UPER"),
		"title":    cfg.Env("app.title", "UPER"),
		"node":     cfg.Env("server.node", 1),
		"debug":    cfg.Env("server.debug", false),
		"location": cfg.Env("server.location", "Asia/Shanghai"),
		"providers": []service.Provider{
			&gorm.ServiceProvider{},
			&redis.ServiceProvider{},
			//&filesystem.ServiceProvider{},
			&snowflake.ServiceProvider{},
			&locker.ServiceProvider{},
			&queue.ServiceProvider{},
			//&validation.ServiceProvider{},
			//&auth.ServiceProvider{},
			&console.ServiceProvider{},
		},
		"consoles": []cons.Provider{
			//&consoles.MigrationProvider{},
			&consoles.ConsumerProvider{},
			//&consoles2.DeveloperProvider{},
		},
	})
}
