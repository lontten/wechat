package mp_pay

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PayConfig struct {
	Appid string `json:"appid"` // 小程序 appId
	Mchid string `json:"mchid"` // 商户号

	MchAPIv3Key    string `json:"mch_api_v3_key"`   // 商户 APIv3 密钥
	PrivateKeyPath string `json:"private_key_path"` // 商户私钥文件路径 /path/to/merchant/apiclient_key.pem

	MchCertificateSerialNumber string `json:"mch_certificate_serial_number"` // 商户证书序列号

	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`

	coreClient *core.Client
}

func (pc PayConfig) InitClient() error {
	ctx := context.Background()
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(pc.PrivateKeyPath)
	if err != nil {
		return err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(pc.Mchid, pc.MchCertificateSerialNumber, mchPrivateKey, pc.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return err
	}
	pc.coreClient = client
	return nil
}
