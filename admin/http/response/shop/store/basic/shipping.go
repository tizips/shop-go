package basic

type ToShippingOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Money     uint   `json:"money"`
	Query     string `json:"query"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToShippingOfOpening struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
