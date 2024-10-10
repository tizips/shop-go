package shop

import "github.com/herhe-com/framework/contracts/http/request"

type DoCartOfCreate struct {
	SKU      string `json:"sku" form:"sku" valid:"required,snowflake" label:"SKU"`
	Quantity uint   `json:"quantity" form:"quantity" valid:"required,gt=0,lte=999" label:"Quantity"`
}

type DoCartOfUpdate struct {
	Quantity uint `json:"quantity" form:"quantity" valid:"required,gt=0,lte=99" label:"Quantity"`

	request.IDOfUint
}

type DoCartOfDelete struct {
	request.IDOfUint
}

type DoCartOfDeletes struct {
	IDS []uint `json:"ids" form:"ids[]" valid:"required,min=1,unique,dive,gt=0" label:"ID"`
}
