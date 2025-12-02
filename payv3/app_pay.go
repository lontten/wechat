package payv3

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/app"
)

type AppPayService struct {
	payConfig PayConfig
	client    app.AppApiService
}

// APP支付,调用 微信APP 支付
func (pc PayConfig) GetAppPayClient() (AppPayService, error) {
	svc := app.AppApiService{Client: pc.coreClient}

	return AppPayService{
		payConfig: pc,
		client:    svc,
	}, nil
}
