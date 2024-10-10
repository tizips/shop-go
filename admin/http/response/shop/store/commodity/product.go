package commodity

type ToProductOfPaginate struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Price          uint   `json:"price"`
	OriginPrice    uint   `json:"origin_price"`
	IsHot          uint8  `json:"is_hot"`
	IsRecommend    uint8  `json:"is_recommend"`
	IsMultiple     uint8  `json:"is_multiple"`
	IsFreeShipping uint8  `json:"is_free_shipping"`
	IsFreeze       uint8  `json:"is_freeze"`
	IsEnable       uint8  `json:"is_enable"`
	CreatedAt      string `json:"created_at"`
}

type ToProductOfInformation struct {
	ID             string                                `json:"id"`
	Categories     []uint                                `json:"categories"`
	Name           string                                `json:"name"`
	Summary        string                                `json:"summary"`
	Pictures       []string                              `json:"pictures"`
	Title          string                                `json:"title"`
	Keyword        string                                `json:"keyword"`
	Description    string                                `json:"description"`
	SKU            *ToProductOfSKU                       `json:"sku,omitempty"`
	Information    string                                `json:"information"`
	Attributes     []ToProductOfInformationWithAttribute `json:"attributes"`
	IsHot          uint8                                 `json:"is_hot"`
	IsRecommend    uint8                                 `json:"is_recommend"`
	IsMultiple     uint8                                 `json:"is_multiple"`
	IsFreeShipping uint8                                 `json:"is_free_shipping"`
	IsFreeze       uint8                                 `json:"is_freeze"`
	IsEnable       uint8                                 `json:"is_enable"`
	CreatedAt      string                                `json:"created_at"`
}

type ToProductOfSpecification struct {
	ID             string                              `json:"id"`
	Specifications []ToProductOfSpecificationWithGroup `json:"specifications"`
	SKUS           []ToProductOfSKU                    `json:"skus"`
}

type ToProductOfSpecificationWithGroup struct {
	ID        uint                                   `json:"id"`
	Name      string                                 `json:"name"`
	Children  []ToProductOfSpecificationWithChildren `json:"children"`
	CreatedAt string                                 `json:"created_at"`
}

type ToProductOfSpecificationWithChildren struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type ToProductOfSKU struct {
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Price       uint   `json:"price"`
	OriginPrice uint   `json:"origin_price"`
	CostPrice   uint   `json:"cost_price"`
	Stock       uint   `json:"stock"`
	Warn        uint   `json:"warn"`
	Picture     string `json:"picture,omitempty"`
	IsDefault   uint8  `json:"is_default"`
}

type ToProductOfInformationWithAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
