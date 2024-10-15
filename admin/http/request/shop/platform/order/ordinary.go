package order

import "github.com/herhe-com/framework/contracts/http/request"

type ToOrdinaryOfPaginate struct {
	ID     string `query:"id" valid:"omitempty,snowflake" label:"订单号"`
	Status string `query:"status" valid:"omitempty,oneof=pay shipment receipt completed closed" label:"状态"`

	request.Paginate
}

type ToOrdinaryOfAddress struct {
	request.IDOfSnowflake
}

type ToOrdinaryOfInformation struct {
	request.IDOfSnowflake
}

type ToOrdinaryOfLogs struct {
	request.IDOfSnowflake
}
