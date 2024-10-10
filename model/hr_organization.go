package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type HROrganization struct {
	ID          string         `gorm:"column:id"`
	Platform    uint16         `gorm:"column:platform"`
	CliqueID    *string        `gorm:"column:clique_id"`
	BrandID     uint           `gorm:"column:brand_id"`
	ParentID    *string        `gorm:"column:parent_id"`
	Name        string         `gorm:"column:name"`
	ValidStart  carbon.Carbon  `gorm:"column:valid_start" carbon:"type:date"`
	ValidEnd    carbon.Carbon  `gorm:"column:valid_end" carbon:"type:date"`
	User        string         `gorm:"column:user"`
	Telephone   string         `gorm:"column:telephone"`
	ProvinceId  uint           `gorm:"column:province_id"`
	CityID      uint           `gorm:"column:city_id"`
	AreaID      uint           `gorm:"column:area_id"`
	Address     string         `gorm:"column:address"`
	Longitude   float64        `gorm:"column:longitude"`
	Latitude    float64        `gorm:"column:latitude"`
	Description string         `gorm:"column:description"`
	IsEnable    uint8          `gorm:"column:is_enable"`
	CreatedAt   carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt   carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`

	Brand    *HRBrand                `gorm:"foreignKey:ID;references:BrandID"`
	Thumb    *HROrganizationThumb    `gorm:"foreignKey:OrganizationID;references:ID"`
	Pictures []HROrganizationPicture `gorm:"foreignKey:OrganizationID;references:ID"`
}

const TableHROrganization = "hr_organization"

func (HROrganization) TableName() string {
	return TableHROrganization
}
