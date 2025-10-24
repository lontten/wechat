package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

var CacheEngine MemberCache

type MemberCache struct {
	CacheEngine *ristretto.Cache
}

func InitMemberCache() MemberCache {
	// 初始化缓存：配置最大成本（可理解为最大内存占用估算）、缓冲区大小等
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // 最大成本（示例：1GB，根据实际需求调整）
		BufferItems: 64,      // 缓冲区大小（影响写入性能，建议 64-256）
	})
	if err != nil {
		panic(err)
	}
	CacheEngine = MemberCache{
		CacheEngine: c,
	}
	return CacheEngine
}
func (e MemberCache) Get(key string) (any, bool) {
	v, ok := e.CacheEngine.Get(key)
	return v, ok
}

func (e MemberCache) Set(key string, v any) {
	e.CacheEngine.SetWithTTL(key, v, 1, 0)
}
func (e MemberCache) SetExpire(key string, v any, expire int64) {
	e.CacheEngine.SetWithTTL(key, v, 1, time.Duration(expire)*time.Second)
}
