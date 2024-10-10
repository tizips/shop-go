package shop

import "github.com/herhe-com/framework/contracts/http/request"

type ToBlogOfPaginate struct {
	request.Paginate
}

type ToBlogOfInformation struct {
	request.IDOfSnowflake
}
