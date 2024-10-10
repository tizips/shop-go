package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpShipment = "shp_shipment"
)

type ShpShipment struct {
	ID             uint           `gorm:"column:id;primaryKey;autoIncrement"` // ID
	Platform       uint16         `gorm:"column:platform"`                    // 平台
	CliqueID       *string        `gorm:"column:clique_id"`                   // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"`             // 组织ID
	OrderID        string         `gorm:"column:order_id"`                    // 订单ID
	UserID         string         `gorm:"column:user_id"`                     // 用户ID
	ShippingID     uint           `gorm:"column:shipping_id"`                 // 快递ID
	Money          uint           `gorm:"column:money"`                       // 费用
	Company        string         `gorm:"column:company"`                     // 快递公司
	No             string         `gorm:"column:no"`                          // 快递单号
	Remark         string         `gorm:"column:remark"`                      // 备注
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"`   // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"`   // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                  // 删除时间
}

func (ShpShipment) TableName() string {
	return TableShpShipment
}
