package shop

type ToAppraisalOfPaginate struct {
	ID           uint                  `json:"id"`
	StarProduct  uint8                 `json:"star_product"`
	StarShipment uint8                 `json:"star_shipment"`
	Remark       string                `json:"remark"`
	Pictures     []string              `json:"pictures,omitempty"`
	Organization ToOrderOfOrganization `json:"organization"`
	Details      []ToOrderOfDetail     `json:"details"`
	CreatedAt    string                `json:"created_at"`
}
