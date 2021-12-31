package chapter06

import (
	"fmt"
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

func TestGraph_BFS(t *testing.T) {
	A := &GraphNode{
		Value: "A",
	}
	B := &GraphNode{
		Value: "B",
	}
	start := &GraphNode{
		Value: "start",
	}
	start.GraphEdges = []*GraphEdge{
		{
			From: start,
			To:   A,
		},
		{
			From: start,
			To:   B,
		},
	}
	C := &GraphNode{
		Value: "C",
	}
	D := &GraphNode{
		Value: "D",
	}
	A.GraphEdges = []*GraphEdge{
		{
			From: A,
			To:   C,
		},
		{
			From: A,
			To:   D,
		},
	}
	E := &GraphNode{
		Value: "E",
	}
	F := &GraphNode{
		Value: "F",
	}
	B.GraphEdges = []*GraphEdge{
		{
			From: B,
			To:   E,
		},
		{
			From: B,
			To:   F,
		},
	}
	graph := New[*GraphNode, *GraphEdge]([]*GraphNode{start, A, B, C, D, E, F})
	got := graph.BFS(start, F, func(self, other *GraphNode) bool {
		return self.Value == other.Value
	})
	for _, g := range got {
		from, to := g.Nodes()
		fmt.Println(from, to)
	}
	assert.Equal(t, nil, got)
}
