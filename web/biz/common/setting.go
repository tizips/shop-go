package common

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/support/util"
	"github.com/redis/go-redis/v9"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/common"
	"time"
)

func ToSetting(c context.Context, ctx *app.RequestContext) {

	var request req.ToSetting

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := make(map[string]string)

	key := util.Keys("setting", request.Module, *(auth.Organization(ctx)))

	result, err := facades.Redis.HGetAll(c, key).Result()

	if err == nil && len(result) > 0 {
		responses = result
	} else if errors.Is(err, redis.Nil) || len(result) == 0 { // 从数据库获取

		var settings []model.ComSetting

		facades.Gorm.Find(&settings, "`module`=?", request.Module)

		responses = make(map[string]string, len(settings))

		for _, item := range settings {
			responses[item.Key] = item.Val
		}

		if _, err = facades.Redis.HSet(c, key, responses).Result(); err == nil {
			facades.Redis.Expire(c, key, time.Hour*24)
		}
	}

	http.Success(ctx, responses)
}
