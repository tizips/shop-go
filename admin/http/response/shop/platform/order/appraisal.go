package order

type ToAppraisalOfPaginate struct {
	ID           uint                       `json:"id"`
	Order        string                     `json:"order"`
	Organization *ToAppraisalOfOrganization `json:"organization,omitempty"`
	FirstName    string                     `json:"first_name"`
	LastName     string                     `json:"last_name"`
	Email        string                     `json:"email"`
	StarProduct  uint8                      `json:"star_product"`
	StarShipment uint8                      `json:"star_shipment"`
	Remark       string                     `json:"remark"`
	Pictures     []string                   `json:"pictures,omitempty"`
	CreatedAt    string                     `json:"created_at"`
}

type ToAppraisalOfOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
