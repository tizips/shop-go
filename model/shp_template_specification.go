package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableNameShpTemplateSpecification = "shp_template_specification"

type ShpTemplateSpecification struct {
	ID             uint            `gorm:"column:id;primaryKey"`
	Platform       uint16          `gorm:"column:platform"`
	CliqueID       *string         `gorm:"column:clique_id"`
	OrganizationID *string         `gorm:"column:organization_id"`
	Name           string          `gorm:"column:name"`
	Label          string          `gorm:"column:label"`
	Options        []string        `gorm:"column:options;serializer:json"`
	IsEnable       uint8           `gorm:"column:is_enable"`
	CreatedAt      carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (ShpTemplateSpecification) TableName() string {
	return TableNameShpTemplateSpecification
}
