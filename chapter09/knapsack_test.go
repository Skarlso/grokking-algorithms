package chapter09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsack(t *testing.T) {
	items := []Item{
		{Name: "guitar", Value: 1500, Weight: 1},
		{Name: "stereo", Value: 3000, Weight: 4},
		{Name: "laptop", Value: 2000, Weight: 3},
	}
	max := Knapsack(items, 4)
	assert.Equal(t, 3500, max)
	items = append(items, Item{
		Name:   "iphone",
		Value:  2000,
		Weight: 1,
	})
	max = Knapsack(items, 4)
	assert.Equal(t, 4000, max)
}
