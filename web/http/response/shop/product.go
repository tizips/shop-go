package shop

type ToProductOfPaginate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	OriginPrice uint   `json:"origin_price,omitempty"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
}

type ToProductOfHot struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	OriginPrice uint   `json:"origin_price,omitempty"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
}

type ToProductOfRecommended struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	OriginPrice uint   `json:"origin_price,omitempty"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
}

type ToProductOfInformation struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Summary        string      `json:"summary"`
	Pictures       []string    `json:"pictures"`
	Price          uint        `json:"price"`
	OriginPrice    uint        `json:"origin_price,omitempty"`
	Information    string      `json:"information"`
	Attributes     []Attribute `json:"attributes,omitempty"`
	SKU            *SKU        `json:"sku,omitempty"`
	IsMultiple     uint8       `json:"is_multiple"`
	IsFreeShipping uint8       `json:"is_free_shipping"`
	CreateAt       string      `json:"create_at"`
}

type ToProductOfSpecification struct {
	ID             string          `json:"id"`
	Specifications []Specification `json:"specifications"`
	SKUS           []SKU           `json:"skus"`
}

type Specification struct {
	ID      uint            `json:"id"`
	Name    string          `json:"name"`
	Options []Specification `json:"options,omitempty"`
}

type Attribute struct {
	Label string `json:"label"`
	Value string `json:"value,omitempty"`
}

type SKU struct {
	ID          string `json:"id"`
	Code        string `json:"code,omitempty"`
	Price       uint   `json:"price"`
	OriginPrice uint   `json:"origin_price,omitempty"`
	Stock       uint   `json:"stock"`
	Picture     string `json:"picture,omitempty"`
}
