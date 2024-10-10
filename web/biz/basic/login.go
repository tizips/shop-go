package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	"project.io/shop/web/constants"
	req "project.io/shop/web/http/request/basic"
	res "project.io/shop/web/http/response/basic"
)

func DoLoginOfAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.ShpUser

	fu := facades.Gorm.Scopes(scope.Platform(ctx)).First(&user, "`username`=?", request.Username)

	if fu.Error != nil {
		http.Fail(ctx, "Incorrect username or password.")
		return
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		http.Fail(ctx, "Incorrect username or password.")
		return
	}

	lifetime := facades.Cfg.GetInt("jwt.lifetime")

	var err error
	var token string

	if token, err = auth.NewJWToken(constants.JwtOfIssuerWithWeb, user.ID, user.OrganizationID, user.CliqueID, lifetime, true, nil, auth.Platform(ctx)); err != nil {
		http.Login(ctx)
		return
	}

	responses := res.DoLogin{
		Token:    token,
		Lifetime: lifetime,
	}

	http.Success(ctx, responses)
}

func DoLoginOfEMail(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfEMail

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.ShpUser

	fu := facades.Gorm.Scopes(scope.Platform(ctx)).First(&user, "`email`=?", request.EMail)

	if fu.Error != nil {
		http.Fail(ctx, "Incorrect email or password.")
		return
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		http.Fail(ctx, "Incorrect email or password.")
		return
	}

	lifetime := facades.Cfg.GetInt("jwt.lifetime")

	var err error
	var token string

	if token, err = auth.NewJWToken(constants.JwtOfIssuerWithWeb, user.ID, user.OrganizationID, user.CliqueID, lifetime, true, nil, auth.Platform(ctx)); err != nil {
		http.Login(ctx)
		return
	}

	responses := res.DoLogin{
		Token:    token,
		Lifetime: lifetime,
	}

	http.Success(ctx, responses)
}

func DoLoginOfOut(c context.Context, ctx *app.RequestContext) {

	if ok, _ := auth.BlacklistOfJwtValue(c, ctx); !ok {
		http.Fail(ctx, "退出失败，请稍后重试")
		return
	}

	http.Success[any](ctx)
}
