package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

// ShpCategory 采购-栏目表
type ShpCategory struct {
	ID             uint           `gorm:"column:id;primaryKey"`
	Platform       uint16         `gorm:"column:platform"`
	CliqueID       *string        `gorm:"column:clique_id"`
	OrganizationID *string        `gorm:"column:organization_id"`
	Level          string         `gorm:"column:level;comment:分类层级：lv_1=一级分类；lv_2=二级分类；lv_3=三级分类"`
	ParentID       uint           `gorm:"column:parent_id;comment:父级ID"`
	Name           string         `gorm:"column:name;comment:名称"`
	Order          uint8          `gorm:"column:order;comment:序号：正序"`
	IsEnable       uint8          `gorm:"column:is_enable"`
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime" carbon:"type:dateTime"`
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime" carbon:"type:dateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;default:null"`

	Parent *ShpCategory `gorm:"foreignKey:ID;references:ParentID"`
}

const TableNameShpCategory = "shp_category"

func (ShpCategory) TableName() string {
	return TableNameShpCategory
}

const (
	ShpCategoryOfLevel1 = "lv_1"
	ShpCategoryOfLevel2 = "lv_2"
	ShpCategoryOfLevel3 = "lv_3"
)
