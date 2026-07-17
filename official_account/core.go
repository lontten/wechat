package official_account

import "github.com/lontten/wechat/cache"

type OfficialAccountConfig struct {
	Appid  string `json:"appid"`  // 公众号 appId
	Secret string `json:"secret"` // 公众号 appSecret

	Cache cache.Cacher
}

func OfficialAccount(mpc OfficialAccountConfig) OfficialAccountConfig {
	if mpc.Cache == nil {
		mpc.Cache = cache.InitMemberCache()
	}
	return mpc
}

func (c OfficialAccountConfig) CacheKeyAccessToken() string {
	return "official_account_access_token_" + c.Appid
}

func (c OfficialAccountConfig) Get(key string) (any, bool) {
	v, ok := c.Cache.Get(key)
	return v, ok
}
func (c OfficialAccountConfig) Set(key string, v any) {
	c.Cache.Set(key, v)
}
func (c OfficialAccountConfig) SetExpire(key string, v any, expire int64) {
	c.Cache.SetExpire(key, v, expire)
}
