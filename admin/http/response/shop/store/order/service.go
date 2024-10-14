package order

type ToServiceOfPaginate struct {
	ID                   string               `json:"id"`
	Order                string               `json:"order"`
	Type                 string               `json:"type"`
	Details              []ToOrderOfDetail    `json:"details"`
	Refund               *ToServiceOfRefund   `json:"refund"`
	Status               string               `json:"status"`
	Result               string               `json:"result,omitempty"`
	Reason               string               `json:"reason"`
	Subtotal             uint                 `json:"subtotal"`
	Shipping             uint                 `json:"shipping"`
	Refunds              uint                 `json:"refunds"`
	Pictures             []string             `json:"pictures,omitempty"`
	ShipmentUser         *ToServiceOfShipment `json:"shipment_user,omitempty"`
	ShipmentOrganization *ToServiceOfShipment `json:"shipment_organization,omitempty"`
	CreatedAt            string               `json:"created_at"`
}

type ToServiceOfLogs struct {
	ID        uint   `json:"id"`
	Action    string `json:"action"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type ToServiceOfRefund struct {
	ID         string `json:"id"`
	No         string `json:"no"`
	Channel    string `json:"channel"`
	Price      uint   `json:"price"`
	Currency   string `json:"currency"`
	RefundedAt string `json:"refunded_at"`
}

type ToServiceOfShipment struct {
	Company string `json:"company"`
	No      string `json:"no"`
	Remark  string `json:"remark,omitempty"`
}
