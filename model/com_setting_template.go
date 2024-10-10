package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableComSettingTemplate = "com_setting_template"

type ComSettingTemplate struct {
	ID         uint           `gorm:"column:id;primaryKey"`
	Module     string         `gorm:"column:module"`
	Type       string         `gorm:"column:type"`
	Label      string         `gorm:"column:label"`
	Key        string         `gorm:"column:key"`
	IsRequired uint8          `gorm:"column:is_required"`
	Order      uint8          `gorm:"column:order"`
	CreatedAt  carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt  carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ComSettingTemplate) TableName() string {
	return TableComSettingTemplate
}
