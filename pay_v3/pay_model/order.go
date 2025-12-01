package pay_model

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type Order struct {
	// 必填：【商品描述】商品信息描述，用户微信账单的商品字段中可见，商户需传递能真实代表商品信息的描述，不能超过127个字符。
	Description string `json:"description"`
	// 必填：商户订单号，要求6-32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。
	OutTradeNo string `json:"out_trade_no"`

	// 必填：【订单金额】订单金额信息
	Amount Amount `json:"amount"`

	// 必填：用户在商户appid下的唯一标识。
	Openid string `json:"openid,omitempty"`

	// 选填：订单失效时间;
	// 若未指定支付结束时间，系统默认以下单时间为起始点计算时效；超过 7 天未支付的订单，无法再支付
	// 传递的支付结束时间需在下单时间的90天以内，如超过90天，微信支付会自动将该时间调整为下单时间后的第90天
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	// 选填：附加数据，总长度限制在128字符以内
	Attach string `json:"attach,omitempty"`

	// 选填：商品标记，代金券或立减优惠功能的参数。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 指定支付方式
	//LimitPay []string `json:"limit_pay,omitempty"`
	// 选填：【电子发票入口开放标识】 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao *bool `json:"support_fapiao,omitempty"`

	// 选填：【优惠功能】 优惠功能
	Detail *jsapi.Detail `json:"detail,omitempty"`
	// 选填：场景信息
	SceneInfo *jsapi.SceneInfo `json:"scene_info,omitempty"`
	// 选填：【结算信息】 结算信息
	SettleInfo *jsapi.SettleInfo `json:"settle_info,omitempty"`
}

type EasyOrder struct {
	// 必填：【商品描述】商品信息描述，用户微信账单的商品字段中可见，商户需传递能真实代表商品信息的描述，不能超过127个字符。
	Title string
	// 必填：商户订单号，要求6-32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。
	OutTradeNo string
	// 必填：【订单金额】订单金额信息
	Money decimal.Decimal
	// 必填：用户在商户appid下的唯一标识。
	Openid string
	// 选填：附加数据，总长度限制在128字符以内
	Attach string `json:"attach,omitempty"`
}

type Amount struct {
	// 必填：订单总金额，单位为分
	Total int64 `json:"total"`
	// 选填：CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
}
