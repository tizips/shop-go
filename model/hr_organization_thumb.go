package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type HROrganizationThumb struct {
	ID             uint           `gorm:"column:id"`
	OrganizationID string         `gorm:"column:organization_id"`
	CliqueID       *string        `gorm:"column:clique_id"`
	URL            string         `gorm:"column:url"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

const TableHROrganizationThumb = "hr_organization_thumb"

func (HROrganizationThumb) TableName() string {
	return TableHROrganizationThumb
}
