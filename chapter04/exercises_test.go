package chapter04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4})
	assert.Equal(t, 10, got)
}

func TestCountItems(t *testing.T) {
	got := CountItems[int, int]([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 5, got)
}
