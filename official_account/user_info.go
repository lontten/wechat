package official_account

import (
	"fmt"
	"time"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/lutil/netutil"
)

type GetFansResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	Total      int         `json:"total"`       // 关注该公众账号的总用户数
	Count      int         `json:"count"`       // 拉取的OPENID个数，最大值为10000
	Data       GetFansData `json:"data"`        // 列表数据，OPENID的列表
	NextOpenid string      `json:"next_openid"` // 拉取列表后一个用户的OPENID，为空表示列表结束
}

type GetFansData struct {
	Openid []string `json:"openid"` // 用户唯一ID列表
}

func (r GetFansResp) Ok() bool {
	return r.ErrCode == 0
}

// GetFans 获取关注用户列表
// 本接口用来获取账号的关注者列表，关注者列表由一串OpenID组成。
// 一次拉取最多10000个关注者的OpenID，可通过 next_openid 多次拉取。
// 最后一次返回时 next_openid 可能为空表示列表结束。
// https://developers.weixin.qq.com/doc/subscription/api/usermanage/userinfo/api_getfans.html
func (c OfficialAccountConfig) GetFans(nextOpenid ...string) (GetFansResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/get"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GetFansResp{}, err
	}
	url += "?access_token=" + accessToken
	if len(nextOpenid) > 0 && nextOpenid[0] != "" {
		url += "&next_openid=" + nextOpenid[0]
	}
	return netutil.Get[GetFansResp](url)
}

// GetAllFans 分页拉取全部关注用户 OpenID
// 每批（最多 10000）调用一次 handler；handler 返回 error 则中止
// total 为公众号关注者总数（接口返回的 total）
// https://developers.weixin.qq.com/doc/subscription/api/usermanage/userinfo/api_getfans.html
func (c OfficialAccountConfig) GetAllFans(handler func(openids []string, total int) error) error {
	var nextOpenid string
	for {
		resp, err := c.GetFans(nextOpenid)
		if err != nil {
			return err
		}
		if !resp.Ok() {
			return fmt.Errorf("GetFans,err:%v", resp)
		}
		if resp.Count > 0 {
			if err := handler(resp.Data.Openid, resp.Total); err != nil {
				return err
			}
		}
		if resp.NextOpenid == "" || resp.Count == 0 {
			return nil
		}
		nextOpenid = resp.NextOpenid
	}
}

// GetAllFansEach 分页拉取全部关注用户，每次回调一个 OpenID；handler 返回 error 则中止
// total 为公众号关注者总数（接口返回的 total）
// https://developers.weixin.qq.com/doc/subscription/api/usermanage/userinfo/api_getfans.html
func (c OfficialAccountConfig) GetAllFansEach(handler func(openid string, total int) error) error {
	return c.GetAllFans(func(openids []string, total int) error {
		for _, openid := range openids {
			if err := handler(openid, total); err != nil {
				return err
			}
		}
		return nil
	})
}

type UserInfo struct {
	Subscribe      int                 // 用户是否订阅该公众号标识，值为0时代表未关注，拉取不到其余信息
	Openid         string              // 用户的标识，对当前公众号唯一
	Language       string              // 【注意：该字段不再提供】用户的语言，简体中文为zh_CN
	SubscribeTime  types.LocalDateTime // 用户关注时间；若曾多次关注则取最后关注时间
	Unionid        string              // 只有在用户将公众号绑定到微信开放平台账号后才会出现
	Remark         string              // 公众号运营者对粉丝的备注
	Groupid        int                 // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int               // 用户被打上的标签ID列表
	SubscribeScene string              // 用户关注的渠道来源
	QrScene        int                 // 二维码扫码场景（开发者自定义）
	QrSceneStr     string              // 二维码扫码场景描述（开发者自定义）
}

