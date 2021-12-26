package chapter03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainRecurse(t *testing.T) {
	result := PlainRecurse(10, []int{})
	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result)
}
