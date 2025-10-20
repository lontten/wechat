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

func (mp MiniProgramConfig) CacheKeyAccessToken() string {
	return "miniprogram_access_token_" + mp.Appid
}

func (mp MiniProgramConfig) Get(key string) (any, bool) {
	v, ok := mp.Cache.Get(key)
	return v, ok
}
func (mp MiniProgramConfig) Set(key string, v any) {
	mp.Cache.Set(key, v)
}
func (mp MiniProgramConfig) SetExpire(key string, v any, expire int64) {
	mp.Cache.SetExpire(key, v, expire)
}
