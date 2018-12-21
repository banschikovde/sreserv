package cache

type Cache interface {
	Set(key, value string)
	Get(key string)
}
