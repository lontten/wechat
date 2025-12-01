package pay_model

import "time"

type Order struct {
	// 必填：【商品描述】商品信息描述，用户微信账单的商品字段中可见，商户需传递能真实代表商品信息的描述，不能超过127个字符。
	Description *string `json:"description"`
	// 必填：商户订单号，要求6-32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。
	OutTradeNo *string `json:"out_trade_no"`
	// 选填：订单失效时间;
	// 若未指定支付结束时间，系统默认以下单时间为起始点计算时效；超过 7 天未支付的订单，无法再支付
	// 传递的支付结束时间需在下单时间的90天以内，如超过90天，微信支付会自动将该时间调整为下单时间后的第90天
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	// 选填：附加数据，总长度限制在128字符以内
	Attach *string `json:"attach,omitempty"`

	// 必填：【订单金额】订单金额信息
	Amount Amount `json:"amount"`

	// 必填：用户在商户appid下的唯一标识。
	Openid string `json:"openid,omitempty"`

	// 选填：商品标记，代金券或立减优惠功能的参数。
	GoodsTag *string `json:"goods_tag,omitempty"`
	// 指定支付方式
	//LimitPay []string `json:"limit_pay,omitempty"`
	// 选填：【电子发票入口开放标识】 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao *bool `json:"support_fapiao,omitempty"`

	// 选填：【优惠功能】 优惠功能
	Detail *Detail `json:"detail,omitempty"`
	// 选填：场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 选填：【结算信息】 结算信息
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

type Amount struct {
	// 必填：订单总金额，单位为分
	Total int64 `json:"total"`
	// 选填：CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
}

// Detail 优惠功能
type Detail struct {
	// 1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	CostPrice *int64 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity *int64 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice *int64 `json:"unit_price"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	PayerClientIp *string `json:"payer_client_ip"`
	// 商户端设备号
	DeviceId  *string    `json:"device_id,omitempty"`
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// StoreInfo 商户门店信息
type StoreInfo struct {
	// 商户侧门店编号
	Id *string `json:"id"`
	// 商户侧门店名称
	Name *string `json:"name,omitempty"`
	// 地区编码，详细请见微信支付提供的文档
	AreaCode *string `json:"area_code,omitempty"`
	// 详细的商户门店地址
	Address *string `json:"address,omitempty"`
}

// SettleInfo 【结算信息】 结算信息
type SettleInfo struct {
	// 是否指定分账,需要分账（传入true）,不需要分账（传入false或不传，默认为false）
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}
