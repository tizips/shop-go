package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpCart = "shp_cart"
)

type ShpCart struct {
	ID             uint           `gorm:"column:id;primaryKey"`                  // ID
	Platform       uint16         `gorm:"column:platform"`                       // 平台
	CliqueID       *string        `gorm:"column:clique_id;index"`                // 集团ID
	OrganizationID *string        `gorm:"column:organization_id;index"`          // 组织ID
	UserID         string         `gorm:"column:user_id;index"`                  // 用户ID
	ProductID      string         `gorm:"column:product_id;index"`               // 产品ID
	SkuID          string         `gorm:"column:sku_id;index"`                   // SkuID
	Code           string         `gorm:"column:code"`                           // 代码
	Specifications []string       `gorm:"column:specifications;serializer:json"` // 规格
	Name           string         `gorm:"column:name"`                           // 名称
	Picture        string         `gorm:"column:picture"`                        // 图片
	Price          uint           `gorm:"column:price"`                          // 价格
	Quantity       uint           `gorm:"column:quantity"`                       // 数量
	IsInvalid      uint8          `gorm:"column:is_invalid"`                     // 是否失效：1=是；2=否
	CreatedAt      carbon.Carbon  `gorm:"column:created_at;autoCreateTime"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at;autoUpdateTime"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`                     // 删除时间

	Product *ShpProduct `gorm:"foreignKey:ID;references:SKU"`
	SKU     *ShpSku     `gorm:"foreignKey:ID;references:SkuID"`
}

func (s *ShpCart) TableName() string {
	return TableShpCart
}
