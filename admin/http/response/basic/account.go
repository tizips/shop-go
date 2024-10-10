package basic

type ToAccountOfInformation struct {
	Nickname string `json:"nickname"`
	Username string `json:"username,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

type ToAccountOfPlatform struct {
	Platform     uint16 `json:"platform"`
	Organization string `json:"organization"`
	Org          string `json:"org,omitempty"`
	Back         string `json:"back,omitempty"`
}

type ToAccountOfModules struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
