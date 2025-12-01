package mp_pay

import "github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"

type PcWebPayService struct {
	client native.NativeApiService
}

// PC端网页浏览器, 扫码支付
func (pc PayConfig) PcWebPayClient() PcWebPayService {
	svc := native.NativeApiService{Client: pc.coreClient}
	return PcWebPayService{
		client: svc,
	}
}
