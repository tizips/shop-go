package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpSku = "shp_sku"

type ShpSku struct {
	ID             string          `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	ProductID      string          `gorm:"column:product_id"`
	Code           string          `gorm:"column:code"`
	Price          uint            `gorm:"column:price"`
	OriginPrice    uint            `gorm:"column:origin_price"`
	CostPrice      uint            `gorm:"column:cost_price"`
	Stock          uint            `gorm:"column:stock"`
	Warn           uint            `gorm:"column:warn"`
	Picture        string          `gorm:"column:picture"`
	IsDefault      uint8           `gorm:"column:is_default"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`

	Product *ShpProduct `gorm:"foreignKey:ID;references:ProductID"`
}

func (ShpSku) TableName() string {
	return TableShpSku
}
