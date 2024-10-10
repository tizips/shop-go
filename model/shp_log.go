package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpLog = "shp_log"
)

type ShpLog struct {
	ID             uint           `gorm:"column:id;primaryKey;autoIncrement"` // ID
	Platform       uint16         `gorm:"column:platform"`                    // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`             // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`       // 组织ID
	UserID         string         `gorm:"column:user_id;index"`               // 用户ID
	OrderID        string         `gorm:"column:order_id;index"`              // 订单ID
	DetailID       *string        `gorm:"column:detail_id;index"`             // 明细ID
	ServiceID      *string        `gorm:"column:service_id;index"`            // 售后ID
	Action         string         `gorm:"column:action"`                      // 操作
	Content        string         `gorm:"column:content"`                     // 内容
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"`   // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"`   // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                  // 删除时间
}

func (ShpLog) TableName() string {
	return TableShpLog
}
