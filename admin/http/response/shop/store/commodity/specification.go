package commodity

type ToSpecificationOfPaginate struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Label    string   `json:"label"`
	Options  []string `json:"options"`
	IsEnable uint8    `json:"is_enable"`
}

type ToSpecificationOfOpening struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	Label   string   `json:"label"`
	Options []string `json:"options"`
}
