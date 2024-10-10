package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoBannerOfCreate struct {
	Name        string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Description string `form:"description" json:"description" valid:"omitempty,max=1000" label:"描述"`
	Button      string `form:"button" json:"button" valid:"required,max=120" label:"按钮"`
	Picture     string `json:"picture" form:"picture" valid:"required,url,max=255" label:"图片"`
	Target      string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL         string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.Order
	request.Enable
}

type DoBannerOfUpdate struct {
	Name        string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Description string `form:"description" json:"description" valid:"omitempty,max=1000" label:"描述"`
	Button      string `form:"button" json:"button" valid:"required,max=120" label:"按钮"`
	Picture     string `json:"picture" form:"picture" valid:"required,url,max=255" label:"图片"`
	Target      string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL         string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoBannerOfDelete struct {
	request.IDOfUint
}

type DoBannerOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToBannerOfPaginate struct {
	request.Paginate
}
