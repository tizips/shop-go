package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoAdvertiseOfCreate struct {
	Page   string `json:"page" form:"page" valid:"required,oneof=home" label:"页面"`
	Block  string `json:"block" form:"block" valid:"required,oneof=new_product" label:"广告位"`
	Title  string `json:"title" form:"title" valid:"required,max=255" label:"标题"`
	Target string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL    string `json:"url" form:"url" valid:"required,max=255,uri" label:"链接"`
	Thumb  string `json:"thumb" form:"thumb" valid:"required,max=255,http_url" label:"图片"`

	request.Order
	request.Enable
}

type DoAdvertiseOfUpdate struct {
	Page   string `json:"page" form:"page" valid:"required,oneof=home" label:"页面"`
	Block  string `json:"block" form:"block" valid:"required,oneof=new_product" label:"广告位"`
	Title  string `json:"title" form:"title" valid:"required,max=255" label:"标题"`
	Target string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL    string `json:"url" form:"url" valid:"required,max=255,uri" label:"链接"`
	Thumb  string `json:"thumb" form:"thumb" valid:"required,max=255,http_url" label:"图片"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoAdvertiseOfDelete struct {
	request.IDOfUint
}

type DoAdvertiseOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToAdvertiseOfPaginate struct {
	Pages string `query:"pages" valid:"omitempty,oneof=home" label:"页面"`
	Block string `query:"block" valid:"omitempty,oneof=new_product" label:"广告位"`

	request.Paginate
}
