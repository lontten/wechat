package cache

type CacheKeyer interface {
	CacheKey() string
}

type Cacher interface {
	Set(key string, value any)
	SetExpire(key string, value any, expire int64)
	Get(key string) (any, bool)
}
