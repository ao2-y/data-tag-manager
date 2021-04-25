package inmemory

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache interface {
	Store(key string, obj interface{})
	Restore(key string) interface{}
	Delete(key string)
}

type inMemoryCache struct {
	cache *cache.Cache
}

func (i *inMemoryCache) Store(key string, obj interface{}) {
	i.cache.Set(key, obj, cache.DefaultExpiration)
}

func (i *inMemoryCache) Restore(key string) interface{} {
	obj, _ := i.cache.Get(key)
	return obj
}

func (i *inMemoryCache) Delete(key string) {
	i.cache.Delete(key)
}

func NewInMemoryCache() Cache {
	return &inMemoryCache{cache: cache.New(5*time.Minute, 10*time.Minute)}
}

type noCache struct {
}

func (n noCache) Store(key string, obj interface{}) {
	return
}

func (n noCache) Restore(key string) interface{} {
	return nil
}

func (n noCache) Delete(key string) {
	return
}

func NewNoCache() Cache {
	return &noCache{}
}
