package basic

type ToBannerOfPaginate struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Button      string `json:"button"`
	Picture     string `json:"picture"`
	Target      string `json:"target"`
	URL         string `json:"url"`
	Order       uint8  `json:"order"`
	IsEnable    uint8  `json:"is_enable"`
	CreatedAt   string `json:"created_at"`
}
