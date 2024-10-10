package org

type ToOrganization struct {
	ID          string   `json:"id"`
	Brand       uint     `json:"brand,omitempty"`
	Name        string   `json:"name"`
	ValidStart  string   `json:"valid_start"`
	ValidEnd    string   `json:"valid_end"`
	User        string   `json:"user"`
	Telephone   string   `json:"telephone"`
	Province    uint     `json:"province,omitempty"`
	City        uint     `json:"city,omitempty"`
	Area        uint     `json:"area,omitempty"`
	Address     string   `json:"address,omitempty"`
	Longitude   float64  `json:"longitude"`
	Latitude    float64  `json:"latitude"`
	Description string   `json:"description"`
	Thumb       string   `json:"thumb"`
	Pictures    []string `json:"pictures"`
	CreatedAt   string   `json:"created_at"`
}
