package chapter05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
	// Create a new `[]byte` based data cache.
	cache := NewCache[[]byte]()
	got := GetPage(cache, "new")
	assert.Equal(t, []byte("this is generated data"), got)
	cache["new2"] = []byte("this is cached data")
	got = GetPage(cache, "new2")
	assert.Equal(t, []byte("this is cached data"), got)
	// Create a new `string` based data cache.
	cache2 := NewCache[string]()
	got2 := GetPage(cache2, "new")
	assert.Equal(t, "this is generated data", got2)
}
