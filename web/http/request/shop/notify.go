package shop

type DoNotifyOfPaypal struct {
	Token   string `query:"token" valid:"required" label:"token"`
	PayerID string `query:"PayerID" valid:"required" label:"PayerID"`
}
