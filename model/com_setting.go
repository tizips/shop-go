package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableComSetting = "com_setting"

type ComSetting struct {
	ID             uint           `gorm:"column:id;primaryKey"`
	Platform       uint16         `gorm:"column:platform"`
	OrganizationID *string        `gorm:"column:organization_id"`
	Module         string         `gorm:"column:module"`
	Type           string         `gorm:"column:type"`
	Label          string         `gorm:"column:label"`
	Key            string         `gorm:"column:key"`
	Val            string         `gorm:"column:val"`
	IsRequired     uint8          `gorm:"column:is_required"`
	Order          uint8          `gorm:"column:order"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ComSetting) TableName() string {
	return TableComSetting
}

const (
	ComSettingForTypeOfInput    = "input"
	ComSettingForTypeOfEnable   = "enable"
	ComSettingForTypeOfURL      = "url"
	ComSettingForTypeOfEmail    = "email"
	ComSettingForTypeOfPicture  = "picture"
	ComSettingForTypeOfTextarea = "textarea"

	ComSettingForIsRequiredOfYes = 1
	ComSettingForIsRequiredOfNo  = 2
)
