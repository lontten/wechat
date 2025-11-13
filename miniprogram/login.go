package miniprogram

import (
	"github.com/lontten/lcore/v2/netutil"
)

type Code2SessionResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回。
	OpenID     string `json:"openid"`      // 用户唯一标识
}

// Code2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (c MiniProgramConfig) Code2Session(code string) (Code2SessionResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session"
	url += "?appid=" + c.Appid
	url += "&secret=" + c.Secret
	url += "&js_code=" + code
	url += "&grant_type=" + "authorization_code"

	return netutil.Get[Code2SessionResp](url)
}
