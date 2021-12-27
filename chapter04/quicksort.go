package chapter04

import "constraints"

// Quicksort uses inductive proof to verify that it can sort a list
// with zero or one item; therefor, if the list is reduced to that number
// we know we can sort it. How that's done is by choosing a pivot ( which
// can be any one of the list's items ) then dividing the list into lesser than,
// and greater than numbers based on the pivot. We then call quicksort on both of those
// and merge them together when returning. I said the pivot can be any one of the
// list's items. However, choosing the right pivot can speed quicksort to it's
// best case which is O(nLogn). Choosing a bad pivot, results in the worst case
// which is O(n^2).
func Quicksort[T constraints.Ordered](list []T) []T {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]
	less := make([]T, 0)
	more := make([]T, 0)
	for _, v := range list[1:] {
		if v <= pivot {
			less = append(less, v)
		} else if v > pivot {
			more = append(more, v)
		}
	}
	less = append(Quicksort(less), pivot)
	return append(less, Quicksort(more)...)
}
