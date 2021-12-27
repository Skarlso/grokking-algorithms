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

func TestFindMax(t *testing.T) {
	got := FindMax([]int{1, 2, 5, 4, 1}, 0)
	assert.Equal(t, 5, got)
	got = FindMax([]int{1, 2, 5, 4, 1, 9, 12}, 0)
	assert.Equal(t, 12, got)
}

func TestBinarySearch(t *testing.T) {
	var list []int
	for i := 1; i <= 100; i++ {
		list = append(list, i)
	}
	for i := 1; i <= 100; i++ {
		got, err := BinarySearch(list, i, -1)
		assert.NoError(t, err)
		assert.Equal(t, i-1, got)
	}
	_, err := BinarySearch(list, 999, -1)
	assert.EqualError(t, err, "not found")
}

func TestBinarySearchWithLowAndHigh(t *testing.T) {
	var list []int
	for i := 1; i <= 100; i++ {
		list = append(list, i)
	}
	for i := 1; i <= 100; i++ {
		got, err := BinarySearchWitLowHigh(list, i, 0, len(list)-1)
		assert.NoError(t, err)
		assert.Equal(t, i-1, got)
	}
	_, err := BinarySearchWitLowHigh(list, 999, 0, len(list)-1)
	assert.EqualError(t, err, "not found")
}

func BenchmarkBinarySearchWithChopping(b *testing.B) {
	var list []int
	for i := 1; i <= 100; i++ {
		list = append(list, i)
	}
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearch(list, 6, -1)
	}
}

// go1.18beta1 test -bench=. -run=BenchmarkBinarySearch -benchmem
// goos: darwin
// goarch: amd64
// pkg: grokking-algorithms/chapter04
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkBinarySearchWithChopping-12            78129811                15.23 ns/op            0 B/op          0 allocs/op
// BenchmarkBinarySearchWithHighAndLow-12          166489551                7.167 ns/op           0 B/op          0 allocs/op
// PASS
// ok      grokking-algorithms/chapter04   3.240s
// WITH GENERICS:
// go1.18beta1 test -bench=. -run=BenchmarkBinarySearch -benchmem
// goos: darwin
// goarch: amd64
// pkg: grokking-algorithms/chapter04
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkBinarySearchWithChopping-12            89500888                12.98 ns/op            0 B/op          0 allocs/op
// BenchmarkBinarySearchWithHighAndLow-12          165573376                7.264 ns/op           0 B/op          0 allocs/op
// PASS
// ok      grokking-algorithms/chapter04   3.224s
// Wow... The chopping part actually got a bit faster.
func BenchmarkBinarySearchWithHighAndLow(b *testing.B) {
	var list []int
	for i := 1; i <= 100; i++ {
		list = append(list, i)
	}
	l := len(list) - 1
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearchWitLowHigh(list, 6, 0, l)
	}
}
