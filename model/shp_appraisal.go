package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpAppraisal = "shp_appraisal"
)

type ShpAppraisal struct {
	ID             uint           `gorm:"column:id;primaryKey"`            // 主键ID
	Platform       uint16         `gorm:"column:platform"`                 // 平台
	CliqueID       *string        `gorm:"column:clique_id"`                // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"`          // 组织ID
	UserID         string         `gorm:"column:user_id"`                  // 用户ID
	OrderID        string         `gorm:"column:order_id"`                 // 订单ID
	StarProduct    uint8          `gorm:"column:star_product"`             // 商品评分：1-5
	StarShipment   uint8          `gorm:"column:star_shipment"`            // 物流评分：1-5
	Remark         string         `gorm:"column:remark"`                   // 备注
	Pictures       []string       `gorm:"column:pictures;serializer:json"` // 证据图
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`               // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`               // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`               // 删除时间

	User         *ShpUser         `gorm:"foreignKey:ID;references:UserID"`
	Organization *HROrganization  `gorm:"foreignKey:ID;references:OrganizationID"`
	Order        *ShpOrder        `gorm:"foreignKey:ID;references:OrderID"`
	Details      []ShpOrderDetail `gorm:"foreignKey:OrderID;references:OrderID"`
}

func (ShpAppraisal) TableName() string {
	return TableShpAppraisal
}
