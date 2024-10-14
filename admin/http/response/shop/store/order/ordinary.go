package order

type ToOrdinaryOfPaginate struct {
	ID          string            `json:"id"`
	TotalPrice  uint              `json:"total_price"`
	CouponPrice uint              `json:"coupon_price"`
	Prices      uint              `json:"prices"`
	Refund      uint              `json:"refund"`
	Details     []ToOrderOfDetail `json:"details"`
	Payment     *ToOrderOfPayment `json:"payment,omitempty"`
	Shipping    string            `json:"shipping"`
	Status      string            `json:"status"`
	IsPaid      uint8             `json:"is_paid"`
	IsInvoice   uint8             `json:"is_invoice"`
	CreatedAt   string            `json:"created_at"`
}

type ToOrderOfUser struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

type ToOrderOfPayment struct {
	ID       string `json:"id"`
	No       string `json:"no"`
	Channel  string `json:"channel"`
	Money    uint   `json:"money"`
	Currency string `json:"currency"`
	PaidAt   string `json:"paid_at"`
}

type ToOrderOfDetail struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Picture        string            `json:"picture"`
	Specifications []string          `json:"specifications,omitempty"`
	Price          uint              `json:"price"`
	Quantity       uint              `json:"quantity"`
	TotalPrice     uint              `json:"total_price"`
	CouponPrice    uint              `json:"coupon_price"`
	Prices         uint              `json:"prices"`
	Refund         uint              `json:"refund"`
	Returned       uint              `json:"returned"`
	Service        *ToOrderOfService `json:"service,omitempty"`
	IsServiced     uint8             `json:"is_serviced,omitempty"`
}

type ToOrderOfService struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Quantity uint   `json:"quantity"`
}

type ToOrderOfOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ToOrdinaryOfAddress struct {
	Shipping   string `json:"shipping"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Company    string `json:"company,omitempty"`
	Country    string `json:"country"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	Street     string `json:"street"`
	Detail     string `json:"detail"`
	Postcode   string `json:"postcode"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type ToOrderOfInvoice struct {
	Type      string   `json:"type"`
	Header    string   `json:"header"`
	No        string   `json:"no"`
	Bank      string   `json:"bank,omitempty"`
	Card      string   `json:"card,omitempty"`
	Address   string   `json:"address,omitempty"`
	Telephone string   `json:"telephone,omitempty"`
	Status    string   `json:"status"`
	Files     []string `json:"files,omitempty"`
	Reason    string   `json:"reason,omitempty"`
	Remark    string   `json:"remark,omitempty"`
}

type ToOrderOfLogs struct {
	ID        uint   `json:"id"`
	Action    string `json:"action"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type ToOrderOfShipment struct {
	Company   string `json:"company"`
	No        string `json:"no"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
}

type ToOrdinaryOfInformation struct {
	ID           string              `json:"id"`
	Details      []ToOrderOfDetail   `json:"details"`
	Shipping     string              `json:"shipping"`
	Address      ToOrdinaryOfAddress `json:"address"`
	Invoice      *ToOrderOfInvoice   `json:"invoice,omitempty"`
	Shipment     *ToOrderOfShipment  `json:"shipment,omitempty"`
	Payment      *ToOrderOfPayment   `json:"payment,omitempty"`
	CostShipping uint                `json:"cost_shipping"`
	TotalPrice   uint                `json:"total_price"`
	CouponPrice  uint                `json:"coupon_price"`
	Prices       uint                `json:"prices"`
	Refund       uint                `json:"refund"`
	Status       string              `json:"status"`
	Remark       string              `json:"remark"`
	IsInvoice    uint8               `json:"is_invoice"`
	CreateAt     string              `json:"create_at"`
}
