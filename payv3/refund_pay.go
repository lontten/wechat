package payv3

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

// Refund 退款申请
// https://pay.weixin.qq.com/doc/v3/merchant/4012791903
// outTradeNo 微信支付 商户自定义订单号
func (p WxPayService) Refund(req refunddomestic.CreateRequest) (*refunddomestic.Refund, error) {
	ctx := context.Background()
	resp, _, err := p.refundClient.Create(ctx, req)
	return resp, err
}

// QueryRefundByOutTradeNo 商户订单号查询 退款订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791904
// outTradeNo 微信支付 商户自定义 退款订单号
func (p WxPayService) QueryRefundByOutTradeNo(outTradeNo string) (*refunddomestic.Refund, error) {
	ctx := context.Background()
	resp, _, err := p.refundClient.QueryByOutRefundNo(ctx, refunddomestic.QueryByOutRefundNoRequest{
		OutRefundNo: &outTradeNo,
	})
	return resp, err
}
