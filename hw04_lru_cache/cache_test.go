package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 1) // a => 1,
		c.Set("b", 2) // b => 2, a => 1,
		c.Set("c", 3) // c => 3, b => 2, a => 1,
		c.Set("d", 4) // d => 4, c => 3, b => 2,

		_, ok := c.Get("a")
		require.False(t, ok)

		c.Get("b")     // b => 2, d => 4, c => 3,
		c.Get("d")     // d => 4, b => 2, c => 3,
		c.Set("c", 7)  // c => 7, d => 4, b => 2,
		c.Set("f", 15) // f => 15, c => 7, d => 4,

		_, ok = c.Get("b")
		require.False(t, ok)

		c.Clear()
		_, ok = c.Get("f")
		require.False(t, ok)
		_, ok = c.Get("c")
		require.False(t, ok)
		_, ok = c.Get("d")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
