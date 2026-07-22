package official_account

import (
	"fmt"

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
