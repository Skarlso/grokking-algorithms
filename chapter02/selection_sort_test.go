package chapter02

import (
	"constraints"
	"testing"

	"github.com/stretchr/testify/assert"
)

// List this type serves no other purpose then experimenting with
// generic types.
type List[T constraints.Ordered] []T

func TestSelectionSortInts(t *testing.T) {
	list := List[int]{1, 4, 2, 3, 6, 5, 8, 7, 9}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := SelectionSort(list)
	assert.Equal(t, want, got)
}

func TestSelectionSortFloats(t *testing.T) {
	list := List[float64]{1.2, 4.3, 2.4, 3.5, 6.6, 5.7, 8.8, 7.9, 9.0}
	want := []float64{1.2, 2.4, 3.5, 4.3, 5.7, 6.6, 7.9, 8.8, 9.0}
	got := SelectionSort(list)
	assert.Equal(t, want, got)
}

func TestSelectionSortEmptyList(t *testing.T) {
	list := []int{}
	got := SelectionSort(list)
	assert.Nil(t, got)
}
