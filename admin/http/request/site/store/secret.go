package store

import "github.com/herhe-com/framework/contracts/http/request"

type ToSecretOfPaginate struct {
	request.Paginate
}

type DoSecretOfCreate struct {
	Name string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
}

type DoSecretOfDelete struct {
	request.IDOfSnowflake
}
