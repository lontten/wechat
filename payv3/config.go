package payv3

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

type PayConfig struct {
	Appid string `json:"appid"`

	Mchid                      string `json:"mchid"`                         // 商户号
	MchCertificateSerialNumber string `json:"mch_certificate_serial_number"` // 商户证书序列号
	PrivateKeyPath             string `json:"private_key_path"`              // 商户私钥文件路径 /path/to/merchant/apiclient_key.pem

	MchAPIv3Key string `json:"mch_api_v3_key"` // 商户 APIv3 密钥

	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`

	coreClient *core.Client
}
