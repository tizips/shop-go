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
