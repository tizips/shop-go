package platform

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"gorm.io/gorm"
	req "project.io/shop/admin/http/request/site/platform"
	res "project.io/shop/admin/http/response/site/platform"
	"project.io/shop/model"
)

func ToAreaOfPaginate(ctx context.Context, c *app.RequestContext) {

	var request req.ToAreaOfPaginate

	if err := c.BindAndValidate(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	responses := response.Paginate[res.ToAreaOfPaginate]{
		Size: request.GetSize(),
		Page: request.GetPage(),
	}

	tx := facades.Gorm.Where("`parent_id`=?", request.Parent)

	tx.Model(&model.SysArea{}).Count(&responses.Total)

	if responses.Total > 0 {

		var areas []model.SysArea

		tx.
			Order("`id` asc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&areas)

		responses.Data = make([]res.ToAreaOfPaginate, len(areas))

		for index, item := range areas {
			responses.Data[index] = res.ToAreaOfPaginate{
				ID:        item.ID,
				Level:     item.Level,
				Name:      item.Name,
				Code:      item.Code,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(c, responses)
}

func DoAreaOfCreate(ctx context.Context, c *app.RequestContext) {

	var request req.DoAreaOfCreate

	if err := c.BindAndValidate(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var level = model.SysAreaOfLevelLv1

	if request.Parent > 0 {

		var parent model.SysArea

		fp := facades.Gorm.First(&parent, "`id`=?", request.Parent)

		if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
			http.NotFound(c, "未找到该父级")
			return
		} else if fp.Error != nil {
			http.Fail(c, "父级查询失败：%v", fp.Error)
			return
		}

		if parent.Level == model.SysAreaOfLevelLv1 {
			level = model.SysAreaOfLevelLv2
		} else if parent.Level == model.SysAreaOfLevelLv2 {
			level = model.SysAreaOfLevelLv3
		} else if parent.Level == model.SysAreaOfLevelLv3 {
			http.Fail(c, "该父级下无法添加管辖区域")
			return
		}
	}

	area := model.SysArea{
		Level:    level,
		ParentID: request.Parent,
		Name:     request.Name,
		Code:     request.Code,
	}

	if result := facades.Gorm.Create(&area); result.Error != nil {
		http.Fail(c, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](c)
}

func DoAreaOfUpdate(ctx context.Context, c *app.RequestContext) {

	var request req.DoAreaOfUpdate

	if err := c.BindAndValidate(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var area model.SysArea

	fb := facades.Gorm.First(&area, "`id`=?", request.ID)

	if errors.Is(fb.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fb.Error != nil {
		http.Fail(c, "数据查询失败：%v", fb.Error)
		return
	}

	area.Name = request.Name
	area.Code = request.Code

	if result := facades.Gorm.Save(&area); result.Error != nil {
		http.Fail(c, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](c)
}

func DoAreaOfDelete(ctx context.Context, c *app.RequestContext) {

	var request req.DoAreaOfDelete

	if err := c.BindAndValidate(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if result := facades.Gorm.Delete(&model.SysArea{}, "`id`=?", request.ID); result.Error != nil {
		http.Fail(c, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](c)
}
