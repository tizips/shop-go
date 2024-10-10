package commodity

import "github.com/herhe-com/framework/contracts/http/request"

type DoCategoryOfCreate struct {
	Parent uint   `json:"parent" form:"parent" valid:"omitempty,gte=0" label:"父级"`
	Name   string `json:"name" form:"name" valid:"required,max=64" label:"名称"`

	request.Order
	request.Enable
}

type DoCategoryOfUpdate struct {
	Name string `json:"name" form:"name" valid:"required,max=64" label:"名称"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoCategoryOfDelete struct {
	request.IDOfUint
}
