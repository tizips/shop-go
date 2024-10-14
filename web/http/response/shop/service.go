package shop

type ToServiceOfPaginate struct {
	ID        string              `json:"id"`
	Type      string              `json:"type"`
	Status    string              `json:"status"`
	Reason    string              `json:"reason"`
	Details   []ToServiceOfDetail `json:"details"`
	Subtotal  uint                `json:"subtotal"`
	Shipping  uint                `json:"shipping"`
	Refund    uint                `json:"refund,omitempty"`
	CreatedAt string              `json:"created_at"`
}

type ToServiceOfOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ToServiceOfDetail struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Picture        string   `json:"picture"`
	Specifications []string `json:"specifications,omitempty"`
	Price          uint     `json:"price"`
	Quantity       uint     `json:"quantity"`
	Refund         uint     `json:"refund,omitempty"`
	CreatedAt      string   `json:"created_at"`
}

type ToServiceOfLog struct {
	Action    string `json:"action"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type ToServiceOfInformation struct {
	ID           string                   `json:"id"`
	Order        string                   `json:"order"`
	Type         string                   `json:"type"`
	Status       string                   `json:"status"`
	Reason       string                   `json:"reason"`
	Pictures     []string                 `json:"pictures,omitempty"`
	Details      []ToServiceOfDetail      `json:"details"`
	Organization *ToServiceOfOrganization `json:"organization"`
	Logs         []ToServiceOfLog         `json:"logs"`
	Subtotal     uint                     `json:"subtotal"`
	Shipping     uint                     `json:"shipping"`
	Refund       uint                     `json:"refund"`
	CreatedAt    string                   `json:"created_at"`
}
