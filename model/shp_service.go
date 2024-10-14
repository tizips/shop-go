package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpService = "shp_service"
)

type ShpService struct {
	ID                   string                `gorm:"column:id;primaryKey"`                         // 主键ID
	Platform             uint16                `gorm:"column:platform"`                              // 平台
	CliqueID             *string               `gorm:"column:clique_id"`                             // 集团ID
	OrganizationID       *string               `gorm:"column:organization_id"`                       // 组织ID
	UserID               string                `gorm:"column:user_id"`                               // 用户ID
	OrderID              string                `gorm:"column:order_id"`                              // 订单ID
	DetailID             *string               `gorm:"column:detail_id"`                             // 明细ID
	Type                 string                `gorm:"column:type"`                                  // 类型：un_receipt=未收到货；refund=退货退款；exchange=换货；express=物流
	Status               string                `gorm:"column:status"`                                // 状态：pending=待处理；user=等待用户发货；org=等待商家发货；confirm_user=等待用户确认；confirm_org=等待商家确认；finish=完成；closed=已关闭
	Result               string                `gorm:"column:result"`                                // 结果：agree=同意；refuse=拒绝
	Reason               string                `gorm:"column:reason"`                                // 原因
	Pictures             []string              `gorm:"column:pictures;serializer:json"`              // 证据图
	Subtotal             uint                  `gorm:"column:subtotal"`                              // 商品退款
	Shipping             uint                  `gorm:"column:shipping"`                              // 运费退款
	Refunds              uint                  `gorm:"column:refund"`                                // 退款金额
	ShipmentUser         *ShpServiceOfShipment `gorm:"column:shipment_user;serializer:json"`         // 用户发货信息
	ShipmentOrganization *ShpServiceOfShipment `gorm:"column:shipment_organization;serializer:json"` // 商家发货信息
	CreatedAt            carbon.Carbon         `gorm:"column:created_at"`                            // 创建时间
	UpdatedAt            carbon.Carbon         `gorm:"column:updated_at"`                            // 更新时间
	DeletedAt            gorm.DeletedAt        `gorm:"column:deleted_at"`                            // 删除时间

	Organization *HROrganization    `gorm:"foreignKey:ID;references:OrganizationID"`
	Order        *ShpOrder          `gorm:"foreignKey:ID;references:OrderID"`
	Detail       *ShpOrderDetail    `gorm:"foreignKey:ID;references:DetailID"`
	Details      []ShpOrderDetail   `gorm:"foreignKey:OrderID;references:OrderID"`
	Products     []ShpServiceDetail `gorm:"foreignKey:ServiceID;references:ID"`
	Logs         []ShpLog           `gorm:"foreignKey:ServiceID;references:ID"`

	Refund *ShpRefund `gorm:"foreignKey:ServiceID;references:ID"`
}

func (ShpService) TableName() string {
	return TableShpService
}

const (
	ShpServiceOfTypeUnReceipt = "un_receipt"
	ShpServiceOfTypeRefund    = "refund"
	ShpServiceOfTypeExchange  = "exchange"

	ShpServiceOfStatusPending     = "pending"
	ShpServiceOfStatusUser        = "user"
	ShpServiceOfStatusOrg         = "org"
	ShpServiceOfStatusConfirmUser = "confirm_user"
	ShpServiceOfStatusConfirmOrg  = "confirm_org"
	ShpServiceOfStatusFinish      = "finish"
	ShpServiceOfStatusClosed      = "closed"

	ShpServiceOfResultAgree  = "agree"
	ShpServiceOfResultRefuse = "refuse"
)

type ShpServiceOfShipment struct {
	Company string `json:"company"`
	No      string `json:"no"`
	Remark  string `json:"remark"`
}
