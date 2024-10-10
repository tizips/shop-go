package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableHROrganizationPicture = "hr_organization_picture"

type HROrganizationPicture struct {
	ID             uint           `gorm:"column:id"`
	OrganizationID string         `gorm:"column:organization_id"`
	CliqueID       *string        `gorm:"column:clique_id"`
	URL            string         `gorm:"column:url"`
	Order          uint8          `gorm:"column:order"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (HROrganizationPicture) TableName() string {
	return TableHROrganizationPicture
}
