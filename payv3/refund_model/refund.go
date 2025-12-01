package refund_model

type CreateRequest struct {
	// 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数
	SubMchid *string `json:"sub_mchid,omitempty"`
	// 原支付交易对应的微信订单号
	TransactionId *string `json:"transaction_id,omitempty"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 若商户传入，会在下发给用户的退款消息中体现退款原因
	Reason *string `json:"reason,omitempty"`
	// 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）  枚举值： - AVAILABLE：可用余额账户    * `AVAILABLE` - 可用余额
	FundsAccount *ReqFundsAccount `json:"funds_account,omitempty"`
	// 订单金额信息
	Amount *AmountReq `json:"amount"`
	// 指定商品退款需要传此参数，其他场景无需传递
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称
	GoodsName *string `json:"goods_name,omitempty"`
	// 商品单价金额，单位为分
	UnitPrice *int64 `json:"unit_price"`
	// 商品退款金额，单位为分
	RefundAmount *int64 `json:"refund_amount"`
	// 对应商品的退货数量
	RefundQuantity *int64 `json:"refund_quantity"`
}

// ReqFundsAccount * `AVAILABLE` - 可用余额, 仅对老资金流商户适用，指定从可用余额账户出资
type ReqFundsAccount string

// AmountReq
type AmountReq struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）。 同时指定多个账户出资退款的使用场景需要满足以下条件：1、未开通退款支出分离产品功能；2、订单属于分账订单，且分账处于待分账或分账中状态。 参数传递需要满足条件：1、基本账户可用余额出资金额与基本账户不可用余额出资金额之和等于退款金额；2、账户类型不能重复。 上述任一条件不满足将返回错误
	From []FundsFromItem `json:"from,omitempty"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total *int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

// FundsFromItem
type FundsFromItem struct {
	// 下面枚举值多选一。 枚举值： AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额
	Account *Account `json:"account"`
	// 对应账户出资金额
	Amount *int64 `json:"amount"`
}

// Account * `AVAILABLE` - 可用余额, 多账户资金准备退款可用余额出资账户类型 * `UNAVAILABLE` - 不可用余额, 多账户资金准备退款不可用余额出资账户类型
type Account string
