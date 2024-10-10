package common

import "github.com/herhe-com/framework/contracts/http/request"

type ToRoleByPaginate struct {
	request.Paginate
}

type DoRoleByCreate struct {
	Name        string   `json:"name" form:"name" valid:"required,max=32" label:"名称"`
	Permissions []string `json:"permissions" form:"permissions[]" valid:"required,min=1,unique" label:"权限"`
	Summary     string   `json:"summary" form:"summary" valid:"omitempty,max=255" label:"简介"`
}

type DoRoleByUpdate struct {
	Name        string   `json:"name" form:"name" valid:"required,max=32" label:"名称"`
	Permissions []string `json:"permissions" form:"permissions" valid:"required,min=1,unique" label:"权限"`
	Summary     string   `json:"summary" form:"summary" valid:"omitempty,max=255" label:"简介"`

	request.IDOfUint
}
