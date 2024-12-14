package dbclients

import "sync"

var cache = struct {
	instances map[string]Client
	mu        *sync.Mutex
}{
	instances: make(map[string]Client),
	mu:        &sync.Mutex{},
}

func findInCache(key string) (Client, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	client, ok := cache.instances[key]
	return client, ok
}

func addToCache(key string, c Client) Client {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.instances[key] = c
	return c
}

func removeFromCache(key string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.instances, key)
}
