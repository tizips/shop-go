package common

import "github.com/herhe-com/framework/contracts/http/request"

type ToUserByPaginate struct {
	request.Paginate
}

type DoUserByCreate struct {
	Username string `json:"username" form:"username" valid:"required,username" label:"用户名"`
	Nickname string `json:"nickname" form:"nickname" valid:"required,max=32" label:"昵称"`
	Mobile   string `json:"mobile" form:"mobile" valid:"omitempty,mobile" label:"手机号"`
	Email    string `json:"email" form:"email" valid:"omitempty,email,max=64" label:"邮箱"`
	Password string `json:"password" form:"password" valid:"required,password" label:"密码"`
	Roles    []uint `json:"roles" form:"roles" valid:"required,min=1,unique" label:"角色"`

	request.Enable
}

type DoUserByUpdate struct {
	request.IDOfSnowflake

	Nickname string `json:"nickname" form:"nickname" valid:"required,max=32" label:"昵称"`
	Mobile   string `json:"mobile" form:"mobile" valid:"omitempty,mobile" label:"手机号"`
	Email    string `json:"email" form:"email" valid:"omitempty,email,max=64" label:"邮箱"`
	Password string `json:"password" form:"password" valid:"omitempty,password" label:"密码"`
	Roles    []uint `json:"roles" form:"roles" valid:"required,min=1,unique" label:"角色"`

	request.Enable
}

type DoUserByEnable struct {
	request.IDOfSnowflake
	request.Enable
}
