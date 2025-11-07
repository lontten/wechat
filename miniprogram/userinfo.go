package miniprogram

import (
	"github.com/lontten/lcore/v2/netutil"
)

type GetPhoneNumberReq struct {
	Code   string `json:"code"`
	OpenID string `json:"openid"`
}
type GetPhoneNumberResp struct {
	ErrMsg    string    `json:"errmsg"`     // 错误信息，请求失败时返回
	ErrCode   int       `json:"errcode"`    // 错误码，请求失败时返回, 0 表示成功
	PhoneInfo PhoneInfo `json:"phone_info"` // 用户手机号信息

}
type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string    `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string    `json:"countryCode"`     // 区号
	Watermark       Watermark `json:"watermark"`       // 数据水印
}
type Watermark struct {
	Timestamp int64  `json:"timestamp"` // 用户获取手机号操作的时间戳
	Appid     string `json:"appid"`     // 小程序appid
}

// GetPhoneNumber 获取手机号
// 该接口用于将code换取用户手机号。 说明，每个code只能使用一次，code的有效期为5min。
// code 前端通过<button open-type="getPhoneNumber" bindgetphonenumber="getPhoneNumber"></button> 获取
// openid 选填，有开发者反馈，传入 openid 后反而报格式错误。建议不填。
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (p MiniProgramConfig) GetPhoneNumber(code string, openid ...string) (GetPhoneNumberResp, error) {
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	accessToken, err := p.GetAccessTokenCache()
	if err != nil {
		return GetPhoneNumberResp{}, err
	}
	url += "?access_token=" + accessToken
	var data = GetPhoneNumberReq{
		Code: code,
	}
	if len(openid) > 0 {
		data.OpenID = openid[0]
	}
	return netutil.PostJson[GetPhoneNumberResp](url, data)
}
