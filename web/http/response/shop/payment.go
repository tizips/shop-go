package shop

type DoPaymentOfPaypal struct {
	Link string `json:"link"`
}

type ToPaymentOfChannel struct {
	ID      uint   `json:"id"`
	Channel string `json:"channel"`
}
