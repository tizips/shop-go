package basic

type ToBlogOfPaginate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Thumb     string `json:"thumb"`
	PostedAt  string `json:"posted_at"`
	IsTop     uint8  `json:"is_top"`
	CreatedAt string `json:"created_at"`
}

type ToBlogOfInformation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Thumb       string `json:"thumb"`
	IsTop       uint8  `json:"is_top"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	PostedAt    string `json:"posted_at"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
}
