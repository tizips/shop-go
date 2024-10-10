package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableShpProductInformation = "shp_product_information"

type ShpProductInformation struct {
	ID             uint                               `gorm:"column:id;primaryKey"`
	Platform       uint16                             `gorm:"column:platform"`
	CliqueID       *string                            `gorm:"column:clique_id"`
	OrganizationID *string                            `gorm:"column:organization_id"`
	ProductID      string                             `gorm:"column:product_id"`
	Description    string                             `gorm:"column:description"`
	Attributes     []ShpProductInformationOfAttribute `gorm:"column:attributes;serializer:json"`
	CreatedAt      carbon.DateTime                    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime                    `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt                     `gorm:"column:deleted_at"`
}

func (ShpProductInformation) TableName() string {
	return TableShpProductInformation
}

type ShpProductInformationOfAttribute struct {
	Label string `json:"label"`
	Value string `json:"value,omitempty"`
}
