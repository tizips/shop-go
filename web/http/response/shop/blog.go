package shop

type ToBlogOfPaginate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Thumb     string `json:"thumb"`
	Summary   string `json:"summary"`
	PostedAt  string `json:"posted_at"`
	CreatedAt string `json:"created_at"`
}

type ToBlogOfInformation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Thumb     string `json:"thumb"`
	Summary   string `json:"summary"`
	PostedAt  string `json:"posted_at"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
