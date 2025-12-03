package payv3

import (
	"context"
	"log"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
)

// PayNotify 支付回调通知 通用
func (p WxPayService) PayNotify(req *http.Request) (*payments.Transaction, error) {
	ctx := context.Background()
	transaction := new(payments.Transaction)
	_, err := p.notifyClient.ParseNotifyRequest(ctx, req, transaction)
	if err != nil {
		log.Printf("❌ 回调处理失败: %s", err)
		return transaction, err
	}
	return transaction, nil
}

// RefundNotify 退款 回调通知 通用
func (p WxPayService) RefundNotify(req *http.Request) (map[string]any, error) {
	ctx := context.Background()
	content := make(map[string]any)
	_, err := p.notifyClient.ParseNotifyRequest(ctx, req, content)
	if err != nil {
		log.Printf("❌ 回调处理失败: %s", err)
		return content, err
	}
	return content, nil
}

func (p WxPayService) Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (p WxPayService) Fail(w http.ResponseWriter) {
	http.Error(w, "FAIL", http.StatusBadRequest)
}
