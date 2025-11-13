package miniprogram

import "github.com/lontten/wechat/cache"

type MiniProgramConfig struct {
	Appid  string `json:"appid"`  // 小程序 appId
	Secret string `json:"secret"` // 小程序 appSecret

	Cache cache.Cacher
}

func MiniProgram(mpc MiniProgramConfig) MiniProgramConfig {
	if mpc.Cache == nil {
		mpc.Cache = cache.InitMemberCache()
	}
	return mpc
}

func (c MiniProgramConfig) CacheKeyAccessToken() string {
	return "miniprogram_access_token_" + c.Appid
}

func (c MiniProgramConfig) Get(key string) (any, bool) {
	v, ok := c.Cache.Get(key)
	return v, ok
}
func (c MiniProgramConfig) Set(key string, v any) {
	c.Cache.Set(key, v)
}
func (c MiniProgramConfig) SetExpire(key string, v any, expire int64) {
	c.Cache.SetExpire(key, v, expire)
}
