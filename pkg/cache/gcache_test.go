package cache_test

import (
	"github.com/justtrackio/gosoline/pkg/mdl"
	"math/rand"
	"testing"
	"time"

	"github.com/justtrackio/gosoline/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestGCache_Set(t *testing.T) {
	assert.NotPanics(t, func() {
		cache := cache.NewGCache[string](1, 0, 1)

		cache.Set("key", "value")
	})
}

func TestCCache_Contains(t *testing.T) {
	assert.NotPanics(t, func() {
		cache := cache.NewGCache[string](1, 0, 1*time.Second)

		cache.Set("key", "value")
		assert.Equal(t, true, cache.Contains("key"))
	})
}

func TestCCache_Expire(t *testing.T) {
	assert.NotPanics(t, func() {
		cache := cache.NewGCache[string](1, 0, 0)

		cache.Set("key", "value")
		assert.Equal(t, false, cache.Contains("key"))
	})
}

var (
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
)

var benchRes []int

func BenchmarkGCache_SetPrimitive(b *testing.B) {
	c := cache.NewGCache[int](2_000_000, 0, 0)
	for i := 0; i < b.N; i++ {
		c.Set("key", 1)
	}
}

func BenchmarkCache_SetPrimitive(b *testing.B) {
	c := cache.New(2_000_000, 0, 0)
	for i := 0; i < b.N; i++ {
		c.Set("key", 1)
	}
}

func BenchmarkGCache_GetPrimitive(b *testing.B) {
	c := cache.NewGCache[int](2_000_000, 0, time.Minute)
	c.Set("key", 1)
	for i := 0; i < b.N; i++ {
		_, ok := c.Get("key")
		if !ok {
			b.Fail()
		}
	}
}

func BenchmarkCache_GetPrimitive(b *testing.B) {
	c := cache.New(2_000_000, 0, time.Minute)
	c.Set("key", 1)
	for i := 0; i < b.N; i++ {
		_, ok := c.Get("key")
		if !ok {
			b.Fail()
		}
	}
}

type test struct {
	Int    int
	String string
	Sl     []int
	Ptr    *time.Time
}

var testStruct = test{
	Int:    1,
	String: "hallo",
	Sl:     randomIntSlice(100),
	Ptr:    mdl.Box(time.Now()),
}

func BenchmarkGCache_SetStruct(b *testing.B) {
	c := cache.NewGCache[test](2_000_000, 0, 0)
	for i := 0; i < b.N; i++ {
		c.Set("key", testStruct)
	}
}

func BenchmarkCache_SetStruct(b *testing.B) {
	c := cache.New(2_000_000, 0, 0)
	for i := 0; i < b.N; i++ {
		c.Set("key", testStruct)
	}
}

func BenchmarkGCache_GetStruct(b *testing.B) {
	c := cache.NewGCache[test](2_000_000, 0, time.Minute)
	c.Set("key", testStruct)
	for i := 0; i < b.N; i++ {
		_, ok := c.Get("key")
		if !ok {
			b.Fail()
		}
	}
}

func BenchmarkCache_GetStruct(b *testing.B) {
	c := cache.New(2_000_000, 0, time.Minute)
	c.Set("key", testStruct)
	for i := 0; i < b.N; i++ {
		_, ok := c.Get("key")
		if !ok {
			b.Fail()
		}
	}
}

func randomIntSlice(length int) []int {
	var res []int
	for i := 0; i < length; i++ {
		res = append(res, r.Int())
	}

	return res
}
