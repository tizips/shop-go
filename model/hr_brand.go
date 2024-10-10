package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableHRBrand = "hr_brand"

type HRBrand struct {
	ID             uint           `gorm:"column:id"`
	Platform       uint16         `gorm:"column:platform"`
	OrganizationID *string        `gorm:"column:organization_id"`
	Name           string         `gorm:"column:name"`
	Logo           string         `gorm:"column:logo"`
	Order          uint8          `gorm:"column:order"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (HRBrand) TableName() string {
	return TableHRBrand
}
