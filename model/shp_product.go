package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpProduct = "shp_product"

type ShpProduct struct {
	ID             string          `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	I1CategoryID   uint            `gorm:"column:i1_category_id"`
	I2CategoryID   uint            `gorm:"column:i2_category_id"`
	I3CategoryID   uint            `gorm:"column:i3_category_id"`
	Name           string          `gorm:"column:name"`
	Summary        string          `gorm:"column:summary"`
	IsHot          uint8           `gorm:"column:is_hot"`
	IsRecommend    uint8           `gorm:"column:is_recommend"`
	IsMultiple     uint8           `gorm:"column:is_multiple"`
	IsFreeShipping uint8           `gorm:"column:is_free_shipping"`
	IsFreeze       uint8           `gorm:"column:is_freeze"`
	IsEnable       uint8           `gorm:"column:is_enable"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`

	SEO            ShpSEO                `gorm:"foreignKey:ChannelID;references:ID"`
	Picture        *ShpProductPicture    `gorm:"foreignKey:ProductID;references:ID"`
	Pictures       []ShpProductPicture   `gorm:"foreignKey:ProductID;references:ID"`
	Information    ShpProductInformation `gorm:"foreignKey:ProductID;references:ID"`
	Price          *ShpSku               `gorm:"foreignKey:ProductID;references:ID"`
	SKUS           []ShpSku              `gorm:"foreignKey:ProductID;references:ID"`
	Specifications []ShpSpecification    `gorm:"foreignKey:ProductID;references:ID"`
}

func (ShpProduct) TableName() string {
	return TableShpProduct
}
