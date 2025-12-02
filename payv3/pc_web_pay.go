package payv3

import "github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"

type PcWebPayService struct {
	payConfig PayConfig
	client    native.NativeApiService
}

// GetPcWebPayClient PC端网页浏览器, 扫码支付
func (pc PayConfig) GetPcWebPayClient() PcWebPayService {
	svc := native.NativeApiService{Client: pc.coreClient}
	return PcWebPayService{
		payConfig: pc,
		client:    svc,
	}
}
