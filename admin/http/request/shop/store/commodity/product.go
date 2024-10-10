package commodity

import "github.com/herhe-com/framework/contracts/http/request"

type DoProductOfCreate struct {
	Category       uint                   `json:"category" form:"category" valid:"required,gt=0" label:"栏目"`
	Name           string                 `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Summary        string                 `json:"summary" form:"summary" valid:"omitempty" label:"简介"`
	IsHot          uint8                  `json:"is_hot" form:"is_hot" valid:"required,oneof=1 2" label:"是否热销"`
	IsRecommend    uint8                  `json:"is_recommend" form:"is_recommend" valid:"required,oneof=1 2" label:"是否推荐"`
	IsMultiple     uint8                  `json:"is_multiple" form:"is_multiple" valid:"required,oneof=1 2" label:"是否多规格"`
	IsFreeShipping uint8                  `json:"is_free_shipping" form:"is_free_shipping" valid:"required,oneof=1 2" label:"是否包邮"`
	IsFreeze       uint8                  `json:"is_freeze" form:"is_freeze" valid:"required,oneof=1 2" label:"是否冻结"`
	Pictures       []string               `json:"pictures" form:"pictures[]" valid:"required,min=1,max=12,unique,dive,required,max=255,http_url" label:"轮播大图"`
	Title          string                 `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword        string                 `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description    string                 `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Price          uint                   `json:"price" form:"price" valid:"required_if=IsMultiple 2,omitempty,gt=0,max=999999999" label:"价格"`
	OriginPrice    uint                   `json:"origin_price" form:"origin_price" valid:"omitempty,gte=0,max=999999999" label:"原价"`
	CostPrice      uint                   `json:"cost_price" form:"cost_price" valid:"omitempty,gte=0,max=999999999" label:"成本"`
	Stock          uint                   `json:"stock" form:"stock" valid:"omitempty,gt=0,max=999999999" label:"库存"`
	Warn           uint                   `json:"warn" form:"warn" valid:"omitempty,gte=0,max=999999999" label:"库存警告"`
	Attributes     []DoProductOfAttribute `json:"attributes" form:"attributes[]" valid:"required,min=1,unique=Label,dive" label:"属性"`
	Information    string                 `json:"information" form:"information" valid:"required" label:"详情"`

	request.Enable
}

type DoProductOfUpdate struct {
	Category       uint                   `json:"category" form:"category" valid:"required,gt=0" label:"栏目"`
	Name           string                 `json:"name" form:"name" valid:"required,max=255" label:"标题"`
	Summary        string                 `json:"summary" form:"summary" valid:"omitempty" label:"简介"`
	IsHot          uint8                  `json:"is_hot" form:"is_hot" valid:"required,oneof=1 2" label:"是否热销"`
	IsRecommend    uint8                  `json:"is_recommend" form:"is_recommend" valid:"required,oneof=1 2" label:"是否推荐"`
	IsMultiple     uint8                  `json:"is_multiple" form:"is_multiple" valid:"required,oneof=1 2" label:"是否多规格"`
	IsFreeShipping uint8                  `json:"is_free_shipping" form:"is_free_shipping" valid:"required,oneof=1 2" label:"是否包邮"`
	IsFreeze       uint8                  `json:"is_freeze" form:"is_freeze" valid:"required,oneof=1 2" label:"是否冻结"`
	Pictures       []string               `json:"pictures" form:"pictures[]" valid:"required,min=1,unique,dive,required,max=255,http_url" label:"轮播大图"`
	Title          string                 `json:"title" form:"title" valid:"omitempty,max=255" label:"SEO 标题"`
	Keyword        string                 `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"SEO 关键词"`
	Description    string                 `json:"description" form:"description" valid:"omitempty,max=255" label:"SEO 描述"`
	Price          uint                   `json:"price" form:"price" valid:"required_if=IsMultiple 2,omitempty,gt=0,max=999999999" label:"价格"`
	OriginPrice    uint                   `json:"origin_price" form:"origin_price" valid:"omitempty,gte=0,max=999999999" label:"原价"`
	CostPrice      uint                   `json:"cost_price" form:"cost_price" valid:"omitempty,gte=0,max=999999999" label:"成本"`
	Stock          uint                   `json:"stock" form:"stock" valid:"omitempty,gt=0,max=999999999" label:"库存"`
	Warn           uint                   `json:"warn" form:"warn" valid:"omitempty,gte=0,max=999999999" label:"库存警告"`
	Attributes     []DoProductOfAttribute `json:"attributes" form:"attributes[]" valid:"required,min=1,unique=Label,dive" label:"属性"`
	Information    string                 `json:"information" form:"information" valid:"required" label:"详情"`

	request.Enable
	request.IDOfSnowflake
}

