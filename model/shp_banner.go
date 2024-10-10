package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpBanner = "shp_banner"

type ShpBanner struct {
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	ID             uint            `gorm:"column:id;primaryKey"`
	Name           string          `gorm:"column:name"`
	Description    string          `gorm:"column:description"`
	Button         string          `gorm:"column:button"`
	Picture        string          `gorm:"column:picture"`
	Target         string          `gorm:"column:target"`
	URL            string          `gorm:"column:url"`
	Order          uint8           `gorm:"column:order"`
	IsEnable       uint8           `gorm:"column:is_enable"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ShpBanner) TableName() string {
	return TableShpBanner
}
