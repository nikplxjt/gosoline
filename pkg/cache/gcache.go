package cache

import (
	"time"

	"github.com/karlseguin/ccache"
)

type GCache[T any] interface {
	Contains(key string) bool
	Expire(key string) bool
	Get(key string) (*T, bool)
	Set(key string, value T)
	SetX(key string, value T, ttl time.Duration)
}

type gcache[T any] struct {
	base *ccache.Cache
	size int
	ttl  time.Duration
}

// Contains implements GCache
func (c *gcache[T]) Contains(key string) bool {
	_, ok := c.Get(key)

	return ok
}

// Expire implements GCache
func (c *gcache[T]) Expire(key string) bool {
	item := c.base.Get(key)

	if item == nil {
		return false
	}

	remaining := time.Until(item.Expires())
	item.Extend(-1 * remaining)

	return true
}

// Get implements GCache
func (c *gcache[T]) Get(key string) (*T, bool) {
	item := c.base.Get(key)

	if item == nil {
		return nil, false
	}

	if item.Expired() {
		return nil, false
	}

	val := item.Value().(T)

	return &val, true
}

// Set implements GCache
func (c *gcache[T]) Set(key string, value T) {
	c.base.Set(key, value, c.ttl)
}

// SetX implements GCache
func (c *gcache[T]) SetX(key string, value T, ttl time.Duration) {
	c.base.Set(key, value, ttl)
}

func NewGCache[T any](maxSize int64, pruneCount uint32, ttl time.Duration) GCache[T] {
	cache := &gcache[T]{
		size: 0,
		ttl:  ttl,
	}

	config := ccache.Configure()
	config.MaxSize(maxSize)
	config.ItemsToPrune(pruneCount)

	cache.base = ccache.New(config)

	return cache
}
