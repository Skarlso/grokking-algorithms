// Package chapter04 demonstrates Divide and Conquer algorithms using Quicksort.
// To use divide and conquer, you define these two things:
// 1. Figure out the base case. This should be the simplest possible case.
// 2. Divide or decrease your problem until it becomes the base case.
package chapter04

import (
	"constraints"
	"errors"
)

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

// FindMax what is our base case here.
// Base case: Return the maximum value
// Recur: Compare the first item in the list with max and if
// it is bigger than max, then set max and call the function with
// the new max value and the rest of the list.
// Another way of doing this, if we don't want to keep track of max,
// is to compare in case the len is 2, otherwise call max with the
// rest of the list and if that is greater, return that otherwise
// the item in the list. I find this to be more readable and shorter.
func FindMax[T constraints.Ordered](list []T, max T) T {
	if len(list) == 0 {
		return max
	}
	if list[0] > max {
		max = list[0]
	}
	return FindMax(list[1:], max)
}

// BinarySearch defines the binary search algorithm using divide and conquer.
// I choose to chop up the list instead of keeping track of high and low.
// Base case is if the value of list equals with T or not. ( in case the list is
// one item long ).
// Recur: Reduce the list... aka, call with tail in case it's higher and call with
// head in case it's lower.
// The book defines D&C with breaking down the array. But in this case, that's a sub-par
// approach. See below for a better approach.
func BinarySearch[T constraints.Ordered, I constraints.Signed](list []T, item T, index I) (I, error) {
	var i I
	if len(list) == 0 {
		return i, errors.New("not found")
	}
	if len(list) == 1 {
		if list[0] == item {
			return index, nil
		}
		return i, errors.New("not found")
	}
	middle := len(list) / 2
	if index == -1 { // Signed integer
		index = I(middle)
	}
	if list[middle] > item {
		newList := list[:middle]
		if len(newList)%2 == 0 {
			l := len(newList) / 2
			index = index - I(l) // need to convert because this is a generic type
		} else {
			index -= I((len(newList) / 2) + 1)
		}
		return BinarySearch[T](newList, item, index)
	} else if list[middle] < item {
		newList := list[middle+1:]
		index += I((len(newList) / 2) + 1)
		return BinarySearch[T](newList, item, index)
	}

	return index, nil
}

// BinarySearchWitLowHigh as expected, this version is much faster, also reads a bit better.
// And there is no need to mess around with the index, trying to calculate it properly.
func BinarySearchWitLowHigh[T constraints.Ordered](list []T, item T, low int, high int) (int, error) {
	if low > high {
		return -1, errors.New("not found")
	}
	middle := (low + high) / 2
	if item == list[middle] {
		return middle, nil
	}
	if item < list[middle] {
		return BinarySearchWitLowHigh(list, item, low, middle-1)
	}
	return BinarySearchWitLowHigh(list, item, middle+1, high)
}
