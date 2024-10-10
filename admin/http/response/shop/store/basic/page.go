package basic

type ToPageOfPaginate struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	IsSystem  uint8  `json:"is_system"`
	CreatedAt string `json:"created_at"`
}

type ToPageOfInformation struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Content     string `json:"content"`
	IsSystem    uint8  `json:"is_system,omitempty"`
	CreatedAt   string `json:"created_at"`
}
