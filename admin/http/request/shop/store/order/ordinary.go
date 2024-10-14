package order

import "github.com/herhe-com/framework/contracts/http/request"

type ToOrdinaryOfPaginate struct {
	ID     string `query:"id" valid:"omitempty,snowflake" label:"订单号"`
	Status string `query:"status" valid:"omitempty,oneof=pay shipment receipt completed closed" label:"状态"`

	request.Paginate
}

type DoOrdinaryOfShipment struct {
	Company string `json:"company" form:"company" valid:"required,max=64" label:"快递公司"`
	No      string `json:"no" form:"no" valid:"required,max=64" label:"快递单号"`
	Remark  string `json:"remark" form:"remark" valid:"omitempty,max=255" label:"备注"`

	request.IDOfSnowflake
}

type ToOrdinaryOfAddress struct {
	request.IDOfSnowflake
}

type ToOrdinaryOfInformation struct {
	request.IDOfSnowflake
}

type DoOrdinaryOfRemark struct {
	Remark string `json:"remark" form:"remark" valid:"required,max=255" label:"备注"`

	request.IDOfSnowflake
}

type ToOrdinaryOfLogs struct {
	request.IDOfSnowflake
}
