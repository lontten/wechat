package payv3

import (
	"context"

	"github.com/lontten/wechat/payv3/pay_model"
	"github.com/lontten/wechat/wxutil"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type MpPayService struct {
	payConfig PayConfig
	client    jsapi.JsapiApiService
}

// GetMpPayClient 微信小程序支付,调用 微信APP 支付
func (pc PayConfig) GetMpPayClient() MpPayService {
	svc := jsapi.JsapiApiService{Client: pc.coreClient}
	return MpPayService{
		payConfig: pc,
		client:    svc,
	}
}

// CreateOrder 小程序调起支付
// https://pay.weixin.qq.com/doc/v3/merchant/4012791897
func (p MpPayService) CreateOrder(o pay_model.Order) (pay_model.PrepayWithRequestPaymentResponse, error) {
	var res pay_model.PrepayWithRequestPaymentResponse
	ctx := context.Background()
	var req = jsapi.PrepayRequest{
		Appid:       &p.payConfig.Appid,
		Mchid:       &p.payConfig.Mchid,
		Description: &o.Description,
		OutTradeNo:  &o.OutTradeNo,
		NotifyUrl:   &p.payConfig.NotifyUrl,
		Amount: &jsapi.Amount{
			Total: core.Int64(100),
		},
		Payer: &jsapi.Payer{
			Openid: &o.Openid,
		},
		Attach:        &o.Attach,
		TimeExpire:    o.TimeExpire,
		GoodsTag:      &o.GoodsTag,
		SupportFapiao: o.SupportFapiao,
		Detail:        o.Detail,
		SceneInfo:     o.SceneInfo,
		SettleInfo:    o.SettleInfo,
	}
	resp, _, err := p.client.PrepayWithRequestPayment(ctx, req)
	if err != nil {
		return res, err
	}
	res = pay_model.PrepayWithRequestPaymentResponse{
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.Package,
		SignType:  resp.SignType,
		PaySign:   resp.PaySign,
	}
	return res, err
}

// CreateOrderEasy 小程序调起支付
// https://pay.weixin.qq.com/doc/v3/merchant/4012791897
func (p MpPayService) CreateOrderEasy(o pay_model.EasyOrder) (pay_model.PrepayWithRequestPaymentResponse, error) {
	var res pay_model.PrepayWithRequestPaymentResponse
	ctx := context.Background()
	var req = jsapi.PrepayRequest{
		Appid:       &p.payConfig.Appid,
		Mchid:       &p.payConfig.Mchid,
		Description: &o.Title,
		OutTradeNo:  &o.OutTradeNo,
		NotifyUrl:   &p.payConfig.NotifyUrl,
		Amount: &jsapi.Amount{
			Total: core.Int64(wxutil.WxMoneyToFen(o.Money)),
		},
		Payer: &jsapi.Payer{
			Openid: &o.Openid,
		},
		Attach: &o.Attach,
	}
	resp, _, err := p.client.PrepayWithRequestPayment(ctx, req)
	if err != nil {
		return res, err
	}
	res = pay_model.PrepayWithRequestPaymentResponse{
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.Package,
		SignType:  resp.SignType,
		PaySign:   resp.PaySign,
	}
	return res, err
}

// QueryById 微信支付订单号查询订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791899
// wxOderId 微信支付 官方生成的 订单id
func (p MpPayService) QueryById(wxOderId string) (*payments.Transaction, error) {
	ctx := context.Background()
	resp, _, err := p.client.QueryOrderById(ctx, jsapi.QueryOrderByIdRequest{
		TransactionId: &wxOderId,
		Mchid:         &p.payConfig.Mchid,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryByOutTradeNo 商户订单号查询订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791900
// outTradeNo 微信支付 商户自定义订单号
func (p MpPayService) QueryByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	ctx := context.Background()
	resp, _, err := p.client.QueryOrderByOutTradeNo(ctx, jsapi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: &outTradeNo,
		Mchid:      &p.payConfig.Mchid,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CloseOrder 关闭订单
//
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
// https://pay.weixin.qq.com/doc/v3/merchant/4012791901
// outTradeNo 微信支付 商户自定义订单号
func (p MpPayService) CloseOrder(outTradeNo string) error {
	ctx := context.Background()
	_, err := p.client.CloseOrder(ctx, jsapi.CloseOrderRequest{
		OutTradeNo: &outTradeNo,
		Mchid:      &p.payConfig.Mchid,
	})
	if err != nil {
		return err
	}
	return nil
}
