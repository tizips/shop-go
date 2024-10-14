package queue

import "github.com/golang-module/carbon/v2"

const (
	ShopOrderPaid      = "shop_order_paid"
	ShopOrderRefund    = "shop_order_refund"
	ShopOrderClosed    = "shop_order_closed"
	ShopOrderReceived  = "shop_order_received"
	ShopOrderCompleted = "shop_order_completed"
	ShopOrderLog       = "shop_order_log"

	ShopPaymentSuccess = "shop_payment_success"
	ShopPaymentRefund  = "shop_payment_refund"

	ShopServiceAgree  = "shop_service_agree"
	ShopServiceFinish = "shop_service_finish"
	ShopServiceRefund = "shop_service_refund"
)

type ShopPaymentSuccessMessage struct {
	ID       string         `json:"id"`               // 支付ID
	No       string         `json:"no"`               // 第三方支付单号
	Channel  string         `json:"channel"`          // 支付渠道
	Total    uint           `json:"total"`            // 总额
	Payer    uint           `json:"payer"`            // 到账
	Currency string         `json:"currency"`         // 币种
	OpenID   string         `json:"openid,omitempty"` // 用户ID
	PaidAt   carbon.Carbon  `json:"paid_at"`          // 支付时间
	Ext      map[string]any `json:"ext,omitempty"`
}

type ShopPaymentRefundMessage struct {
	ID      string         `json:"id"`               // 支付ID
	Order   string         `json:"order"`            // 订单ID
	Service *string        `json:"service"`          // 售后ID
	Detail  *string        `json:"detail"`           // 明细ID
	Money   uint           `json:"money"`            // 金额「为空从表中查询」
	Reason  string         `json:"reason"`           // 退款原因
	OpenID  *string        `json:"openid,omitempty"` // 用户ID
	Ext     map[string]any `json:"ext,omitempty"`
}

type ShopOrderPaidMessage struct {
	ID    string `json:"id"`    // 支付ID
	No    string `json:"no"`    // 第三方支付单号
	Order string `json:"order"` // 订单ID
}

type ShopOrderRefundMessage struct {
	ID      string                           `json:"id"` // 订单号
	Detail  *string                          `json:"detail,omitempty"`
	Service *string                          `json:"service,omitempty"`
	Refund  uint                             `json:"refund,omitempty"`
	Reason  string                           `json:"reason"`
	Details []ShopOrderRefundOfDetailMessage `json:"details"` // 明细订单
}

type ShopOrderRefundOfDetailMessage struct {
	ID       string `json:"id"`       // 明细ID
	Quantity uint   `json:"quantity"` // 退款数量
	Refund   uint   `json:"refund"`   // 退款金额
}

type ShopOrderClosedMessage struct {
	UserID string `json:"user_id"`
	Order  string `json:"order"` // 订单ID
}

type ShopOrderLogMessage struct {
	Order     string        `json:"order"` // 订单ID
	Action    string        `json:"action"`
	Content   string        `json:"content"`
	CreatedAt carbon.Carbon `json:"created_at"`
}

type ShopServiceRefundMessage struct {
	ID    string `json:"id"`
	Money uint   `json:"money"`
}
