package miniprogram

import "github.com/lontten/lcore/v2"

type GetPhoneNumberReq struct {
	Code   string `json:"code"`   // 前端通过<button open-type="getPhoneNumber" bindgetphonenumber="getPhoneNumber"></button> 获取
	OpenID string `json:"openid"` // 用户唯一标识，非必填
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
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (p MiniProgramConfig) GetPhoneNumber(req GetPhoneNumberReq) (GetPhoneNumberResp, error) {
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	url += "?access_token=" + p.AccessToken
	return lcore.PostJson[GetPhoneNumberResp](url, req)
}
