package pay_model

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type PrepayWithRequestPaymentResponse struct {
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}
type QueryOrderByIdRequest struct {
	TransactionId *string `json:"transaction_id"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}
