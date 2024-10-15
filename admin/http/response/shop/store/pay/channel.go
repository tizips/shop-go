package pay

type ToChannelOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Channel   string `json:"channel"`
	Key       string `json:"key"`
	IsDebug   uint8  `json:"is_debug"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToChannelOfInformation struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Channel   string             `json:"channel"`
	Key       string             `json:"key"`
	Secret    string             `json:"secret"`
	IsDebug   uint8              `json:"is_debug"`
	PayPal    *ToChannelOfPayPal `json:"paypal,omitempty"`
	Order     uint8              `json:"order"`
	IsEnable  uint8              `json:"is_enable"`
	CreatedAt string             `json:"created_at"`
}

type ToChannelOfPayPal struct {
	URL struct {
		Return string `json:"return"`
		Cancel string `json:"cancel"`
	} `json:"url"`
}
