package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpServiceDetail = "shp_service_detail"
)

type ShpServiceDetail struct {
	ID             uint           `gorm:"column:id;primaryKey"`   // 主键ID
	Platform       uint16         `gorm:"column:platform"`        // 平台
	CliqueID       *string        `gorm:"column:clique_id"`       // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"` // 组织ID
	UserID         string         `gorm:"column:user_id"`         // 用户ID
	OrderID        string         `gorm:"column:order_id"`        // 订单ID
	ProductID      string         `gorm:"column:product_id"`      // 产品ID
	ServiceID      string         `gorm:"column:service_id"`      // 售后ID
	DetailID       string         `gorm:"column:detail_id"`       // 明细ID
	Quantity       uint           `gorm:"column:quantity"`        // 数量
	Refund         uint           `gorm:"column:refund"`          // 退款
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`      // 删除时间

	Detail *ShpOrderDetail `gorm:"foreignKey:ID;references:DetailID"`
}

func (ShpServiceDetail) TableName() string {
	return TableShpServiceDetail
}
