package basic

import "github.com/herhe-com/framework/contracts/http/request"

type DoShippingOfCreate struct {
	Name  string `form:"name" json:"name" valid:"required,max=64" label:"名称"`
	Money uint   `json:"money" form:"money" valid:"omitempty,gte=0" label:"运费"`
	Query string `json:"query" form:"query" valid:"omitempty,max=255" label:"查询地址"`

	request.Order
	request.Enable
}

type DoShippingOfUpdate struct {
	Name  string `form:"name" json:"name" valid:"required,max=64" label:"名称"`
	Money uint   `json:"money" form:"money" valid:"omitempty,gte=0" label:"运费"`
	Query string `json:"query" form:"query" valid:"omitempty,max=255" label:"查询地址"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoShippingOfDelete struct {
	request.IDOfUint
}

type DoShippingOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToShippingOfPaginate struct {
	request.Paginate
}
