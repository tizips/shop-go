package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpBlog = "shp_blog"
)

type ShpBlog struct {
	ID             string         `gorm:"column:id;primaryKey"`   // 主键ID
	Platform       uint16         `gorm:"column:platform"`        // 平台
	CliqueID       *string        `gorm:"column:clique_id"`       // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"` // 组织ID
	Name           string         `gorm:"column:name"`            // 名称
	Thumb          string         `gorm:"column:thumb"`           // 图片
	Summary        string         `gorm:"column:summary"`         // 简介
	PostedAt       carbon.Carbon  `gorm:"column:posted_at"`       // 发布日期
	IsTop          uint8          `gorm:"column:is_top"`          // 是否置顶：1=是；2=否
	Content        string         `gorm:"column:content"`         // 内容
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`      // 删除时间

	SEO *ShpSEO `gorm:"foreignKey:ChannelID;references:ID"`
}

func (ShpBlog) TableName() string {
	return TableShpBlog
}
