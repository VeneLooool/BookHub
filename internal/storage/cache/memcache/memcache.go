package memcache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
)

type Cache struct {
	client *memcache.Client
}

func New(client *memcache.Client) *Cache {
	return &Cache{
		client: client,
	}
}

func (ch *Cache) Set(key any, item any) error {
	keyStr, err := ch.keyToString(key)
	if err != nil {
		return err
	}

	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("Marshal: %w", err)
	}

	chItem := memcache.Item{
		Key:        keyStr,
		Value:      data,
		Expiration: 60 * 10,
	}
	if err = ch.client.Add(&chItem); err != nil {
		return fmt.Errorf("Cache add: %w", err)
	}
	return nil
}
func (ch *Cache) Get(key any) (item any, err error) {
	keyStr, err := ch.keyToString(key)
	if err != nil {
		return nil, err
	}

	chItem, err := ch.client.Get(keyStr)
	if err != nil {
		return nil, fmt.Errorf("cache get: %w", err)
	}
	if err = json.Unmarshal(chItem.Value, &item); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return item, nil
}
func (ch *Cache) Exist(key any) (exist bool, err error) {
	keyStr, err := ch.keyToString(key)
	if err != nil {
		return false, err
	}

	_, err = ch.client.Get(keyStr)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("cache get: %w", err)
	}
	return true, nil
}
func (ch *Cache) Delete(key any) error {
	keyStr, err := ch.keyToString(key)
	if err != nil {
		return err
	}

	err = ch.client.Delete(keyStr)
	if err == nil || errors.Is(err, memcache.ErrCacheMiss) {
		return nil
	}
	return fmt.Errorf("cache delete: %w", err)
}

func (ch *Cache) keyToString(key any) (keyStr string, err error) {
	switch v := key.(type) {
	case int:
		keyStr = strconv.Itoa(v)
	case string:
		keyStr = v
	default:
		return "", errors.New("key type isn't comparable")
	}
	return keyStr, nil
}
