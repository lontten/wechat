package official_account

import (
	"fmt"

	"github.com/lontten/lutil/netutil"
)

type GetAccessTokenResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	AccessToken string `json:"access_token"` // 接口调用凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
}

// GetAccessToken 获取接口调用凭据
// 本接口用于获取获取全局唯一后台接口调用凭据（Access Token），token 有效期为 7200 秒，开发者需要进行妥善保存，使用注意事项请参考此文档。
// 推荐使用 获取稳定版接口调用凭据
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getaccesstoken.html
func (c OfficialAccountConfig) GetAccessToken() (GetAccessTokenResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token"
	url += "?appid=" + c.Appid
	url += "&grant_type=" + "client_credential"
	url += "&secret=" + c.Secret

	return netutil.Get[GetAccessTokenResp](url)
}

type GetStableAccessTokenReq struct {
	Appid     string `json:"appid"`      // 公众号 appId
	Secret    string `json:"secret"`     // 公众号 appSecret
	GrantType string `json:"grant_type"` // 填写 client_credential
	// 默认使用 false。
	// 1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；
	// 2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
	ForceRefresh bool `json:"force_refresh"` //
}

// GetStableAccessToken 获取稳定版接口调用凭据
// 本接口用于获取获取全局唯一后台接口调用凭据（Access Token），token 有效期为 7200 秒，但此接口和 getAccessToken 互相隔离，且比其更加稳定，推荐使用此接口替代。
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getstableaccesstoken.html
func (c OfficialAccountConfig) GetStableAccessToken(forceRefresh ...bool) (GetAccessTokenResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/stable_token"
	var data = GetStableAccessTokenReq{
		Appid:     c.Appid,
		Secret:    c.Secret,
		GrantType: "client_credential",
	}
	if len(forceRefresh) > 0 {
		data.ForceRefresh = forceRefresh[0]
	}
	return netutil.PostJsonOk[GetAccessTokenResp](url, data)
}

func (c OfficialAccountConfig) GetAccessTokenCache() (string, error) {
	v, ok := c.Get(c.CacheKeyAccessToken())
	if ok {
		return v.(string), nil
	}
	token, err := c.GetStableAccessToken()
	if err != nil {
		return "", err
	}
	if token.ErrCode != 0 {
		return "", fmt.Errorf("GetStableAccessToken,err:%v", token)
	}
	c.SetExpire(c.CacheKeyAccessToken(), token.AccessToken, int64(token.ExpiresIn)-200)
	return token.AccessToken, nil
}
