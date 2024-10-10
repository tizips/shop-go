package common

type ToAreaOfOpening struct {
	Parent uint `json:"parent" form:"parent" query:"parent" valid:"omitempty,gte=0" label:"父级"`
}
