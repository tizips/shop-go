package basic

type ToCategoryOfChildren struct {
	Parent uint  `query:"parent" valid:"omitempty,gte=0" label:"父级"`
	All    uint8 `query:"all" valid:"omitempty,eq=1" label:"全部"`
}
