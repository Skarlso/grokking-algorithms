package chapter01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	listOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	intItem, err := BinarySearch(listOfInts, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, intItem)
	_, err = BinarySearch(listOfInts, 11)
	assert.EqualError(t, err, "not found")
	listOfFloats := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	floatItem, err := BinarySearch(listOfFloats, 3.3)
	assert.NoError(t, err)
	assert.Equal(t, 3.3, floatItem)
	_, err = BinarySearch([]int{1, 2, 3}, 5)
	assert.EqualError(t, err, "not found")
}
