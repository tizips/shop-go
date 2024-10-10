package shop

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/database/gorm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"project.io/shop/model"
	req "project.io/shop/web/http/request/shop"
	res "project.io/shop/web/http/response/shop"
)

func ToSEO(c context.Context, ctx *app.RequestContext) {

	var request req.ToSEO

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var seo model.ShpSEO

	tx := facades.Gorm.WithContext(c)

	if lo.Contains([]string{model.ShpSEOForChannelOfProduct, model.ShpSEOForChannelOfCategory, model.ShpSEOForChannelOfBlog}, request.Channel) {
		tx = tx.Scopes(scope.Platform(ctx))
	}

	if request.Channel == model.ShpSEOForChannelOfPage {
		tx = tx.Where("exists (?)", facades.Gorm.
			Model(&model.ShpPage{}).
			Select("1").
			Where(fmt.Sprintf("`%s`.`channel_id`=`%s`.`id` and `%s`.`code`=?", model.TableShpSEO, model.TableShpPage, model.TableShpPage), request.ID),
		)
	} else {
		tx = tx.Where("`channel_id`=?", request.ID)
	}

	fs := tx.First(&seo, "`channel`=?", request.Channel)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该 SEO 信息")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "数据查询失败：%v", fs.Error)
		return
	}

	responses := res.ToSEO{
		Title:       seo.Title,
		Keyword:     seo.Keyword,
		Description: seo.Description,
	}

	http.Success(ctx, responses)
}
