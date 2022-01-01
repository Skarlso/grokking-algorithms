package chapter07

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	// Nodes with weights
	g := graph{
		"start": map[string]int{
			"a": 6,
			"b": 2,
		},
		"a": map[string]int{
			"fin": 1,
		},
		"b": map[string]int{
			"a":   3,
			"fin": 5,
		},
		"fin": map[string]int{},
	}
	shortestPath := Dijkstra(g, "start", "fin")
	// now, travers the node back from `fin` to `start`.
	start := "fin"
	var path []string
	path = append(path, start)
	for start != "start" {
		start = shortestPath[start]
		path = append(path, start)
	}
	// reverse, since we traversed backwards
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	assert.Equal(t, "start,b,a,fin", strings.Join(path, ","))
}

func TestDijkstraWithPQ(t *testing.T) {
	// Nodes with weights
	g := graph{
		"start": map[string]int{
			"a": 6,
			"b": 2,
		},
		"a": map[string]int{
			"fin": 1,
		},
		"b": map[string]int{
			"a":   3,
			"fin": 5,
		},
		"fin": map[string]int{},
	}
	shortestPath := DijkstraWithPQ(g, "start", "fin")
	// assemble the path
	current := "fin"
	var path []string
	for current != "start" {
		path = append(path, current)
		current = shortestPath[current]
	}
	path = append(path, current)
	// reverse, since we traversed backwards
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	assert.Equal(t, "start,b,a,fin", strings.Join(path, ","))
}
