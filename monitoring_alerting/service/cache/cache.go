package cache

// Cache is an interface for cache implementations.
type Cache interface {
	// Set sets new key value pair.
	Set(key, value string)
	// Get returns value by key if it is exists
	Get(key string) (string, bool)
}
