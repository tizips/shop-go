package common

type ToSetting struct {
	ID         uint   `json:"id"`
	Type       string `json:"type"`
	Label      string `json:"label"`
	Key        string `json:"key"`
	Val        string `json:"val"`
	IsRequired uint8  `json:"is_required"`
	CreatedAt  string `json:"created_at"`
}
