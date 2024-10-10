package basic

type ToAdvertiseOfPaginate struct {
	ID        uint   `json:"id"`
	Page      string `json:"page"`
	Block     string `json:"block"`
	Title     string `json:"title"`
	Target    string `json:"target"`
	URL       string `json:"url"`
	Thumb     string `json:"thumb"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
