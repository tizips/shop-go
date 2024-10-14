package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpRefund = "shp_refund"
)

type ShpRefund struct {
	ID             string         `gorm:"column:id;primaryKey"`             // ID
	Platform       uint16         `gorm:"column:platform"`                  // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`     // 组织ID
	UserID         string         `gorm:"column:user_id;index"`             // 用户ID
	OrderID        *string        `gorm:"column:order_id;index"`            // 订单ID
	DetailID       *string        `gorm:"column:detail_id;index"`           // 明细ID
	PaymentID      *string        `gorm:"column:payment_id;index"`          // 支付ID
	ServiceID      *string        `gorm:"column:service_id;index"`          // 售后ID
	No             string         `gorm:"column:no;index"`                  // 第三方支付单号
	Channel        string         `gorm:"column:channel"`                   // 支付渠道：paypal=贝宝
	Money          uint           `gorm:"column:money"`                     // 退款金额
	Currency       string         `gorm:"column:currency"`                  // 币种
	IsConfirmed    uint8          `gorm:"column:is_confirmed"`              // 是否确认：1=是；2=否
	Remark         string         `gorm:"column:remark"`                    // 备注
	Ext            string         `gorm:"column:ext"`                       // 扩展信息
	RefundedAt     *carbon.Carbon `gorm:"column:refunded_at"`               // 支付时间
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`          // 删除时间
}

func (ShpRefund) TableName() string {
	return TableShpRefund
}
