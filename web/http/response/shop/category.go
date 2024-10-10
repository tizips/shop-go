package shop

type ToCategories struct {
	ID       uint           `json:"id"`
	Level    string         `json:"level"`
	Name     string         `json:"name"`
	Children []ToCategories `json:"children,omitempty"`
}
