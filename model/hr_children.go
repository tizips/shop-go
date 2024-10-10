package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableHRChildren = "hr_children"

type HRChildren struct {
	ID             uint           `gorm:"column:id"`
	OrganizationID string         `gorm:"column:organization_id"`
	ChildID        string         `gorm:"column:child_id"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (HRChildren) TableName() string {
	return TableHRChildren
}
