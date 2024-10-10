package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-module/carbon/v2"
	"github.com/golang-module/dongle"
	"github.com/herhe-com/framework/cache"
	"github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/support/util"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"net/url"
	"project.io/shop/model"
	"strings"
	"time"
)

func Authorize() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {

		if ctx.GetBool("API_NOT_AUTHORIZE") {
			ctx.Next(c)
			return
		}

		if total, err := limiter(c, ctx); err != nil && !errors.Is(err, redis.Nil) || total >= 3 {
			ctx.Abort()
			http.Fail(ctx, "The operation is too frequent!")
			return
		}

		var request authorize

		if err := ctx.BindAndValidate(&request); err != nil {
			ctx.Abort()
			http.BadRequest(ctx, err)
			return
		}

		timer := carbon.Parse(request.Time)

		if timer.IsZero() {
			ctx.Abort()
			http.Fail(ctx, "Datetime parsing failed")
			return
		}

		now := carbon.Now()

		if timer.Gt(now) || timer.Lt(now.SubHours(2)) {
			ctx.Abort()
			http.Fail(ctx, "Datetime is not available")
			return
		}

		params := url.Values{}

		ctx.QueryArgs().VisitAll(func(key []byte, value []byte) {
			params.Add(string(key), string(value))
		})

		params.Del("sign")

		for index, item := range params {

			values := lo.Filter(item, func(value string, key int) bool { return value != "" })

			params.Del(index)

			for _, value := range values {
				params.Add(index, value)
			}
		}

		var secret model.SysSecret

		if err := cache.FindByID(c, &secret, request.Key); err != nil {
			ctx.Abort()
			http.Forbidden(ctx)
			return
		}

		if secret.OrganizationID == nil {
			ctx.Abort()
			http.Forbidden(ctx)
			return
		}

		params.Set("secret", secret.Secret)

		sign := dongle.Encrypt.FromString(params.Encode()).ByMd5().ToHexString()

		sign = strings.ToUpper(sign)

		spew.Dump(sign)

		if sign != request.Sign {

			limit(c, ctx)

			ctx.Abort()
			http.Fail(ctx, "Signature verification failed")
			return
		}

		ctx.Set(auth.ContextOfPlatform, secret.Platform)
		ctx.Set(auth.ContextOfOrganization, secret.OrganizationID)

		ctx.Next(c)
	}
}

func NotAuthorize() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {

		ctx.Set("API_NOT_AUTHORIZE", true)

		ctx.Next(c)
	}
}

type authorize struct {
	Key  string `query:"key" valid:"required,min=1,max=64" label:"KEY"`
	Time string `query:"time" valid:"required,datetime=2006-01-02 15:04:05" label:"Datetime"`
	Sign string `query:"sign" valid:"required,min=1,max=255" label:"Sign"`
}

func limit(c context.Context, ctx *app.RequestContext) {

	k := key(ctx)

	total, err := facades.Redis.Incr(c, k).Result()

	if err != nil || total > 3 {
		ctx.Abort()
		http.Fail(ctx, "The operation is too frequent!")
		return
	} else if total == 1 {
		facades.Redis.Expire(c, k, time.Minute)
	}
}

func limiter(c context.Context, ctx *app.RequestContext) (total int, err error) {

	k := key(ctx)

	return facades.Redis.Get(c, k).Int()
}

func key(ctx *app.RequestContext) string {
	return util.Keys("limit", "authorize", ctx.ClientIP())
}
