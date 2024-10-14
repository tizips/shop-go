package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpOrderAddress = "shp_order_address"
)

type ShpOrderAddress struct {
	ID             uint           `gorm:"column:id;primaryKey"`             // ID
	Platform       uint16         `gorm:"column:platform"`                  // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`     // 组织ID
	OrderID        string         `gorm:"column:order_id;index"`            // 订单ID
	UserID         string         `gorm:"column:user_id;index"`             // 用户id
	FirstName      string         `gorm:"column:first_name"`                // first name
	LastName       string         `gorm:"column:last_name"`                 // last name
	Company        string         `gorm:"column:company"`                   // 公司
	Country        string         `gorm:"column:country"`                   // 国家
	Prefecture     string         `gorm:"column:prefecture"`                // 州府
	City           string         `gorm:"column:city"`                      // 城市
	Street         string         `gorm:"column:street"`                    // 街道
	Detail         string         `gorm:"column:detail"`                    // 详细地址
	Postcode       string         `gorm:"column:postcode"`                  // 邮编
	Phone          string         `gorm:"column:phone"`                     // 电话
	Email          string         `gorm:"column:email"`                     // 邮箱
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                // 删除时间

	Shipment *ShpShipment `gorm:"foreignKey:OrderID;references:OrderID"`
}

func (ShpOrderAddress) TableName() string {
	return TableShpOrderAddress
}