type userInfoRaw struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int64  `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	Groupid        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

func toUserInfo(r userInfoRaw) UserInfo {
	info := UserInfo{
		Subscribe:      r.Subscribe,
		Openid:         r.Openid,
		Language:       r.Language,
		Unionid:        r.Unionid,
		Remark:         r.Remark,
		Groupid:        r.Groupid,
		TagidList:      r.TagidList,
		SubscribeScene: r.SubscribeScene,
		QrScene:        r.QrScene,
		QrSceneStr:     r.QrSceneStr,
	}
	if r.SubscribeTime != 0 {
		info.SubscribeTime = types.LocalDateTimeOfLoc(time.Unix(r.SubscribeTime, 0))
	}
	return info
}

type UserInfoResp struct {
	ErrCode int    // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string // 错误信息，请求失败时返回
	UserInfo
}

type userInfoRespRaw struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	userInfoRaw
}

func (r UserInfoResp) Ok() bool {
	return r.ErrCode == 0
}

func toUserInfoResp(r userInfoRespRaw) UserInfoResp {
	return UserInfoResp{
		ErrCode:  r.ErrCode,
		ErrMsg:   r.ErrMsg,
		UserInfo: toUserInfo(r.userInfoRaw),
	}
}

// UserInfo 获取用户基本信息
// 根据 OpenID 获取用户基本信息，包括关注时间、unionid、标签等。
// subscribe 为 0 时表示未关注，拉取不到其余信息。
// https://developers.weixin.qq.com/doc/subscription/api/usermanage/userinfo/api_userinfo.html
func (c OfficialAccountConfig) UserInfo(openid string, lang ...string) (UserInfoResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return UserInfoResp{}, err
	}
	url += "?access_token=" + accessToken
	url += "&openid=" + openid
	if len(lang) > 0 && lang[0] != "" {
		url += "&lang=" + lang[0]
	}
	raw, err := netutil.Get[userInfoRespRaw](url)
	if err != nil {
		return UserInfoResp{}, err
	}
	return toUserInfoResp(raw), nil
}

type BatchUserinfoItem struct {
	Openid string `json:"openid"`         // 是	用户的标识，对当前公众号唯一；必须是已关注的用户的 openid
	Lang   string `json:"lang,omitempty"` // 否	国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语，默认为 zh_CN
}

type BatchUserinfoReq struct {
	UserList []BatchUserinfoItem `json:"user_list"` // 是	用户列表，最多 100 条
}

type BatchUserinfoResp struct {
	ErrCode int    // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string // 错误信息，请求失败时返回

	UserInfoList []UserInfo // 用户列表
}

type batchUserinfoRespRaw struct {
	ErrCode      int           `json:"errcode"`
	ErrMsg       string        `json:"errmsg"`
	UserInfoList []userInfoRaw `json:"user_info_list"`
}

func (r BatchUserinfoResp) Ok() bool {
	return r.ErrCode == 0
}

func toBatchUserinfoResp(r batchUserinfoRespRaw) BatchUserinfoResp {
	list := make([]UserInfo, len(r.UserInfoList))
	for i, item := range r.UserInfoList {
		list[i] = toUserInfo(item)
	}
	return BatchUserinfoResp{
		ErrCode:      r.ErrCode,
		ErrMsg:       r.ErrMsg,
		UserInfoList: list,
	}
}

// BatchUserinfo 批量获取用户基本信息
// 最多支持一次拉取 100 条；所有 openid 必须是已关注的用户。
// https://developers.weixin.qq.com/doc/subscription/api/usermanage/userinfo/api_batchuserinfo.html
func (c OfficialAccountConfig) BatchUserinfo(userList []BatchUserinfoItem) (BatchUserinfoResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return BatchUserinfoResp{}, err
	}
	url += "?access_token=" + accessToken
	raw, err := netutil.PostJsonOk[batchUserinfoRespRaw](url, BatchUserinfoReq{UserList: userList})
	if err != nil {
		return BatchUserinfoResp{}, err
	}
	return toBatchUserinfoResp(raw), nil
}
