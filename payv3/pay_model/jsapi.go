package pay_model

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type PrepayWithRequestPaymentResponse struct {

	// 预支付交易会话标识
	PrepayId string `json:"prepayId"`

	// 时间戳
	TimeStamp string `json:"timeStamp"`
	// 随机字符串
	NonceStr string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package string `json:"package"`
	// 签名方式
	SignType string `json:"signType"`
	// 签名
	PaySign string `json:"paySign"`

	// 二维码链接
	CodeUrl string `json:"codeUrl"`
	// 支付跳转链接
	H5Url string `json:"h5Url"`

	// 本地数据库 订单id
	OrderId string `json:"orderId"`
	// 本地数据库 订单号
	OutTradeNo string
}
type QueryOrderByIdRequest struct {
	TransactionId *string `json:"transaction_id"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}
