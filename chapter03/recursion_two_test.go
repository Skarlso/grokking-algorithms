package chapter03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	got := Factorial(5)
	assert.Equal(t, 120, got)
}
