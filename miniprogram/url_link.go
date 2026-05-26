package miniprogram

import (
	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/lutil/netutil"
)

type CloudBase struct {
	Env    string `json:"env"`    //云开发环境，必填
	Domain string `json:"domain"` //静态网站自定义域名，不填则使用默认域名

	Path          string `json:"path,omitempty"`           //云开发静态网站 H5 页面路径，不可携带 query
	Query         string `json:"query,omitempty"`          //云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	ResourceAppid string `json:"resource_appid,omitempty"` //第三方批量代云开发时必填，表示创建该 env 的 appid （小程序/第三方平台）
}

type GenerateUrlLinkReqParam struct {
	Path  string `json:"path,omitempty"`  //通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Query string `json:"query,omitempty"` //通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%

	ExpireType     int   `json:"expire_type,omitempty"`     //默认值0.小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
	ExpireTime     int64 `json:"expire_time,omitempty"`     //到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为30天。expire_type 为 0 必填
	ExpireInterval int   `json:"expire_interval,omitempty"` //到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为30天。expire_type 为 1 必填

	CloudBase *CloudBase `json:"cloud_base,omitempty"`

	EnvVersion string `json:"env_version,omitempty"` //默认值"release"。要打开的小程序版本。正式版为 "release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
}

type GenerateUrlLinkReq struct {
	Path  string `json:"path,omitempty"`  //通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Query string `json:"query,omitempty"` //通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%

	CloudBase *CloudBase `json:"cloud_base,omitempty"`

	EnvVersion string `json:"env_version,omitempty"` //默认值"release"。要打开的小程序版本。正式版为 "release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。

	ExpireTime types.LocalDateTime // 和 DayNum 二选一
	DayNum     int                 //有效期最长为30天，可以自定义指定 过期天数，或者用ExpireTime指定具体的过期时间点
}

type GenerateUrlLinkResp struct {
	ErrCode int    `json:"errcode"`  // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`   // 错误信息，请求失败时返回
	UrlLink string `json:"url_link"` // 生成的小程序 URL Link
}

func (r GenerateUrlLinkResp) Ok() bool {
	return r.ErrCode == 0
}

// GenerateUrlLink 获取加密URLLink
// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。目前仅针对国内非个人主体的小程序开放。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-link/api_generateurllink.html
func (c MiniProgramConfig) GenerateUrlLink(req GenerateUrlLinkReq) (GenerateUrlLinkResp, error) {
	url := "https://api.weixin.qq.com/wxa/generate_urllink"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GenerateUrlLinkResp{}, err
	}
	url += "?access_token=" + accessToken

	var data = GenerateUrlLinkReqParam{
		CloudBase:      req.CloudBase,
		EnvVersion:     req.EnvVersion,
		Path:           req.Path,
		Query:          req.Query,
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
	result, err := netutil.PostJsonOk[GenerateUrlLinkResp](url, data)
	if err != nil {
		return GenerateUrlLinkResp{}, err
	}
	return result, nil
}

type QueryUrlLinkReq struct {
	UrlLink   string `json:"url_link"`   //小程序加密 url_link
	QueryType int    `json:"query_type"` //查询类型。默认值0，查询 url_link 信息：0， 查询每天剩余访问次数：1
}

type QueryUrlLinkResp struct {
	ErrCode     int         `json:"errcode"`       // 错误码，请求失败时返回, 0 表示成功
	ErrMsg      string      `json:"errmsg"`        // 错误信息，请求失败时返回
	UrlLinkInfo UrlLinkInfo `json:"url_link_info"` // url_link 配置
	QuotaInfo   QuotaInfo   `json:"quota_info"`    // quota 配置
}

func (r QueryUrlLinkResp) Ok() bool {
	return r.ErrCode == 0
}

type UrlLinkInfo struct {
	Appid      string `json:"appid,omitempty"`       //小程序appid
	Path       string `json:"path,omitempty"`        //小程序页面路径
	Query      string `json:"query,omitempty"`       //小程序页面query
	CreateTime int64  `json:"create_time,omitempty"` //创建时间，为 Unix 时间戳
	ExpireTime int64  `json:"expire_time,omitempty"` //到期失效时间，为 Unix 时间戳，0 表示永久生效
	EnvVersion string `json:"env_version,omitempty"` //要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"

}
type QuotaInfo struct {
	RemainVisitQuota int64 `json:"remain_visit_quota"` // URL Scheme（加密+明文）/加密 URL Link 单天剩余访问次数
}

// QueryUrlLink 查询加密URLLink
// 该接口用于查询小程序加密 url_link 配置
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-link/api_queryurllink.html
func (c MiniProgramConfig) QueryUrlLink(req QueryUrlLinkReq) (QueryUrlLinkResp, error) {
	url := "https://api.weixin.qq.com/wxa/query_urllink"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return QueryUrlLinkResp{}, err
	}
	url += "?access_token=" + accessToken

	result, err := netutil.PostJsonOk[QueryUrlLinkResp](url, req)
	if err != nil {
		return QueryUrlLinkResp{}, err
	}
	return result, nil
}
