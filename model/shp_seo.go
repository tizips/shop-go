package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpSEO = "shp_seo"

type ShpSEO struct {
	ID             uint            `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	Channel        string          `gorm:"column:channel"`
	ChannelID      string          `gorm:"column:channel_id"`
	Title          string          `gorm:"column:title"`
	Keyword        string          `gorm:"column:keyword"`
	Description    string          `gorm:"column:description"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ShpSEO) TableName() string {
	return TableShpSEO
}

const (
	ShpSEOForChannelOfProduct  = "product"
	ShpSEOForChannelOfPage     = "page"
	ShpSEOForChannelOfCategory = "category"
	ShpSEOForChannelOfBlog     = "blog"
)
