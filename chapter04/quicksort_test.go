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

func TestQuicksortV2(t *testing.T) {
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
			quicksort(tt.list, 0, len(tt.list)-1)
			assert.Equalf(t, tt.want, tt.list, "Quicksort(%v)", tt.list)
		})
	}
}

func BenchmarkQuicksort(b *testing.B) {
	list := make([]int, 0)
	for i := 1000; i > 0; i-- {
		list = append(list, i)
	}
	for i := 0; i < b.N; i++ {
		_ = Quicksort(list)
	}
}

func BenchmarkQuicksortV2(b *testing.B) {
	list := make([]int, 0)
	for i := 1000; i > 0; i-- {
		list = append(list, i)
	}
	for i := 0; i < b.N; i++ {
		quicksort(list, 0, 999)
	}
}
