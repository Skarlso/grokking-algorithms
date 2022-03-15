package chapter03

import "constraints"

// Factorial demonstrates how the call stack works.
func Factorial[T constraints.Integer](i T) T {
	if i == 1 {
		return i
	}
	return i * Factorial(i-1)
}
