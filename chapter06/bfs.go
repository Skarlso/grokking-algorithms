package chapter06

// BFS I find that BFS is better explained at
// [Red Blob Games](https://www.redblobgames.com/pathfinding/a-star/introduction.html). With
// lovely animations. But this one wasn't too bad either.
// This implementation will use the simple map based search from the book.
// Also note, that this implementation does not work if you happened to start from the seller.
// Since the starting point is actually not in the queue. It starts with visiting the vertices
// immediately.
func BFS(graph map[string][]string, name string) bool {
	queue := graph[name]
	seen := map[string]struct{}{}
	var current string
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if _, ok := seen[current]; !ok {
			if isSeller(current) {
				// found the seller
				return true
			}
			seen[current] = struct{}{}
			queue = append(queue, graph[current]...)
		}
	}
	return false
}

// dummy function which denotes when to stop the search. A seller simply
// has a name which ends with an `m`.
func isSeller(current string) bool {
	return current[len(current)-1] == 'm'
}

// GenericShortestPath will use the graph implementation from the type proposal here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#mutually-referencing-type-parameters

// NodeConstraint is the type constraint for graph nodes:
// they must have an `Edges` method that returns the Edges
// that connect to this Node.
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

// EdgeConstraint is the type constraint for graph edges:
// they must have a `Nodes` method that returns the two Nodes
// that this edge connects.
type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

// Graph is a graph composed of nodes and edges.
type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct {
	// need to implement something here.
}

// New returns a new graph given a list of nodes.
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	// need to implement something here.
	return nil
}

// GraphNode is a node in a graph which fulfills the type constraint of nodes.
type GraphNode struct {
}

// GraphEdge is an edge in a graph which fulfills the type constraint of edges.
type GraphEdge struct {
}

func (g *GraphNode) Edges() []*GraphEdge {
	return nil
}

func (g *GraphEdge) Nodes() (*GraphNode, *GraphNode) {
	return nil, nil
}

// BFS returns the shortest path between two nodes,
// as a list of edges.
func (g *Graph[Node, Edge]) BFS(from, to Node) []Edge {
	return nil
}
