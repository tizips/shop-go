package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoPageOfCreate struct {
	Code        string `json:"code" form:"code" valid:"required,max=64" label:"CODE"`
	IsSystem    uint8  `json:"is_system" form:"is_system" valid:"omitempty,oneof=1 2" label:"是否系统内置"`
	Name        string `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"详情"`
}

type DoPageOfUpdate struct {
	IsSystem    uint8  `json:"is_system" form:"is_system" valid:"omitempty,oneof=1 2" label:"是否系统内置"`
	Name        string `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"详情"`

	request.IDOfUint
}

type DoPageOfDelete struct {
	request.IDOfUint
}

type ToPageOfPaginate struct {
	IsSystem uint8 `query:"is_system" valid:"omitempty,oneof=1 2" label:"是否系统内置"`

	request.Paginate
}

type ToPageOfInformation struct {
	request.IDOfUint
}
