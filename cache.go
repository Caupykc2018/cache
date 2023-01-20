package cache

type Cache struct {
	memory map[string]interface{}
}

func New() Cache {
	return Cache{make(map[string]interface{})}
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.memory[key] = value
}

func (cache *Cache) Get(key string) interface{} {
	return cache.memory[key]
}

func (cache *Cache) Delete(key string) {
	delete(cache.memory, key)
}
