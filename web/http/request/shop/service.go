package shop

import "github.com/herhe-com/framework/contracts/http/request"

type ToServiceOfPaginate struct {
	request.Paginate
}

type ToServiceOfInformation struct {
	request.IDOfSnowflake
}

type DoServiceOfCancel struct {
	request.IDOfSnowflake
}

type DoServiceOfShipment struct {
	Company string `json:"company" form:"company" valid:"required,max=64" label:"快递公司"`
	No      string `json:"no" form:"no" valid:"required,max=64" label:"快递单号"`
	Remark  string `json:"remark" form:"remark" valid:"omitempty,max=255" label:"快递备注"`

	request.IDOfSnowflake
}

type DoServiceOfFinish struct {
	request.IDOfSnowflake
}
