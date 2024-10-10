package shop

type DoOrder struct {
	PayID   string `json:"pay_id"`
	Channel string `json:"channel"`
}

type ToOrderOfPaginate struct {
	ID string `json:"id"`
	//Organization ToOrderOfOrganization `json:"organization"`
	Details     []ToOrderOfDetail `json:"details"`
	Prices      uint              `json:"prices"`
	Status      string            `json:"status"`
	IsAppraisal uint8             `json:"is_appraisal"`
	CanService  uint8             `json:"can_service"`
	CreateAt    string            `json:"create_at"`
}

type ToOrderOfOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ToOrderOfDetail struct {
	ID       string  `json:"id"`
	Service  *string `json:"service,omitempty"`
	Name     string  `json:"name"`
	Picture  string  `json:"picture"`
	Price    uint    `json:"price"`
	Quantity uint    `json:"quantity"`
	//Prices         uint     `json:"prices"`
	Specifications []string `json:"specifications,omitempty"`
	Refund         uint     `json:"refund"`
	Returned       uint     `json:"returned"`
	//IsServiced     uint8    `json:"is_serviced"`
}

type ToOrderOfAddress struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Company    string `json:"company"`
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
	Reason    string   `json:"reason"`
	Remark    string   `json:"remark"`
	CreatedAt string   `json:"created_at"`
}

type ToOrderOfLog struct {
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

type ToOrderOfInformation struct {
	ID string `json:"id"`
	//Organization ToOrderOfOrganization `json:"organization"`
	Details      []ToOrderOfDetail  `json:"details"`
	Address      ToOrderOfAddress   `json:"address"`
	Payment      *ToOrderOfPayment  `json:"payment,omitempty"`
	Invoice      *ToOrderOfInvoice  `json:"invoice,omitempty"`
	Shipment     *ToOrderOfShipment `json:"shipment,omitempty"`
	Services     []ToOrderOfService `json:"services,omitempty"`
	Logs         []ToOrderOfLog     `json:"logs"`
	CostShipping uint               `json:"cost_shipping"`
	TotalPrice   uint               `json:"total_price"`
	CouponPrice  uint               `json:"coupon_price"`
	Prices       uint               `json:"prices"`
	Refund       uint               `json:"refund"`
	Status       string             `json:"status"`
	Shipping     string             `json:"shipping"`
	Remark       string             `json:"remark"`
	IsInvoice    uint8              `json:"is_invoice"`
	IsAppraisal  uint8              `json:"is_appraisal"`
	CanService   uint8              `json:"can_service"`
	CreateAt     string             `json:"create_at"`
}

type ToOrderOfPayment struct {
	ID      string `json:"id"`
	Channel string `json:"channel"`
	No      string `json:"no,omitempty"`
	PaidAt  string `json:"paid_at,omitempty"`
}

type ToOrderOfService struct {
	ID        string                       `json:"id"`
	Type      string                       `json:"type"`
	Status    string                       `json:"status"`
	Detail    *string                      `json:"detail,omitempty"`
	Reason    string                       `json:"reason"`
	Refund    uint                         `json:"refund"`
	Details   []ToOrderOfServiceWithDetail `json:"details"`
	CreatedAt string                       `json:"created_at"`
}

type ToOrderOfServiceWithDetail struct {
	ID       string `json:"id"`
	Quantity uint   `json:"quantity"`
}
