package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type ShpPage struct {
	ID             uint            `gorm:"column:id"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	Code           string          `gorm:"column:code"`
	Name           string          `gorm:"column:name"`
	IsSystem       uint8           `gorm:"column:is_system"`
	Content        string          `gorm:"column:content"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`

	SEO *ShpSEO `gorm:"foreignKey:ChannelID;references:ID"`
}

const TableShpPage = "shp_page"

func (ShpPage) TableName() string {
	return TableShpPage
}
