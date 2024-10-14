package order

import "github.com/herhe-com/framework/contracts/http/request"

type ToServiceOfPaginate struct {
	ID     string `query:"id" valid:"omitempty,snowflake" label:"售后单号"`
	Order  string `query:"order" valid:"omitempty,snowflake" label:"订单号"`
	Status string `query:"status" valid:"omitempty,oneof=pending user org confirm_user confirm_org finish closed" label:"状态"`

	request.Paginate
}

type ToServiceOfLogs struct {
	request.IDOfSnowflake
}

type DoServiceOfHandle struct {
	Result string `json:"result" form:"result" valid:"required,oneof=agree refuse" label:"处理结果"`
	Remark string `json:"remark" form:"remark" valid:"required_if=Result refuse,max=255"`

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

type DoServiceOfClosed struct {
	Remark string `json:"remark" form:"remark" valid:"required,max=255" label:"备注"`

	request.IDOfSnowflake
}
