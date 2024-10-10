package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysRole = "sys_role"

type SysRole struct {
	ID             uint           `gorm:"column:id;primaryKey"`
	Platform       uint16         `gorm:"column:platform"`
	OrganizationID *string        `gorm:"column:organization_id"`
	Name           string         `gorm:"column:name"`
	Summary        string         `gorm:"column:summary"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`

	BindPermissions []SysRoleBindPermission `gorm:"foreignKey:RoleID;references:ID"`
}

func (SysRole) TableName() string {
	return TableSysRole
}
