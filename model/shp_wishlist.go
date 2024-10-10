package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpWishlist = "shp_wishlist"
)

type ShpWishlist struct {
	ID             uint           `gorm:"column:id;primaryKey"`   // 主键ID
	Platform       uint16         `gorm:"column:platform"`        // 平台
	CliqueID       *string        `gorm:"column:clique_id"`       // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"` // 组织ID
	UserID         string         `gorm:"column:user_id"`         // 用户ID
	ProductID      string         `gorm:"column:product_id"`      // 产品ID
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`      // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`      // 删除时间

	Product *ShpProduct `gorm:"foreignKey:ID;references:ProductID"`
	SKU     *ShpSku     `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (ShpWishlist) TableName() string {
	return TableShpWishlist
}
