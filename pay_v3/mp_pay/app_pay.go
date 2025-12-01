package mp_pay

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/app"
)

type AppPayService struct {
	client app.AppApiService
}

// APP支付,调用 微信APP 支付
func (pc PayConfig) AppPayClient() (AppPayService, error) {
	var payService AppPayService
	err := pc.InitClient()
	if err != nil {
		return payService, err
	}
	svc := app.AppApiService{Client: pc.coreClient}

	return AppPayService{
		client: svc,
	}, nil
}
