package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpProductPicture = "shp_product_picture"

type ShpProductPicture struct {
	ID             uint            `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	ProductID      string          `gorm:"column:product_id"`
	URL            string          `gorm:"column:url"`
	Order          uint8           `gorm:"column:order"`
	IsDefault      uint8           `gorm:"column:is_default"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ShpProductPicture) TableName() string {
	return TableShpProductPicture
}
