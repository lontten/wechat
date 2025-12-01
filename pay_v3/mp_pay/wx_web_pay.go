package mp_pay

import "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"

type WxWebPayService struct {
	client jsapi.JsapiApiService
}

// WxWebPayClient 微信内部浏览器网页,调用 微信APP 支付
func (pc PayConfig) WxWebPayClient() WxWebPayService {
	svc := jsapi.JsapiApiService{Client: pc.coreClient}
	return WxWebPayService{
		client: svc,
	}
}
