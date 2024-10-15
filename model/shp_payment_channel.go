package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const (
	TableShpPaymentChannel = "shp_payment_channel"
)

type ShpPaymentChannel struct {
	ID             uint           `gorm:"column:id;primaryKey"`       // 主键ID
	Platform       uint16         `gorm:"column:platform"`            // 平台
	CliqueID       *string        `gorm:"column:clique_id"`           // 集团ID
	OrganizationID *string        `gorm:"column:organization_id"`     // 组织ID
	Name           string         `gorm:"column:name"`                // 名称
	Channel        string         `gorm:"column:channel"`             // 支付渠道
	Key            string         `gorm:"column:key"`                 // KEY
	Secret         string         `gorm:"column:secret"`              // 密钥
	IsDebug        uint8          `gorm:"column:is_debug"`            // 是否调试：1=是；2=否
	Ext            map[string]any `gorm:"column:ext;serializer:json"` // 扩展信息
	Order          uint8          `gorm:"column:order"`               // 序号：正序
	IsEnable       uint8          `gorm:"column:is_enable"`           // 是否启用：1=是；2=否；
	CreatedAt      carbon.Carbon  `gorm:"column:created_at"`          // 创建时间
	UpdatedAt      carbon.Carbon  `gorm:"column:updated_at"`          // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`          // 删除时间
}

func (ShpPaymentChannel) TableName() string {
	return TableShpPaymentChannel
}

type ShpPaymentChannelOfExtPayPal struct {
	URL struct {
		Return string `mapstructure:"return"`
		Cancel string `mapstructure:"cancel"`
	} `mapstructure:"url"`
}
