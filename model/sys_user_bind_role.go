package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysUserBindRole = "sys_user_bind_role"

type SysUserBindRole struct {
	ID             uint           `gorm:"column:id;primaryKey"`
	Platform       uint16         `gorm:"column:platform"`
	OrganizationID *string        `gorm:"column:organization_id"`
	UserID         string         `gorm:"column:user_id"`
	RoleID         uint           `gorm:"column:role_id"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`

	Role *SysRole `gorm:"foreignKey:RoleID;references:ID"`
}

func (SysUserBindRole) TableName() string {
	return TableSysUserBindRole
}
