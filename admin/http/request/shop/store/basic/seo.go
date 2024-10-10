package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoSEOOfCreate struct {
	Code        string `json:"code" form:"code" valid:"required,max=64" label:"CODE"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
}

type DoSEOOfUpdate struct {
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`

	request.IDOfUint
}

type DoSEOOfDelete struct {
	request.IDOfUint
}

type ToSEOOfPaginate struct {
	request.Paginate
}
