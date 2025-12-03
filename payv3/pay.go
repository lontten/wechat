package payv3

import (
	"context"
	"errors"

	"github.com/lontten/lcore/v2/jsonutil"
	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/wechat/payv3/pay_model"
	"github.com/lontten/wechat/payv3/wxpay_type"
	"github.com/lontten/wechat/wxutil"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
)
import "github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
import "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"

// GetMpPayClient 微信小程序支付,调用 微信APP 支付
func (pc PayConfig) GetMpPayClient() WxPayService {
	svc := jsapi.JsapiApiService{Client: pc.coreClient}
	return WxPayService{
		payConfig:   pc,
		jsapiClient: svc,
	}
}

// CreateOrder 调起支付
// https://pay.weixin.qq.com/doc/v3/merchant/4012791897
func (p WxPayService) CreateOrder(typ wxpay_type.WxPayType, o pay_model.Order) (pay_model.PrepayWithRequestPaymentResponse, error) {
	var res pay_model.PrepayWithRequestPaymentResponse
	ctx := context.Background()

	switch typ {
	case wxpay_type.APP:
		req := jsonutil.ToObj[app.PrepayRequest](jsonutil.ToJsonStr(o.SceneInfo))
		req.Appid = &p.payConfig.Appid
		req.Mchid = &p.payConfig.Mchid
		req.NotifyUrl = &p.payConfig.NotifyUrl

		resp, _, err := p.appClient.PrepayWithRequestPayment(ctx, req)
		if err != nil {
			return res, err
		}
		res = pay_model.PrepayWithRequestPaymentResponse{
			PrepayId:   types.NilToZero(resp.PrepayId),
			TimeStamp:  types.NilToZero(resp.TimeStamp),
			NonceStr:   types.NilToZero(resp.NonceStr),
			Package:    types.NilToZero(resp.Package),
			Sign:       types.NilToZero(resp.Sign),
			OutTradeNo: o.OutTradeNo,
		}
		return res, err
	case wxpay_type.Pc_Web:
		req := jsonutil.ToObj[native.PrepayRequest](jsonutil.ToJsonStr(o.SceneInfo))
		req.Appid = &p.payConfig.Appid
		req.Mchid = &p.payConfig.Mchid
		req.NotifyUrl = &p.payConfig.NotifyUrl

		resp, _, err := p.nativeClient.Prepay(ctx, req)
		if err != nil {
			return res, err
		}

		res = pay_model.PrepayWithRequestPaymentResponse{
			CodeUrl:    types.NilToZero(resp.CodeUrl),
			OutTradeNo: o.OutTradeNo,
		}
		return res, err

	case wxpay_type.Phone_Web:
		req := jsonutil.ToObj[h5.PrepayRequest](jsonutil.ToJsonStr(o.SceneInfo))
		req.Appid = &p.payConfig.Appid
		req.Mchid = &p.payConfig.Mchid
		req.NotifyUrl = &p.payConfig.NotifyUrl

		resp, _, err := p.h5Client.Prepay(ctx, req)
		if err != nil {
			return res, err
		}
		res = pay_model.PrepayWithRequestPaymentResponse{
			H5Url:      types.NilToZero(resp.H5Url),
			OutTradeNo: o.OutTradeNo,
		}
		return res, err
	case wxpay_type.MiniProgram, wxpay_type.Wx_Web:
		req := jsonutil.ToObj[jsapi.PrepayRequest](jsonutil.ToJsonStr(o.SceneInfo))
		req.Payer = &jsapi.Payer{
			Openid: &o.Openid,
		}
		req.Appid = &p.payConfig.Appid
		req.Mchid = &p.payConfig.Mchid
		req.NotifyUrl = &p.payConfig.NotifyUrl

		resp, _, err := p.jsapiClient.PrepayWithRequestPayment(ctx, req)
		if err != nil {
			return res, err
		}
		res = pay_model.PrepayWithRequestPaymentResponse{
			PrepayId:   types.NilToZero(resp.PrepayId),
			TimeStamp:  types.NilToZero(resp.TimeStamp),
			NonceStr:   types.NilToZero(resp.NonceStr),
			Package:    types.NilToZero(resp.Package),
			SignType:   types.NilToZero(resp.SignType),
			Sign:       types.NilToZero(resp.PaySign),
			OutTradeNo: o.OutTradeNo,
		}
		return res, err
	default:
		return res, errors.New("微信支付类型错误")
	}

}

// CreateOrderEasy 调起支付
// https://pay.weixin.qq.com/doc/v3/merchant/4012791897
func (p WxPayService) CreateOrderEasy(typ wxpay_type.WxPayType, o pay_model.EasyOrder) (pay_model.PrepayWithRequestPaymentResponse, error) {
	var req = pay_model.Order{
		Description: o.Title,
		OutTradeNo:  o.OutTradeNo,
		Amount: pay_model.Amount{
			Total: wxutil.WxMoneyToFen(o.Money),
		},
		Openid: o.Openid,
		Attach: o.Attach,
	}
	return p.CreateOrder(typ, req)
}

// QueryById 微信支付订单号查询订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791899
// wxOderId 微信支付 官方生成的 订单id
func (p WxPayService) QueryById(wxOderId string) (*payments.Transaction, error) {
	ctx := context.Background()
	resp, _, err := p.jsapiClient.QueryOrderById(ctx, jsapi.QueryOrderByIdRequest{
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
func (p WxPayService) QueryByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	ctx := context.Background()
	resp, _, err := p.jsapiClient.QueryOrderByOutTradeNo(ctx, jsapi.QueryOrderByOutTradeNoRequest{
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
func (p WxPayService) CloseOrder(outTradeNo string) error {
	ctx := context.Background()
	_, err := p.jsapiClient.CloseOrder(ctx, jsapi.CloseOrderRequest{
		OutTradeNo: &outTradeNo,
		Mchid:      &p.payConfig.Mchid,
	})
	if err != nil {
		return err
	}
	return nil
}
