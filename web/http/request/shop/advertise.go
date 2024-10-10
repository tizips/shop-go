package shop

type ToAdvertises struct {
	Page  string `query:"page" valid:"required,oneof=home" label:"页面"`
	Block string `query:"block" valid:"required,oneof=new_product" label:"广告位"`
}
