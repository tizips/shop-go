package shop

type ToBanners struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Button      string `json:"button,omitempty"`
	Picture     string `json:"picture"`
	Target      string `json:"target"`
	URL         string `json:"url,omitempty"`
	CreatedAt   string `json:"created_at"`
}
