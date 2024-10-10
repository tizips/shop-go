package commodity

import "github.com/herhe-com/framework/contracts/http/request"

type DoSpecificationOfCreate struct {
	Name    string   `json:"name" form:"name" valid:"required,max=32" label:"名称"`
	Label   string   `json:"label" form:"label" valid:"required,max=32" label:"标签"`
	Options []string `json:"options" form:"options" valid:"required,min=1,max=20,unique,dive,required,max=32" label:"选项"`

	request.Enable
}

type DoSpecificationOfUpdate struct {
	Name    string   `json:"name" form:"name" valid:"required,max=32" label:"名称"`
	Label   string   `json:"label" form:"label" valid:"required,max=32" label:"标签"`
	Options []string `json:"options" form:"options" valid:"required,min=1,max=20,unique,dive,required,max=32" label:"选项"`

	request.IDOfUint
	request.Enable
}

type DoSpecificationOfDelete struct {
	request.IDOfUint
}

type DoSpecificationOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToSpecificationOfPaginate struct {
	request.Paginate
}
