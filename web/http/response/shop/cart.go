package shop

type ToCarts struct {
	ID             uint     `json:"id"`
	Product        string   `json:"product"`
	Name           string   `json:"name"`
	Picture        string   `json:"picture"`
	Code           string   `json:"code,omitempty"`
	Specifications []string `json:"specifications,omitempty"`
	Price          uint     `json:"price"`
	Quantity       uint     `json:"quantity"`
	IsInvalid      uint8    `json:"is_invalid"`
	CreateAt       string   `json:"create_at"`
}
