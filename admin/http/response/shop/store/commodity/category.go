package commodity

type ToCategories struct {
	ID       uint           `json:"id"`
	Level    string         `json:"level"`
	Name     string         `json:"name"`
	Icon     string         `json:"icon,omitempty"`
	Order    uint8          `json:"order"`
	IsEnable uint8          `json:"is_enable"`
	Children []ToCategories `json:"children,omitempty"`
}

type ToCategoryOfOpening struct {
	ID       uint                  `json:"id"`
	Level    string                `json:"level"`
	Name     string                `json:"name"`
	Icon     string                `json:"icon,omitempty"`
	Children []ToCategoryOfOpening `json:"children,omitempty"`
}
