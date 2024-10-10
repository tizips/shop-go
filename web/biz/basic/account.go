package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	"project.io/shop/model"
	res "project.io/shop/web/http/response/basic"
)

func ToAccountOfInformation(c context.Context, ctx *app.RequestContext) {

	var user model.ShpUser

	fu := facades.Gorm.First(&user, "`id`=?", auth.ID(ctx))

	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfInformation{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.ToDateTimeString(),
	}

	http.Success(ctx, responses)
}

func DoAccount(c context.Context, ctx *app.RequestContext) {

	//var request req.DoAccount
	//
	//if err := ctx.BindAndValidate(&request); err != nil {
	//	http.BadRequest(ctx, err)
	//	return
	//}
	//
	//updates := make(map[string]any)
	//
	//if request.Mobile != "" {
	//	updates["mobile"] = request.Mobile
	//}
	//
	//if request.Email != "" {
	//	updates["email"] = request.Email
	//}
	//
	//if request.Password != "" {
	//	updates["password"] = auth.Password(request.Password)
	//}
	//
	//if len(updates) > 0 {
	//
	//	if result := facades.Gorm.Model(&model.SysUser{}).Where("`id` = ?", auth.ID(ctx)).Updates(updates); result.Error != nil {
	//		http.Fail(ctx, "修改失败：%v", result.Error)
	//		return
	//	}
	//}

	http.Success[any](ctx)
}
