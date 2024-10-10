package org

import "github.com/herhe-com/framework/contracts/http/request"

type ToOrganizationOfPaginate struct {
	Keyword  string `query:"keyword" valid:"omitempty,max=20" label:"关键词"`
	Platform uint16 `query:"platform" valid:"omitempty,oneof=777 888 999" label:"类型"`

	request.Paginate
}

type DoOrganizationOfInformation struct {
	request.IDOfSnowflake
}

type DoOrganizationOfCreate struct {
	Platform    uint16 `json:"platform" form:"platform" valid:"required,oneof=777 999" label:"类型"`
	Name        string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	ValidStart  string `json:"valid_start" form:"valid_start" valid:"required,datetime=2006-01-02" label:"有效期：开始"`
	ValidEnd    string `json:"valid_end" form:"valid_end" valid:"required,datetime=2006-01-02" label:"有效期：结束"`
	User        string `json:"user" form:"user" valid:"required,max=32" label:"联系人"`
	Telephone   string `json:"telephone" form:"telephone" valid:"required,max=32" label:"联系电话"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`

	request.Enable
}

type DoOrganizationOfUpdate struct {
	Name        string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
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
