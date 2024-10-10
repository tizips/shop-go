package model

import (
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/cache"
	"gorm.io/gorm"
)

type ShpUser struct {
	ID             string          `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	Email          string          `gorm:"column:email"`
	FirstName      string          `gorm:"column:first_name"`
	LastName       string          `gorm:"column:last_name"`
	Password       string          `gorm:"column:password"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
	//
	//Lv *MemLevel `gorm:"foreignKey:Code;references:Level"`

	cache.Model
}

const TableShpUser = "shp_user"

func (ShpUser) TableName() string {
	return TableShpUser
}
