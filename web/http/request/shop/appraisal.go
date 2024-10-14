package shop

import "github.com/herhe-com/framework/contracts/http/request"

type DoAppraisalOfCreate struct {
	StarProduct  uint8    `json:"star_product" form:"star_product" valid:"required,gte=1,lte=5" label:"商品评价"`
	StarShipment uint8    `json:"star_shipment" form:"star_shipment" valid:"required,gte=1,lte=5" label:"物流评价"`
	Remark       string   `json:"remark" form:"remark" valid:"required,max=255" label:"评价内容"`
	Pictures     []string `json:"pictures" form:"pictures[]" valid:"omitempty,max=8,unique,dive,max=255,http_url" label:"图片"`

	request.IDOfSnowflake
}

type DoAppraisalOfDelete struct {
	request.IDOfUint
}

type ToAppraisalOfPaginate struct {
	request.Paginate
}
