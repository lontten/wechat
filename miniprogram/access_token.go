package miniprogram

import (
	"github.com/lontten/lcore/v2"
)

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"` // 接口调用凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
}

// GetAccessToken 获取接口调用凭据
// 获取小程序全局唯一后台接口调用凭据，token有效期为7200s，开发者需要进行妥善保存
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func (p MiniProgramConfig) GetAccessToken() (GetAccessTokenResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token"
	url += "&grant_type=" + "client_credential"
	url += "?appid=" + p.Appid
	url += "&secret=" + p.Secret

	return lcore.Get[GetAccessTokenResp](url)
}
func (p MiniProgramConfig) GetAccessTokenCache() (string, error) {
	v, ok := p.Get(p.CacheKeyAccessToken())
	if ok {
		return v.(string), nil
	}
	token, err := p.GetAccessToken()
	if err != nil {
		return "", err
	}
	p.SetExpire(p.CacheKeyAccessToken(), token.AccessToken, int64(token.ExpiresIn)-200)
	return token.AccessToken, nil
}
