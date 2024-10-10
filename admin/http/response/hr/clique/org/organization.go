package org

type ToOrganizationOfPaginate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Platform   uint16 `json:"platform"`
	Brand      string `json:"brand,omitempty"`
	ValidStart string `json:"valid_start"`
	ValidEnd   string `json:"valid_end"`
	IsEnable   uint8  `json:"is_enable"`
	CreatedAt  string `json:"created_at"`
}

type ToOrganizationOfInformation struct {
	ID          string `json:"id"`
	Platform    uint16 `json:"platform"`
	Brand       uint   `json:"brand,omitempty"`
	Name        string `json:"name"`
	ValidStart  string `json:"valid_start"`
	ValidEnd    string `json:"valid_end"`
	Description string `json:"description"`
	User        string `json:"user"`
	Telephone   string `json:"telephone"`
	IsEnable    uint8  `json:"is_enable"`
	CreatedAt   string `json:"created_at"`
}
