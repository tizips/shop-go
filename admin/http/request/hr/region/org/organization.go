package org

import "github.com/herhe-com/framework/contracts/http/request"

type ToOrganizationByPaginate struct {
	Keyword  string `json:"keyword" form:"keyword" valid:"omitempty,max=32" label:"关键词"`
	Parent   string `json:"parent" form:"parent" valid:"omitempty,snowflake" label:"父级"`
	Platform uint16 `query:"platform" valid:"omitempty,oneof=777 888 999" label:"类型"`

	request.Paginate
}

type DoOrganizationOfCreate struct {
	Parent      string `json:"parent" form:"parent" valid:"omitempty,snowflake" label:"父级"`
	Platform    uint16 `json:"platform" form:"platform" valid:"required,oneof=888 999" label:"类型"`
	Name        string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Brand       uint   `json:"brand" form:"brand" valid:"required_if=Platform 999,omitempty,gt=0" label:"品牌"`
	ValidStart  string `json:"valid_start" form:"valid_start" valid:"required,datetime=2006-01-02" label:"有效期：开始"`
	ValidEnd    string `json:"valid_end" form:"valid_end" valid:"required,datetime=2006-01-02" label:"有效期：结束"`
	User        string `json:"user" form:"user" valid:"required,max=32" label:"联系人"`
	Telephone   string `json:"telephone" form:"telephone" valid:"required,max=32" label:"联系电话"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`

	request.Enable
}

type DoOrganizationOfUpdate struct {
	Name        string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Brand       uint   `json:"brand" form:"brand" valid:"omitempty,gt=0" label:"品牌"`
	ValidStart  string `json:"valid_start" form:"valid_start" valid:"required,datetime=2006-01-02" label:"有效期：开始"`
	ValidEnd    string `json:"valid_end" form:"valid_end" valid:"required,datetime=2006-01-02" label:"有效期：结束"`
	User        string `json:"user" form:"user" valid:"required,max=32" label:"联系人"`
	Telephone   string `json:"telephone" form:"telephone" valid:"required,max=32" label:"联系电话"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`

	request.IDOfSnowflake
	request.Enable
}

type DoOrganizationOfDelete struct {
	request.IDOfSnowflake
}

type DoOrganizationOfEnable struct {
	request.IDOfSnowflake
	request.Enable
}

type DoOrganizationOfEnter struct {
	request.IDOfSnowflake
}

type ToOrganizationOfInformation struct {
	request.IDOfSnowflake
}
