package mp_pay

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type RefundService struct {
	client refunddomestic.RefundsApiService
}

// RefundClient 退款 申请 通用
func (pc PayConfig) RefundClient() RefundService {
	svc := refunddomestic.RefundsApiService{Client: pc.coreClient}
	return RefundService{
		client: svc,
	}
}

// Refund 退款申请
// https://pay.weixin.qq.com/doc/v3/merchant/4012791903
// outTradeNo 微信支付 商户自定义订单号
func (p RefundService) Refund(req refunddomestic.CreateRequest) (*refunddomestic.Refund, error) {
	ctx := context.Background()
	resp, _, err := p.client.Create(ctx, req)
	return resp, err
}

// QueryByOutTradeNo 商户订单号查询 退款订单
// https://pay.weixin.qq.com/doc/v3/merchant/4012791904
// outTradeNo 微信支付 商户自定义 退款订单号
func (p RefundService) QueryByOutTradeNo(outTradeNo string) (*refunddomestic.Refund, error) {
	ctx := context.Background()
	resp, _, err := p.client.QueryByOutRefundNo(ctx, refunddomestic.QueryByOutRefundNoRequest{
		OutRefundNo: &outTradeNo,
	})
	return resp, err
}
