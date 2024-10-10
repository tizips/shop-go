package model

import (
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/cache"
	"gorm.io/gorm"
)

const (
	TableSysSecret = "sys_secret"
)

type SysSecret struct {
	ID             string         `gorm:"column:id;primaryKey"`             // id
	Platform       uint16         `gorm:"column:platform"`                  // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`     // 组织ID
	Name           string         `gorm:"column:name"`                      // 名称
	Secret         string         `gorm:"column:secret"`                    // Secret
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`          // 删除时间

	cache.Model
}

func (s *SysSecret) TableName() string {
	return TableSysSecret
}
