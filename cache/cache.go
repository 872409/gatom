package cache

type Cache struct {
	name    string
	storage map[string]interface{}
}

// func init() {
// 	fmt.Println("Cache init...")
// }

func Make(name string) *Cache {
	cache := &Cache{name: name, storage: map[string]interface{}{}}

	return cache
}

func (cache *Cache) Get(key string, def ...interface{}) interface{} {
	if value := cache.storage[key]; value != nil {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

func (cache *Cache) Save(key string, value interface{}) bool {
	cache.storage[key] = value
	return true
}
