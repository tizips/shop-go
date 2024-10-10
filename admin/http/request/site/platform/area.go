package platform

import "github.com/herhe-com/framework/contracts/http/request"

type ToAreaOfPaginate struct {
	Parent uint `query:"parent" valid:"omitempty,gte=0" label:"父级"`

	request.Paginate
}

type DoAreaOfCreate struct {
	Parent uint   `json:"parent" form:"parent" valid:"omitempty,gte=0" label:"父级"`
	Name   string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Code   string `json:"code" form:"code" valid:"required,max=16" label:"代码"`
}

type DoAreaOfUpdate struct {
	Name string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Code string `json:"code" form:"code" valid:"required,max=16" label:"代码"`

	request.IDOfUint
}

type DoAreaOfDelete struct {
	request.IDOfUint
}
