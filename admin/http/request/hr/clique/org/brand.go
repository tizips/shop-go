package org

import "github.com/herhe-com/framework/contracts/http/request"

type DoBrandOfCreate struct {
	Name string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Logo string `json:"logo" form:"logo" valid:"required,max=255,http_url" label:"LOGO"`

	request.Order
}

type DoBrandOfUpdate struct {
	Name string `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	Logo string `json:"logo" form:"logo" valid:"required,max=255,http_url" label:"LOGO"`

	request.IDOfUint
	request.Order
}

type DoBrandOfDelete struct {
	request.IDOfUint
}
