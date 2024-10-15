package shop

import "github.com/herhe-com/framework/contracts/http/request"

type DoOrder struct {
	Shipping   uint   `json:"shipping" form:"shipping" valid:"required,gte=0" label:"Shipping"`
	Payment    uint   `json:"payment" form:"payment" valid:"required,gt=0" label:"Payment"`
	Coupon     string `json:"coupon" form:"coupon" valid:"omitempty,snowflake" label:"Coupon"`
	FirstName  string `json:"first_name" form:"first_name" valid:"required,max=64" label:"FirstName"`
	LastName   string `json:"last_name" form:"last_name" valid:"required,max=64" label:"LastName"`
	Company    string `json:"company" form:"company" valid:"omitempty,max=255" label:"Company Name"`
	Country    string `json:"country" form:"country" valid:"required,max=255" label:"Country"`
	Prefecture string `json:"prefecture" form:"prefecture" valid:"required,max=255" label:"Prefecture"`
	City       string `json:"city" form:"city" valid:"required,max=255" label:"City"`
	Street     string `json:"street" form:"street" valid:"required,max=255" label:"Street"`
	Detail     string `json:"detail" form:"detail" valid:"omitempty,max=255" label:"Detail"`
	Postcode   string `json:"postcode" form:"postcode" valid:"required,max=64" label:"Postcode"`
	Phone      string `json:"phone" form:"phone" valid:"required,max=64" label:"Phone"`
	Email      string `json:"email" form:"email" valid:"required,max=120,email" label:"Email"`
	Remark     string `json:"remark" form:"remark" valid:"omitempty,max=1000" label:"Remark"`
}

type ToOrderOfPaginate struct {
	Status      string `query:"status" valid:"omitempty,oneof=pay shipment receipt evaluate" label:"状态"`
	IsAppraisal uint8  `query:"is_appraisal" valid:"omitempty,oneof=1 2" label:"是否评价"`

	request.Paginate
}

type ToOrderOfInformation struct {
	request.IDOfSnowflake
}

type DoOrderOfReceived struct {
	request.IDOfSnowflake
}

type DoOrderOfCompleted struct {
	request.IDOfSnowflake
}

type DoOrderOfService struct {
	ID       string                       `json:"id" form:"id" valid:"required,snowflake" label:"订单号"`
	Type     string                       `json:"type" form:"type" valid:"required,oneof=un_receipt refund exchange" label:"类型"`
	Details  []DoOrderOfServiceWithDetail `json:"details" form:"details[]" valid:"omitempty,unique=ID,dive" label:"明细列表"`
	Reason   string                       `json:"reason" form:"reason" valid:"omitempty,max=255" label:"原因"`
	Pictures []string                     `json:"pictures" form:"pictures[]" valid:"omitempty,max=8,unique,dive,max=255,http_url" label:"图片"`
}

type DoOrderOfServiceWithDetail struct {
	ID       string `json:"id" form:"id" valid:"required,snowflake" label:"明细ID"`
	Quantity uint   `json:"quantity" form:"quantity" valid:"required,gt=0" label:"售后数量"`
}
