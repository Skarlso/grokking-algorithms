package chapter08

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	m1 := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
	}
	m2 := map[string]struct{}{
		"a": {},
		"d": {},
		"c": {},
	}
	expect := map[string]struct{}{
		"a": {},
		"c": {},
		"d": {},
	}
	got := intersect(m1, m2)
	assert.Equal(t, expect, got)
	m1 = map[string]struct{}{
		"a": {},
		"b": {},
	}
	m2 = map[string]struct{}{
		"a": {},
		"d": {},
		"c": {},
	}
	expect = map[string]struct{}{
		"a": {},
	}
	got = intersect(m1, m2)
	assert.Equal(t, expect, got)
}

func TestStates(t *testing.T) {
	final := States()
	var keys []string
	for k := range final {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	assert.Equal(t, []string{"kfive", "kone", "kthree", "ktwo"}, keys)
}
