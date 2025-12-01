package mp_pay

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
)

type PhoneWebPayService struct {
	client h5.H5ApiService
}

// PhoneWebPayClient 手机浏览器网页,调用 微信APP 支付
func (pc PayConfig) PhoneWebPayClient() PhoneWebPayService {
	svc := h5.H5ApiService{Client: pc.coreClient}
	return PhoneWebPayService{
		client: svc,
	}
}
