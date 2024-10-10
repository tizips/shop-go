package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpAdvertise = "shp_advertise"
)

type ShpAdvertise struct {
	ID             uint           `gorm:"column:id;primaryKey"`   // 主键ID
	Platform       uint16         `gorm:"column:platform"`        // 平台
	CliqueID       *string        `gorm:"column:clique_id"`       // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"` // 组织ID
	Page           string         `gorm:"column:page"`            // 页面：home=首页
	Block          string         `gorm:"column:block"`           // 广告位：new_product=新品
	Title          string         `gorm:"column:title"`           // 标题
	Target         string         `gorm:"column:target"`          // 打开方式：blank=新窗口；self=该窗口
	URL            string         `gorm:"column:url"`             // 链接
	Thumb          string         `gorm:"column:thumb"`           // 图片
	Order          uint8          `gorm:"column:order"`           // 排序（正序）
	IsEnable       uint8          `gorm:"column:is_enable"`       // 是否启用；1=是；2=否
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`      // 删除时间
}

func (ShpAdvertise) TableName() string {
	return TableShpAdvertise
}

const (
	ShpAdvertiseOfPageHome = "home"

	ShpAdvertiseOfBlockNewProduct = "new_product"

	ShpAdvertiseOfTargetBlank = "_blank"
	ShpAdvertiseOfTargetSelf  = "_self"
)
