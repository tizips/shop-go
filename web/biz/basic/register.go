package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/basic"
)

func DoRegisterOfEMail(c context.Context, ctx *app.RequestContext) {

	var request req.DoRegisterOfEMail

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64 = 0

	facades.Gorm.Model(&model.ShpUser{}).Where("`email`=?", request.EMail).Count(&total)

	if total > 0 {
		http.Fail(ctx, "This email has already been registered.")
		return
	}

	user := model.ShpUser{
		ID:             facades.Snowflake.Generate().String(),
		Platform:       auth.Platform(ctx),
		CliqueID:       auth.Clique(ctx),
		OrganizationID: auth.Organization(ctx),
		Email:          request.EMail,
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		Password:       auth.Password(request.Password),
	}

	if result := facades.Gorm.Create(&user); result.Error != nil {
		http.Fail(ctx, "Account registration failed, please try again later.")
		return
	}

	http.Success[any](ctx)
}
