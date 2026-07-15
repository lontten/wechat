package service_account

import (
	"github.com/lontten/lutil/netutil"
)

// SendTemplateMessageReq
// url和miniprogram都是非必填字段，若都不传则模板无跳转；
// 若都传，会优先跳转至小程序。
// 开发者可根据实际需要选择其中一种跳转方式即可。
// 当用户的微信客户端版本不支持跳小程序时，将会跳转至url。
type SendTemplateMessageReq struct {
	ToUser     string `json:"touser"`      // 是	接收消息的用户openid
	TemplateId string `json:"template_id"` // 是	订阅消息模板ID

	Url         string       `json:"url,omitempty"`         // 否	模板跳转链接（海外账号没有跳转能力
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"` // 否	跳小程序配置

	Data map[string]ValueContent `json:"data"` // 是	模板内容，需根据模板给定的格式给出（参考注意事项），格式形如 { "key1": { "value": any }, "key2": { "value": any } }

	ClientMsgId string `json:"client_msg_id,omitempty"` // 否	防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
}
type ValueContent struct {
	Value string `json:"value,omitempty"` // 消息文本
}

type SendTemplateMessageResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	MsgId int `json:"msgid"` // 消息id
}

// SendTemplateMessage 发送模板消息
// 推送订阅模板消息给授权微信用户
// 仅限认证后的服务号
// 用户无需每次授权，开发者可在模板规定场景下发送
// 多个自定义参数 (如 {{name01.DATA}} )
// 结构化的业务通知，如订单、物流、消费提醒等
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_sendtemplatemessage.html
func (c ServiceAccountConfig) SendTemplateMessage(req SendTemplateMessageReq) (SendTemplateMessageResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return SendTemplateMessageResp{}, err
	}
	url += "?access_token=" + accessToken

	return netutil.PostJsonOk[SendTemplateMessageResp](url, req)
}
