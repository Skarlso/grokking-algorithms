package chapter07

import (
	"constraints"
	"container/heap"
	"math"
)

// Number which can be used to track `cost`.
type Number interface {
	constraints.Integer | constraints.Float
}

type graph[T comparable, I Number] map[T]map[T]I

// costs contains the cost of all nodes.
type costs[T comparable, I Number] map[T]I

type tracker[T comparable] map[T]struct{}

// First, I'll implement what the book suggests to use for Dijkstra.
// Second, I will show how to implement Dijkstra using a Priority Queue.
// There is no Priority queue in Go at hand, but it's easy to implement
// one using `container` package's Heap.

// Dijkstra
// The book sets up `parents` manually. We aren't going to do that.
// We are going to return a `cameFrom` map which will track where the
// node came from. That can be then traversed backwards to get the path
// the algorithm took. Further, then it doesn't explain what to do
// with the `parents` dict at all or the costs. The test for this function shows how
// to use the output.
// Note that the book uses costs map. Usually, Dijkstra just uses a
// Priority Queue to solve the updating of the costs which I'll demonstrate
// with a separate function.
func Dijkstra[T comparable](g graph[T, int], start, finish T) map[T]T {
	// initialise costs and cameFrom
	cameFrom := make(map[T]T)
	c := make(costs[T, int])
	// We set up initial costs for all the nodes that come from the `start` node.
	for node, cost := range g[start] {
		c[node] = cost
		cameFrom[node] = start
	}
	// mark the finish as an unknown cost
	c[finish] = math.MaxInt

	// set up the tracker
	seen := make(tracker[T])
	// find the lowest cost node from the current nodes.
	node := findLowestCostNode(c, seen)
	var end T
	// while there are nodes...
	for node != end {
		for k, v := range g[node] {
			newCost := c[node] + v
			if c[k] > newCost {
				c[k] = newCost
				cameFrom[k] = node
			}
		}
		seen[node] = struct{}{}
		node = findLowestCostNode(c, seen)
	}
	return cameFrom
}

func findLowestCostNode[T comparable](c costs[T, int], seen tracker[T]) T {
	min := math.MaxInt
	var lowest T
	for node, cost := range c {
		if _, ok := seen[node]; cost < min && !ok {
			min = cost
			lowest = node
		}
	}
	return lowest
}

// DijkstraWithPQ now with a Priority Queue. Returns the `cameFrom` map. Still tracks the cost,
// but the priority queue will take care of finding the lowest node we don't have to construct
// the cost map with the first couple of nodes, and we don't have to care about assigning
// math.MaxInt to the finish.
// I didn't really see the point in providing generics here, unless, I really go wild
// and provide a proper generic graph like the one in BFS, and implement this there.
func DijkstraWithPQ(g graph[string, int], start, finish string) map[string]string {
	var pq PriorityQueue
	// it's important to use `heap.*` and not `pq.*` directly, because
	// `heap` is the one which will reorder based on priority using `Len, Less, Swap`.
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    start,
		priority: 1,
	})
	cost := map[string]int{
		start: 1,
	}
	cameFrom := map[string]string{
		start: start,
	}
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		if current.value == finish {
			break
		}
		for next, c := range g[current.value] {
			// add the weight of the edge to the current movement cost.
			// that will be the cost of the new, overall move from start for this node.
			newCost := cost[current.value] + c
			if v, ok := cost[next]; !ok || newCost < v {
				cameFrom[next] = current.value
				cost[next] = newCost
				heap.Push(&pq, &Item{
					value:    next,
					priority: newCost,
				})
			}
		}
	}
	return cameFrom
}
