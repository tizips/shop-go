package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpPayment = "shp_payment"
)

type ShpPayment struct {
	ID             string         `gorm:"column:id;primaryKey"`             // ID
	Platform       uint16         `gorm:"column:platform"`                  // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`     // 组织ID
	UserID         string         `gorm:"column:user_id;index"`             // 用户id
	OrderID        string         `gorm:"column:order_id;index"`            // 订单ID
	No             *string        `gorm:"column:no;index"`                  // 第三方支付单号
	Channel        string         `gorm:"column:channel"`                   // 支付渠道：paypal=贝宝
	ChannelID      uint           `gorm:"column:channel_id"`                // 渠道ID
	Money          uint           `gorm:"column:money"`                     // 价格
	Currency       string         `gorm:"column:currency"`                  // 币种
	IsConfirmed    uint8          `gorm:"column:is_confirmed"`              // 是否确认：1=是；2=否
	Remark         string         `gorm:"column:remark"`                    // 备注
	Ext            string         `gorm:"column:ext"`                       // 扩展信息
	PaidAt         *carbon.Carbon `gorm:"column:paid_at"`                   // 支付时间
	ExpiredAt      carbon.Carbon  `gorm:"column:expired_at"`                // 过期时间
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                // 删除时间

	Order    *ShpOrder          `gorm:"foreignKey:ID;references:OrderID"`
	Channels *ShpPaymentChannel `gorm:"foreignKey:ID;references:ChannelID"`
}

func (ShpPayment) TableName() string {
	return TableShpPayment
}

const (
	ShpPaymentOfChannelPaypal = "paypal"
)
