package chapter06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	// finds the seller
	found := BFS(graph, "you")
	assert.True(t, found)
	// it will not find the seller if starting from the seller.
	// This is actually bad.
	found = BFS(graph, "thom")
	assert.False(t, found)
}
