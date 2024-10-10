package shop

type DoNotifyOfPaypal struct {
	PaymentId string `query:"paymentId" valid:"required" label:"paymentId"`
	Token     string `query:"token" valid:"required" label:"token"`
	PayerID   string `query:"PayerID" valid:"required" label:"PayerID"`
}
