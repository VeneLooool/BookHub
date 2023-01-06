package memcache

import (
	"github.com/VeneLooool/BookHub/internal/config"
	"github.com/bradfitz/gomemcache/memcache"
)

func New(c *config.Config) *memcache.Client {
	mc := memcache.New(c.Memcached.URL)
	return mc
}
