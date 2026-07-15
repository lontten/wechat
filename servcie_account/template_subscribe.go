package service_account

import (
	"github.com/lontten/lutil/netutil"
)

// TemplateSubscribeReq
// 1.url和miniprogram都是非必填字段，若都不传则模板无跳转； 若都传会优先跳转至小程序
// 2.用户已关注公众号时消息下发到公众号会话，未关注时下发到服务通知
type TemplateSubscribeReq struct {
	ToUser     string `json:"touser"`        // 是	接收消息的用户openid
	TemplateId string `json:"template_id"`   // 是	订阅消息模板ID
	Url        string `json:"url,omitempty"` // 否	点击消息跳转链接(需ICP备案)

	MiniProgram *MiniProgram `json:"miniprogram,omitempty"` // 否	跳小程序配置

	Scene string `json:"scene"` // 是	订阅场景值
	Title string `json:"title"` // 是	消息标题(15字以内)

	Data Data `json:"data"` // 是	消息内容
}
type MiniProgram struct {
	Appid    string `json:"appid"`    // 小程序appid
	PagePath string `json:"pagepath"` // 小程序页面路径
}
type Data struct {
	Content *Content `json:"content,omitempty"` // 内容信息
}
type Content struct {
	Value string `json:"value,omitempty"` // 消息文本(200字内)
	Color string `json:"color,omitempty"` // 字体颜色
}

type TemplateSubscribeResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回
}

func (r TemplateSubscribeResp) Ok() bool {
	return r.ErrCode == 0
}

// TemplateSubscribe 发送一次性订阅消息
// 推送订阅模板消息给授权微信用户
// 服务号、移动应用
// 用户每次都需要主动授权
// 单一内容 (content 字段)
// 低频、轻量的服务通知
// https://developers.weixin.qq.com/doc/subscription/api/notify/subscribe/api_templatesubscribe.html
func (c ServiceAccountConfig) TemplateSubscribe(req TemplateSubscribeReq) (TemplateSubscribeResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/subscribe"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return TemplateSubscribeResp{}, err
	}
	url += "?access_token=" + accessToken

	return netutil.PostJsonOk[TemplateSubscribeResp](url, req)
}
