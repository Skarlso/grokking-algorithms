package chapter09

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Signed
}

// GenericItem defines a take-able item which has three properties. A weight, a value and a name for displaying.
type GenericItem[T Number] interface {
	Weight() T
	Name() string
	Value() T
}

type KnapItem struct {
	_Name   string
	_Value  int
	_Weight int
}

func (k KnapItem) Weight() int {
	return k._Weight
}

func (k KnapItem) Name() string {
	return k._Name
}

func (k KnapItem) Value() int {
	return k._Value
}

type Item struct {
	Name   string
	Value  int
	Weight int
}

// Knapsack is a popular problem about how to steal the most expensive items given a weight limit.
// I was really expecting the book to provider a proper code sample for this as I was looking forward
// to reading and understanding the nuances that a dynamic programming solution require. But, there were none.
// So I put this together while reading the chapter. One thing that wasn't intuitive is that I should
// start with an empty row full of zeros to make choosing the first row a lot easier.
func Knapsack(items []Item, limit int) int {
	// why +1? Because we add a row with 0-s that can akt as a minimum value for
	// the first items.
	cell := make([][]int, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		cell[i] = make([]int, limit+1)
	}
	// the first row is left empty with 0s, so it can be the `max` in case of the first item.
	// Take care to start from 1 for row and col as well.
	for i := 1; i <= len(items); i++ {
		// We are comparing it to the weight, so we begin from 1. But this could go backwards
		// as well, or just -1 it. In any case, we would have to come up with a nicer way
		// to determine the step of this loop any ways.
		for j := 1; j <= limit; j++ {
			item := items[i-1]
			if item.Weight > j {
				cell[i][j] = cell[i-1][j]
			} else {
				cell[i][j] = max(cell[i-1][j], cell[i-1][j-item.Weight]+item.Value)
			}
		}
	}
	return cell[len(items)][limit]
}

// KnapsackGeneric since you can't use limit as above because j-item.Weight needs to be an integer,
// we have to already have the row markers from somewhere.
// Once we have them, we can then run the same thing but instead of using it as an index,
// we always do an `IndexOf` on the item that we get.
func KnapsackGeneric[I GenericItem[int], T Number](items []I, limits []T) T {
	// why +1? Because we add a row with 0-s that can akt as a minimum value for
	// the first items.
	cell := make([][]T, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		cell[i] = make([]T, len(limits)+1)
	}
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= len(limits); j++ {
			item := items[i-1]
			limit := limits[j-1]
			if T(item.Weight()) > limit {
				cell[i][j] = cell[i-1][j]
			} else {
				next := IndexOf(limit-T(item.Weight()-1), limits)
				cell[i][j] = max(cell[i-1][j], cell[i-1][next]+T(item.Value()))
			}
		}
	}
	return cell[len(items)][len(limits)]
}

func IndexOf[T Number](num T, nums []T) int {
	for i, n := range nums {
		if n == num {
			return i
		}
	}
	return -1
}

// LongestCommonSubstring finds the longest substring that is common in each of the two strings.
func LongestCommonSubstring(a, b string) string {
	// We start from +1 again, so we have the first row empty.
	cell := make([][]int, len(a)+1)
	for i := range cell {
		cell[i] = make([]int, len(b)+1)
	}

	max := 0
	maxIndex := 0
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// -1 because we started from 1
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
				// save the last maximum point, or last sub index
				// because the greatest sub index might not be the last row:col.
				if cell[i][j] > max {
					maxIndex = i
					max = cell[i][j]
				}
			} else {
				cell[i][j] = 0
			}
		}
	}

	// From the last index - the maximum value to the last index.
	// Again, this was not intuitive to figure out.
	return a[maxIndex-max : maxIndex]
}

// LongestCommonSubsequence finds the longest subsequence length between two strings. They don't have to be
// continuous.
func LongestCommonSubsequence(a, b string) int {
	// We start from +1 again, so we have the first row empty.
	cell := make([][]int, len(a)+1)
	for i := range cell {
		cell[i] = make([]int, len(b)+1)
	}

	// almost the same as Knapsack
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// -1 because we started from 1
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
			} else {
				cell[i][j] = max(cell[i-1][j], cell[i][j-1])
			}
		}
	}

	return cell[len(a)][len(b)]
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}
