package pay

import "github.com/herhe-com/framework/contracts/http/request"

type DoChannelOfCreate struct {
	Name    string             `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Channel string             `json:"channel" form:"channel" valid:"required,oneof=paypal" label:"渠道"`
	Key     string             `json:"key" form:"key" valid:"required,max=255" label:"KEY"`
	Secret  string             `json:"secret" form:"secret" valid:"required,max=255" label:"SECRET"`
	IsDebug uint8              `json:"is_debug" form:"is_debug" valid:"required,oneof=1 2" label:"调试模式"`
	PayPal  *DoChannelOfPayPal `json:"paypal" form:"paypal" valid:"required_if=Channel paypal" label:"PayPal"`

	request.Order
	request.Enable
}

type DoChannelOfUpdate struct {
	Name    string             `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Channel string             `json:"channel" form:"channel" valid:"required,oneof=paypal" label:"渠道"`
	Key     string             `json:"key" form:"key" valid:"required,max=255" label:"KEY"`
	Secret  string             `json:"secret" form:"secret" valid:"required,max=255" label:"SECRET"`
	IsDebug uint8              `json:"is_debug" form:"is_debug" valid:"required,oneof=1 2" label:"调试模式"`
	PayPal  *DoChannelOfPayPal `json:"paypal" form:"paypal" valid:"required_if=Channel paypal" label:"PayPal"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoChannelOfPayPal struct {
	ReturnURL string `json:"return_url" form:"return_url" valid:"required,max=255,http_url" label:"成功跳转链接"`
	CancelURL string `json:"cancel_url" form:"cancel_url" valid:"required,max=255,http_url" label:"取消跳转链接"`
}

type DoChannelOfEnable struct {
	request.IDOfUint
	request.Enable
}

type DoChannelOfDelete struct {
	request.IDOfUint
}

type ToChannelOfPaginate struct {
	request.Paginate
}

type ToChannelOfInformation struct {
	request.IDOfUint
}
