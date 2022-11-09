package memcache

import (
	"bookhub/internal/config"
	"bookhub/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

type Cache[Key comparable, T storage.CacheAble] struct {
	client     *memcache.Client
	expiration int32
}

func New[Key comparable, T storage.CacheAble](c *config.Memcached) *Cache[Key, T] {
	var cache Cache[Key, T]
	cache.client = memcache.New(c.URL)
	cache.expiration = c.Expiration
	return &cache
}

func (ch *Cache[Key, T]) Set(key string, item T) error {
	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("Marshal: %w", err)
	}

	chItem := memcache.Item{
		Key:        key,
		Value:      data,
		Expiration: ch.expiration,
	}
	if err = ch.client.Add(&chItem); err != nil {
		return fmt.Errorf("cache add: %w", err)
	}
	return nil
}
func (ch *Cache[Key, T]) Get(key string) (item T, err error) {
	chItem, err := ch.client.Get(key)
	if err != nil {
		return nil, fmt.Errorf("cache get: %w", err)
	}

	err = json.Unmarshal(chItem.Value, &item)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return item, nil
}
func (ch *Cache[Key, T]) Update(key string, item T) error {
	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("Marshal: %w", err)
	}

	chItem := memcache.Item{
		Key:        key,
		Expiration: ch.expiration,
		Value:      data,
	}
	err = ch.client.Replace(&chItem)
	if err == nil {
		return nil
	}
	if !errors.Is(err, memcache.ErrCacheMiss) && !errors.Is(err, memcache.ErrNotStored) {
		return fmt.Errorf("cache replace: %w", err)
	}

	err = ch.client.Add(&chItem)
	if err != nil {
		return fmt.Errorf("cache add: %w", err)
	}
	return nil
}
func (ch *Cache[Key, T]) Exist(key string) (exist bool, err error) {
	_, err = ch.client.Get(key)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("cache get: %w", err)
	}
	return true, nil
}
func (ch *Cache[Key, T]) Delete(key string) error {
	err := ch.client.Delete(key)
	if err == nil || errors.Is(err, memcache.ErrCacheMiss) {
		return nil
	}
	return fmt.Errorf("cache delete: %w", err)
}
