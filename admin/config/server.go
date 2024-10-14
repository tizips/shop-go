package config

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/console"
	"github.com/herhe-com/framework/console/consoles"
	cons "github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/service"
	"github.com/herhe-com/framework/database/gorm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/filesystem"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/herhe-com/framework/microservice/snowflake"
	"github.com/herhe-com/framework/queue"
	"github.com/herhe-com/framework/validation"
	consoles2 "project.io/shop/admin/console"
	"project.io/shop/admin/route"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("server", map[string]any{
		"name":     cfg.Env("server.name", "UPER"),
		"address":  cfg.Env("server.address", "0.0.0.0"),
		"port":     cfg.Env("server.port", "9600"),
		"node":     cfg.Env("server.node", 1),
		"debug":    cfg.Env("server.debug", false),
		"location": cfg.Env("server.location", "Asia/Shanghai"),
		"providers": []service.Provider{
			&gorm.ServiceProvider{},
			&redis.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&snowflake.ServiceProvider{},
			&locker.ServiceProvider{},
			&validation.ServiceProvider{},
			&queue.ServiceProvider{},
			&auth.ServiceProvider{},
			&console.ServiceProvider{},
		},
		"consoles": []cons.Provider{
			&consoles.MigrationProvider{},
			&consoles.ServerProvider{},
			//&consoles2.RoleProvider{},
			&consoles2.DeveloperProvider{},
		},
		"route": route.Router,
		//"handle":  func(server *server.Hertz) {},
		//"options": []config.Option{},
		"middlewares": []app.HandlerFunc{
			middleware.Access(),
		},
	})
}
