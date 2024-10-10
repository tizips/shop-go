package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableSysQueue = "sys_queue"
)

type SysQueue struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"` // ID
	Queue     string         `gorm:"column:queue"`                       // 队列
	Message   string         `gorm:"column:message"`                     // 消息
	Error     string         `gorm:"column:error"`                       // 错误
	IsOmitted uint8          `gorm:"column:is_omitted"`                  // 是否忽略：1=是；2=否
	IsTried   uint8          `gorm:"column:is_tried"`                    // 是否重试：1=是；2=否
	IsHandled uint8          `gorm:"column:is_handled"`                  // 是否处理：1=是；2=否
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime"`   // 创建时间
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"`   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`                  // 删除时间
}

func (SysQueue) TableName() string {
	return TableSysQueue
}
