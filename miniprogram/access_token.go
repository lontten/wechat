package miniprogram

import (
	"fmt"

	"github.com/lontten/lcore/v2/netutil"
)

type GetAccessTokenResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	AccessToken string `json:"access_token"` // 接口调用凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
}

// GetAccessToken 获取接口调用凭据
// 获取小程序全局唯一后台接口调用凭据，token有效期为7200s，开发者需要进行妥善保存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func (p MiniProgramConfig) GetAccessToken() (GetAccessTokenResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token"
	url += "?appid=" + p.Appid
	url += "&grant_type=" + "client_credential"
	url += "&secret=" + p.Secret

	return netutil.Get[GetAccessTokenResp](url)
}

// GetStableAccessToken 获取稳定版接口调用凭据
// 获取小程序全局唯一后台接口调用凭据，token有效期为7200s，开发者需要进行妥善保存
// forceRefresh 默认使用 false。
// 1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；
// 2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
func (p MiniProgramConfig) GetStableAccessToken(forceRefresh ...bool) (GetAccessTokenResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/stable_token"
	url += "?appid=" + p.Appid
	url += "&grant_type=" + "client_credential"
	url += "&secret=" + p.Secret
	if len(forceRefresh) > 0 && forceRefresh[0] {
		url += "&force_refresh=true"
	}

	return netutil.PostJson[GetAccessTokenResp](url, nil)
}

func (p MiniProgramConfig) GetAccessTokenCache() (string, error) {
	v, ok := p.Get(p.CacheKeyAccessToken())
	if ok {
		return v.(string), nil
	}
	token, err := p.GetStableAccessToken()
	if err != nil {
		return "", err
	}
	if token.ErrCode != 0 {
		return "", fmt.Errorf("GetStableAccessToken,err:%v", token)
	}
	p.SetExpire(p.CacheKeyAccessToken(), token.AccessToken, int64(token.ExpiresIn)-200)
	return token.AccessToken, nil
}
