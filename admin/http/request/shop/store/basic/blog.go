package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoBlogOfCreate struct {
	Name        string `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,http_url" label:"图片"`
	Summary     string `json:"summary" form:"summary" valid:"required,max=500" label:"简介"`
	PostedAt    string `json:"posted_at" form:"posted_at" valid:"required,datetime=2006-01-02" label:"发布时间"`
	IsTop       uint8  `json:"is_top" form:"is_top" valid:"omitempty,oneof=1 2" label:"是否置顶"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"详情"`
}

type DoBlogOfUpdate struct {
	Name        string `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,http_url" label:"图片"`
	Summary     string `json:"summary" form:"summary" valid:"required,max=500" label:"简介"`
	PostedAt    string `json:"posted_at" form:"posted_at" valid:"required,datetime=2006-01-02" label:"发布时间"`
	IsTop       uint8  `json:"is_top" form:"is_top" valid:"omitempty,oneof=1 2" label:"是否置顶"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"详情"`

	request.IDOfSnowflake
}

type DoBlogOfDelete struct {
	request.IDOfSnowflake
}

type ToBlogOfPaginate struct {
	IsTop uint8 `query:"is_top" valid:"omitempty,oneof=1 2" label:"是否置顶"`

	request.Paginate
}

type ToBlogOfInformation struct {
	request.IDOfSnowflake
}
