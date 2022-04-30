package chapter02

import "golang.org/x/exp/constraints"

// SelectionSort constructs a new slice with ordered elements by continuously finding
// the lowest element. It alters the original slice by popping that element out of the
// slice.
func SelectionSort[T constraints.Ordered](list []T) []T {
	var newList []T
	for i := 0; i < len(list); i++ {
		smallest := FindSmallestElement(list)
		newList = append(newList, list[smallest])
		list = append(list[:smallest], list[smallest+1:]...)
		i--
	}
	return newList
}

// FindSmallestElement finds the smallest element in a list and returns its index.
func FindSmallestElement[T constraints.Ordered](list []T) int {
	first := list[0]
	smallestIndex := 0
	for i, v := range list {
		if v < first {
			first = v
			smallestIndex = i
		}
	}
	return smallestIndex
}