type DoProductOfAttribute struct {
	Label string `json:"label" form:"label" valid:"required,max=120" label:"属性名"`
	Value string `json:"value" form:"value" valid:"omitempty,max=255" label:"属性值"`
}

type DoProductOfDelete struct {
	request.IDOfSnowflake
}

type DoProductOfEnable struct {
	request.IDOfSnowflake
	request.Enable
}

type ToProductOfPaginate struct {
	Keyword        string `query:"keyword" valid:"omitempty,max=255" label:"关键词"`
	IsHot          uint8  `query:"is_hot" valid:"omitempty,oneof=1 2" label:"是否热销"`
	IsRecommend    uint8  `query:"is_recommend" valid:"omitempty,oneof=1 2" label:"是否推荐"`
	IsMultiple     uint8  `query:"is_multiple" valid:"omitempty,oneof=1 2" label:"是否多规格"`
	IsFreeShipping uint8  `query:"is_free_shipping" valid:"omitempty,oneof=1 2" label:"是否包邮"`
	IsFreeze       uint8  `query:"is_freeze" valid:"omitempty,oneof=1 2" label:"是否冻结"`
	IsEnable       uint8  `query:"is_enable" valid:"omitempty,oneof=1 2" label:"启用"`

	request.Paginate
}

type ToProductOfInformation struct {
	request.IDOfSnowflake
}

type ToProductOfSpecification struct {
	request.IDOfSnowflake
}

type DoProductOfSpecification struct {
	request.IDOfSnowflake
	Specification []DoProductOfSpecificationWithGroup `json:"specification" form:"specification[]" valid:"required,min=1,max=3,unique=ID,unique=Name,dive" label:"属性"`
	SKUS          []DoProductWithSKU                  `json:"skus" form:"skus[]" valid:"required,min=1,max=27,unique=Key,dive" label:"SKU"`
}

type DoProductOfSpecificationWithGroup struct {
	ID       uint                                   `json:"id" form:"id" valid:"omitempty,gt=0" label:"ID"`
	Name     string                                 `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Virtual  bool                                   `json:"virtual" form:"virtual" valid:"omitempty,boolean" label:"是否虚拟 ID"`
	Children []DoProductOfSpecificationWithChildren `json:"children" form:"children[]" valid:"required,min=1,max=9,unique=ID,unique=Name,dive" label:"子属性"`
}

type DoProductOfSpecificationWithChildren struct {
	ID      uint   `json:"id" form:"id" valid:"omitempty,gt=0" label:"ID"`
	Name    string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Virtual bool   `json:"virtual" form:"virtual" valid:"omitempty,boolean" label:"是否虚拟 ID"`
}

type DoProductWithSKU struct {
	ID          string `json:"id" form:"id" valid:"omitempty,snowflake" label:"ID"`
	Key         string `json:"key" form:"key" valid:"required" label:"Key"`
	Picture     string `json:"picture" form:"picture" valid:"omitempty,max=255,http_url" label:"图片"`
	Price       uint   `json:"price" form:"price" valid:"required_if=IsMultiple 2,required,gt=0,max=999999999" label:"价格"`
	OriginPrice uint   `json:"origin_price" form:"origin_price" valid:"required_if=IsMultiple 2,omitempty,gte=0,max=999999999" label:"原价"`
	CostPrice   uint   `json:"cost_price" form:"cost_price" valid:"required_if=IsMultiple 2,omitempty,gte=0,max=999999999" label:"成本"`
	Stock       uint   `json:"stock" form:"stock" valid:"required_if=IsMultiple 2,omitempty,gt=0,max=999999999" label:"库存"`
	Warn        uint   `json:"warn" form:"warn" valid:"required_if=IsMultiple 2,omitempty,gte=0,max=999999999" label:"库存警告"`
	IsDefault   uint8  `json:"is_default" form:"is_default" valid:"required,oneof=1 2" label:"是否默认"`
}
