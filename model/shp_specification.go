package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpSpecification = "shp_specification"

type ShpSpecification struct {
	ID             uint            `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	ProductID      string          `gorm:"column:product_id"`
	ParentID       uint            `gorm:"column:parent_id"`
	Name           string          `gorm:"column:name"`
	Order          uint8           `gorm:"column:order"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`

	Specifications []ShpSpecification `gorm:"foreignKey:ParentID;references:ID"`
}

func (ShpSpecification) TableName() string {
	return TableShpSpecification
}
