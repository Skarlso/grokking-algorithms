// Package chapter04 demonstrates Divide and Conquer algorithms using Quicksort.
// To use divide and conquer, you define these two things:
// 1. Figure out the base case. This should be the simplest possible case.
// 2. Divide or decrease your problem until it becomes the base case.
package chapter04

import "constraints"

// Sum an array.
// Base case is returning 0 if the number of elements is equal to zero.
// While if it's equal to 1, we return the value of the first element.
// The recursion is, that is the divide part, is to converge to the base case.
// In this case, it's reducing the number of elements per call. So we add the
// first element to the reduced list which will converge to 0 list items.
// Get a list ->
// - if the list is empty, return 0
// - otherwise, the list is the sum of first element plus the rest of the list
func Sum[T constraints.Ordered](list []T) T {
	var r T
	if len(list) == 0 {
		return r
	}
	return list[0] + Sum(list[1:])
}

// CountItems counts the number of items in a list. This is similar to sum,
// but now we keep a running total.
// Base case: There are zero items in the list.
// Recursion: Add 1 to the number of items in the list for the reduced list.
func CountItems[T constraints.Ordered, R constraints.Integer](list []T) R {
	if len(list) == 0 {
		return 0
	}
	return 1 + CountItems[T, R](list[1:])
}
