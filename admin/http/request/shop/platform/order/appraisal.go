package order

import "github.com/herhe-com/framework/contracts/http/request"

type ToAppraisalOfPaginate struct {
	Type  string  `query:"type" valid:"omitempty,oneof=product shipment" label:"类型"`
	Stars []uint8 `query:"stars[]" valid:"omitempty,unique,dive,gte=1,lte=5" label:"评分"`
	request.Paginate
}
