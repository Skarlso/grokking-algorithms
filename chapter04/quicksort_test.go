package chapter04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			name: "sorts successfully",
			list: []int{4, 5, 6, 3, 7, 9, 1, 2},
			want: []int{1, 2, 3, 4, 5, 6, 7, 9},
		},
		{
			name: "sorts successfully with repeating items",
			list: []int{4, 5, 6, 6, 6, 9, 1, 2},
			want: []int{1, 2, 4, 5, 6, 6, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Quicksort(tt.list), "Quicksort(%v)", tt.list)
		})
	}
}
