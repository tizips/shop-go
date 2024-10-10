package shop

type ToSEO struct {
	Channel string `json:"channel" form:"channel" valid:"required,oneof=product page category blog" label:"Channel"`
	ID      string `json:"id" form:"id" valid:"required,max=64" label:"ID"`
}
