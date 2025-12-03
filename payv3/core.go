package payv3

import (
	"context"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WxPayService struct {
	payConfig  PayConfig
	coreClient *core.Client

	jsapiClient  jsapi.JsapiApiService
	appClient    app.AppApiService
	nativeClient native.NativeApiService
	h5Client     h5.H5ApiService

	refundClient refunddomestic.RefundsApiService
	notifyClient *notify.Handler
}

func InitClient(c PayConfig) (WxPayService, error) {
	ctx := context.Background()
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(c.PrivateKeyPath)
	if err != nil {
		return WxPayService{}, fmt.Errorf("加载私钥失败: %v", err)
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(c.Mchid, c.MchCertificateSerialNumber, mchPrivateKey, c.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return WxPayService{}, fmt.Errorf("初始化微信支付client失败: %v", err)
	}

	certVisitor := downloader.MgrInstance().GetCertificateVisitor(c.Mchid)
	verifier := verifiers.NewSHA256WithRSAVerifier(certVisitor)

	// 4. ✅ 使用新的 RSA 通知处理器（推荐方式）
	handler, err := notify.NewRSANotifyHandler(c.MchAPIv3Key, verifier)
	if err != nil {
		return WxPayService{}, err
	}

	return WxPayService{
		payConfig:    c,
		coreClient:   client,
		jsapiClient:  jsapi.JsapiApiService{Client: client},
		appClient:    app.AppApiService{Client: client},
		nativeClient: native.NativeApiService{Client: client},
		h5Client:     h5.H5ApiService{Client: client},
		refundClient: refunddomestic.RefundsApiService{Client: client},
		notifyClient: handler,
	}, nil
}
