package service_account

import "github.com/lontten/wechat/cache"

type ServiceAccountConfig struct {
	Appid  string `json:"appid"`  // 服务号 appId
	Secret string `json:"secret"` // 服务号 appSecret

	Cache cache.Cacher
}

func ServiceAccount(mpc ServiceAccountConfig) ServiceAccountConfig {
	if mpc.Cache == nil {
		mpc.Cache = cache.InitMemberCache()
	}
	return mpc
}

func (c ServiceAccountConfig) CacheKeyAccessToken() string {
	return "official_account_access_token_" + c.Appid
}

func (c ServiceAccountConfig) Get(key string) (any, bool) {
	v, ok := c.Cache.Get(key)
	return v, ok
}
func (c ServiceAccountConfig) Set(key string, v any) {
	c.Cache.Set(key, v)
}
func (c ServiceAccountConfig) SetExpire(key string, v any, expire int64) {
	c.Cache.SetExpire(key, v, expire)
}
