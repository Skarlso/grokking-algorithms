package chapter03

import (
	"constraints"
)

// PlainRecurse just prints something to the console.
func PlainRecurse[T constraints.Integer](i T, list []T) []T {
	list = append(list, i)
	// base case, aka, the case when the function doesn't call itself again
	if i == 0 {
		return list
	}
	return PlainRecurse(i-1, list) // tail recursion case
}
