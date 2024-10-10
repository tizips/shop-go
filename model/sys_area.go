package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type SysArea struct {
	ID        uint           `gorm:"column:id"`
	Level     string         `gorm:"column:level"`
	ParentID  uint           `gorm:"column:parent_id"`
	Name      string         `gorm:"column:name"`
	Code      string         `gorm:"column:code"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

const TableSysArea = "sys_area"

func (SysArea) TableName() string {
	return TableSysArea
}

const (
	SysAreaOfLevelLv1 = "lv_1"
	SysAreaOfLevelLv2 = "lv_2"
	SysAreaOfLevelLv3 = "lv_3"
)
