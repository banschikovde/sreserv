package cache

// Mem is a cache implementation in memory.
type Mem struct {
	pointer int
	size    int
	values  []kv
}

// NewMem initializes new memory cacher with fixed size.
func NewMem(size int) *Mem {
	return &Mem{pointer: 0, size: size, values: make([]kv, size)}
}

// Set sets new key, value pair.
func (m *Mem) Set(key, value string) {
	// if item with such key and value exists, return
	for _, kv := range m.values {
		if kv.key == key && kv.value == value {
			return
		}
	}

	m.values[m.pointer] = kv{key: key, value: value}

	m.pointer++
	if m.pointer >= m.size {
		m.pointer = 0
	}
}

// Get finds value by key and returns it if it exists.
func (m *Mem) Get(key string) (string, bool) {
	for _, kv := range m.values {
		if kv.key == key {
			return kv.value, true
		}
	}

	return "", false
}

// kv contains key and value.
type kv struct {
	key   string
	value string
}
