package miniprogram

import (
	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/lutil/netutil"
)

type GenerateSchemeReq struct {
	JumpWxa    JumpWxa
	ExpireTime types.LocalDateTime // 和 DayNum 二选一
	DayNum     int                 //有效期最长为30天，可以自定义指定 过期天数，或者用ExpireTime指定具体的过期时间点
}
type GenerateSchemeReqParam struct {
	JumpWxa        JumpWxa `json:"jump_wxa,omitempty"`
	ExpireTime     int64   `json:"expire_time,omitempty"`     //到期失效的 scheme 码的失效时间，为 Unix 时间戳。生成的到期失效 scheme 码在该时间前有效。最长有效期为30天。is_expire 为 true 且 expire_type 为 0 时必填
	ExpireType     int     `json:"expire_type,omitempty"`     //默认值0，到期失效的 scheme 码失效类型，失效时间：0，失效间隔天数：1
	ExpireInterval int     `json:"expire_interval,omitempty"` //到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为30天。is_expire 为 true 且 expire_type 为 1 时必填
}

type JumpWxa struct {
	Path       string `json:"path,omitempty"`        //通过 scheme 码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query。path 为空时会跳转小程序主页
	Query      string `json:"query,omitempty"`       //通过 scheme 码进入小程序时的 query，最大1024个字符
	EnvVersion string `json:"env_version,omitempty"` //默认值"release"。要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"
}

type GenerateSchemeResp struct {
	ErrCode  int    `json:"errcode"`  // 错误码，请求失败时返回, 0 表示成功
	ErrMsg   string `json:"errmsg"`   // 错误信息，请求失败时返回
	Openlink string `json:"openlink"` // 生成的小程序 scheme 码
}

// GenerateScheme 获取加密scheme码
// 该接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-scheme/api_generatescheme.html
func (c MiniProgramConfig) GenerateScheme(req GenerateSchemeReq) (GenerateSchemeResp, error) {
	url := "https://api.weixin.qq.com/wxa/generatescheme"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GenerateSchemeResp{}, err
	}
	url += "?access_token=" + accessToken

	var data = GenerateSchemeReqParam{
		JumpWxa:        req.JumpWxa,
		ExpireTime:     0,
		ExpireType:     0,
		ExpireInterval: 0,
	}
	if req.DayNum > 0 {
		data.ExpireInterval = req.DayNum
		data.ExpireType = 1
	} else {
		data.ExpireTime = req.ExpireTime.ToGoTime().Unix()
		data.ExpireType = 0
	}
	result, err := netutil.PostJsonOk[GenerateSchemeResp](url, data)
	if err != nil {
		return GenerateSchemeResp{}, err
	}
	return result, nil
}

type GenerateNFCSchemeReq struct {
	JumpWxa JumpWxa `json:"jump_wxa,omitempty"`
	ModelId string  `json:"model_id"`
	Sn      string  `json:"sn"`
}

// GenerateNFCScheme 获取NFC的小程序scheme
// 该接口用于获取用于 NFC 的小程序 scheme 码，适用于 NFC 拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-scheme/api_generatenfcscheme.html
func (c MiniProgramConfig) GenerateNFCScheme(req GenerateNFCSchemeReq) (GenerateSchemeResp, error) {
	url := "https://api.weixin.qq.com/wxa/generatenfcscheme"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GenerateSchemeResp{}, err
	}
	url += "?access_token=" + accessToken

	result, err := netutil.PostJsonOk[GenerateSchemeResp](url, req)
	if err != nil {
		return GenerateSchemeResp{}, err
	}
	return result, nil
}

type QuerySchemeReq struct {
	Scheme    string `json:"scheme"`     //小程序 scheme 码。支持加密 scheme 和明文 scheme
	QueryType int    `json:"query_type"` //查询类型。默认值0，查询 scheme 码信息：0， 查询每天剩余访问次数：1
}

// QueryScheme 查询scheme码
// 永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-scheme/api_queryscheme.html
func (c MiniProgramConfig) QueryScheme(req QuerySchemeReq) (GetQRCodeResp, error) {
	url := "https://api.weixin.qq.com/wxa/queryscheme"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GetQRCodeResp{}, err
	}
	url += "?access_token=" + accessToken

	body, result, err := netutil.PostJsonByteOk[GetQRCodeResp](url, req)
	if err != nil {
		return GetQRCodeResp{}, err
	}
	if body == nil {
		return result, nil
	}
	result.Buffer = body
	return result, nil
}
