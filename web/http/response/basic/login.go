package basic

type DoLogin struct {
	Token    string `json:"token"`
	Lifetime int    `json:"lifetime"`
}
