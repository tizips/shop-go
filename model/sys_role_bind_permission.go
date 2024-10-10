package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysRoleBindPermission = "sys_role_bind_permission"

type SysRoleBindPermission struct {
	ID             uint           `gorm:"column:id;primaryKey"`
	Platform       uint16         `gorm:"column:platform"`
	OrganizationID *string        `gorm:"column:organization_id"`
	Module         string         `gorm:"column:module"`
	RoleID         uint           `gorm:"column:role_id"`
	Permission     string         `gorm:"column:permission"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (SysRoleBindPermission) TableName() string {
	return TableSysRoleBindPermission
}
