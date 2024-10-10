package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableSysUser = "sys_user"

type SysUser struct {
	ID        string         `gorm:"column:id;primaryKey"`
	Username  *string        `gorm:"column:username"`
	Mobile    *string        `gorm:"column:mobile"`
	Email     *string        `gorm:"column:email"`
	Nickname  string         `gorm:"column:nickname"`
	Avatar    string         `gorm:"column:avatar"`
	Password  string         `gorm:"column:password"`
	IsEnable  uint8          `gorm:"column:is_enable"`
	CreatedAt carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	BindRoles []SysUserBindRole `gorm:"foreignKey:UserID;references:ID"`
}

func (SysUser) TableName() string {
	return TableSysUser
}
