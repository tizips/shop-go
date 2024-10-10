package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type ShpShipping struct {
	ID             uint            `gorm:"column:id"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	Name           string          `gorm:"column:name"`
	Money          uint            `gorm:"column:money"`
	Query          string          `gorm:"column:query"`
	Order          uint8           `gorm:"column:order"`
	IsEnable       uint8           `gorm:"column:is_enable"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableShpShipping = "shp_shipping"

func (ShpShipping) TableName() string {
	return TableShpShipping
}
