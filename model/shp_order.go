package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpOrder = "shp_order"
)

type ShpOrder struct {
	ID             string         `gorm:"column:id;primaryKey"`             // ID
	Platform       uint16         `gorm:"column:platform"`                  // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`     // 组织ID
	UserID         string         `gorm:"column:user_id;index"`             // 用户ID
	CostShipping   uint           `gorm:"column:cost_shipping"`             // 运费
	TotalPrice     uint           `gorm:"column:total_price"`               // 总价
	CouponPrice    uint           `gorm:"column:coupon_price"`              // 优惠
	Prices         uint           `gorm:"column:prices"`                    // 合计
	CostPrices     uint           `gorm:"column:cost_prices"`               // 成本合计
	Refund         uint           `gorm:"column:refund"`                    // 退款
	Status         string         `gorm:"column:status"`                    // 订单状态：pay=待支付；shipment=待发货；receipt=待收货；received=已收货；completed=已完成；cancel=已取消；closed=已关闭
	Remark         string         `gorm:"column:remark"`                    // 备注
	PaymentID      *string        `json:"payment_id"`                       // 支付ID
	IsPaid         uint8          `gorm:"column:is_paid"`                   // 是否支付：1=是；2=否
	IsInvoice      uint8          `gorm:"column:is_invoice"`                // 是否开发票：1=是；2=否
	IsAppraisal    uint8          `gorm:"column:is_appraisal"`              // 是否评价：1=是；2=否
	CompletedAt    *carbon.Carbon `gorm:"column:completed_at"`              // 完成时间
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"` // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                // 删除时间

	Organization *HROrganization  `gorm:"foreignKey:ID;references:OrganizationID"`
	Details      []ShpOrderDetail `gorm:"foreignKey:OrderID;references:ID"`
	Payment      *ShpPayment      `gorm:"foreignKey:OrderID;references:ID"`
	Shipment     *ShpShipment     `gorm:"foreignKey:OrderID;references:ID"`
	Address      *ShpOrderAddress `gorm:"foreignKey:OrderID;references:ID"`
	//Invoice      *ShpOrderInvoice `gorm:"foreignKey:OrderID;references:ID"`
	Logs     []ShpLog     `gorm:"foreignKey:OrderID;references:ID"`
	Service  *ShpService  `gorm:"foreignKey:OrderID;references:ID"`
	Services []ShpService `gorm:"foreignKey:OrderID;references:ID"`
}

func (s *ShpOrder) TableName() string {
	return TableShpOrder
}

const (
	ShpOrderOfStatusPay       = "pay"
	ShpOrderOfStatusShipment  = "shipment"
	ShpOrderOfStatusReceipt   = "receipt"
	ShpOrderOfStatusReceived  = "received"
	ShpOrderOfStatusCompleted = "completed"
	ShpOrderOfStatusClosed    = "closed"
)
