package pay_v3

import (
	"context"
	"log"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
)

type NotifyService struct {
	client *notify.Handler
}

// GetNotifyClient 回调通知 通用
func (pc PayConfig) GetNotifyClient() (NotifyService, error) {
	var payService NotifyService
	// 3. 创建证书访问器和验证器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(pc.Mchid)
	verifier := verifiers.NewSHA256WithRSAVerifier(certVisitor)

	// 4. ✅ 使用新的 RSA 通知处理器（推荐方式）
	handler, err := notify.NewRSANotifyHandler(pc.MchAPIv3Key, verifier)
	if err != nil {
		return payService, err
	}

	return NotifyService{
		client: handler,
	}, nil
}

// PayNotify 支付回调通知 通用
func (p NotifyService) PayNotify(req *http.Request) (*payments.Transaction, error) {
	ctx := context.Background()
	transaction := new(payments.Transaction)
	_, err := p.client.ParseNotifyRequest(ctx, req, transaction)
	if err != nil {
		log.Printf("❌ 回调处理失败: %s", err)
		return transaction, err
	}
	return transaction, nil
}

// RefundNotify 退款 回调通知 通用
func (p NotifyService) RefundNotify(req *http.Request) (map[string]any, error) {
	ctx := context.Background()
	content := make(map[string]any)
	_, err := p.client.ParseNotifyRequest(ctx, req, content)
	if err != nil {
		log.Printf("❌ 回调处理失败: %s", err)
		return content, err
	}
	return content, nil
}

func (p NotifyService) Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (p NotifyService) Fail(w http.ResponseWriter) {
	http.Error(w, "FAIL", http.StatusBadRequest)
}
