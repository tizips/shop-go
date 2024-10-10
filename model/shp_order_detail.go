package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpOrderDetail = "shp_order_detail"
)

type ShpOrderDetail struct {
	ID             string         `gorm:"column:id;primaryKey"`                  // ID
	Platform       uint16         `gorm:"column:platform"`                       // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`                // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`          // 组织ID
	UserID         string         `gorm:"column:user_id;index"`                  // 用户id
	OrderID        string         `gorm:"column:order_id;index"`                 // 订单ID
	ProductID      string         `gorm:"column:product_id;index"`               // 产品ID
	SkuID          string         `gorm:"column:sku_id;index"`                   // SkuID
	AppraisalID    uint           `gorm:"column:appraisal_id;index"`             // 评价ID
	Name           string         `gorm:"column:name"`                           // 名称
	Specifications []string       `gorm:"column:specifications;serializer:json"` // 规格
	Picture        string         `gorm:"column:picture"`                        // 图片
	Price          uint           `gorm:"column:price"`                          // 价格
	CostPrice      uint           `gorm:"column:cost_price"`                     // 价格
	Quantity       uint           `gorm:"column:quantity"`                       // 数量
	TotalPrice     uint           `gorm:"column:total_price"`                    // 总价
	CouponPrice    uint           `gorm:"column:coupon_price"`                   // 优惠
	CostPrices     uint           `gorm:"column:cost_prices"`                    // 成本合计
	Prices         uint           `gorm:"column:prices"`                         // 合计
	Weight         uint           `gorm:"column:weight"`                         // 重量
	Refund         uint           `gorm:"column:refund"`                         // 退款
	Returned       uint           `gorm:"column:returned"`                       // 退货数量
	IsInvoiced     uint8          `gorm:"column:is_invoiced"`                    // 是否已开发票：1=是；2=否
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                     // 删除时间
}

func (s *ShpOrderDetail) TableName() string {
	return TableShpOrderDetail
}
