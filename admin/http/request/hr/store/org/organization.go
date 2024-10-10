package org

type DoOrganization struct {
	Name        string   `json:"name" form:"name" valid:"required,max=64" label:"名称"`
	ValidStart  string   `json:"valid_start" form:"valid_start" valid:"required,datetime=2006-01-02" label:"有效期：开始"`
	ValidEnd    string   `json:"valid_end" form:"valid_end" valid:"required,datetime=2006-01-02" label:"有效期：结束"`
	User        string   `json:"user" form:"user" valid:"required,max=32" label:"联系人"`
	Telephone   string   `json:"telephone" form:"telephone" valid:"required,max=32" label:"联系电话"`
	Province    uint     `json:"province" form:"province" valid:"required,gt=0" label:"省份"`
	City        uint     `json:"city" form:"city" valid:"required,gt=0" label:"城市"`
	Area        uint     `json:"area" form:"area" valid:"required,gt=0" label:"区县"`
	Address     string   `json:"address" form:"address" valid:"required,max=255" label:"详细地址"`
	Longitude   float64  `json:"longitude" form:"longitude" valid:"required,gte=-180,lte=180" label:"经度"`
	Latitude    float64  `json:"latitude" form:"latitude" valid:"required,gte=-90,lte=90" label:"纬度"`
	Description string   `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`
	Thumb       string   `json:"thumb" form:"thumb" valid:"required,max=255,http_url" label:"封面图"`
	Pictures    []string `json:"pictures" form:"pictures[]" valid:"required,min=1,max=8,unique,dive,max=255,http_url" label:"轮播图"`
}
