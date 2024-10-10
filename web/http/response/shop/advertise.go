package shop

type ToAdvertises struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Thumb     string `json:"thumb"`
	Target    string `json:"target"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}
