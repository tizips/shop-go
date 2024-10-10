package shop

type ToPage struct {
	Code string `query:"code" valid:"required,max=64" label:"Code"`
}
