package shop

import "github.com/herhe-com/framework/contracts/http/request"

type ToProductOfPaginate struct {
	Category uint `query:"category" valid:"omitempty,gte=0" label:"Category"`

	request.Paginate
}

type ToProductOfInformation struct {
	request.IDOfSnowflake
}

type ToProductOfSpecification struct {
	request.IDOfSnowflake
}
