package chapter06

// BFS I find that BFS is better explained at
// [Red Blob Games](https://www.redblobgames.com/pathfinding/a-star/introduction.html). With
// lovely animations. But this one wasn't too bad either.
// This implementation will use the simple map based search from the book.
// Also note, that this implementation does not work if you happened to start from the seller.
// Since the starting point is actually not in the queue. It starts with visiting the vertices
// immediately. Also, this doesn't return the path in any format, it just says it found the
// thing you are looking for. There is no proof that the path is the shortest.
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
	Nodes []Node
}

// New returns a new graph given a list of nodes.
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	// need to implement something here.
	return &Graph[Node, Edge]{
		Nodes: nodes,
	}
}

// GraphNode is a node in a graph which fulfills the type constraint of nodes.
type GraphNode struct {
	Value      string
	GraphEdges []*GraphEdge
}

// GraphEdge is an edge in a graph which fulfills the type constraint of edges.
type GraphEdge struct {
	From *GraphNode
	To   *GraphNode
}

func (g *GraphNode) Edges() []*GraphEdge {
	return g.GraphEdges
}

func (g *GraphEdge) Nodes() (*GraphNode, *GraphNode) {
	return g.From, g.To
}

// BFS returns the shortest path between two nodes,
// as a list of edges. Eq is used to determine parity between nodes. The Eq functions makes
// this function generic enough instead of trying to shoehorn some other `comparable` type
// into `Node`.
// Note, this would be much more user-friendly, if it would return a map of `cameFrom`s which
// then could be traversed back to find the shortest path, rather than returning a slice of
// edges which is difficult to follow back.
func (g *Graph[Node, Edge]) BFS(from, to Node, eq func(self, other Node) bool) []Edge {
	queue := []Node{from}
	var path []Edge
	// this should be a map for O(1) recall, but in that case I would have to also get a function
	// which returns a unique identifier for the nodes. But since I already have an Eq function
	// I can use that.
	var seen []Node
	var current Node
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		edges := current.Edges()
		// For each edge, gather the nodes. The edges contain from -> to syntax,
		// so we ignore the `from` one. We are only interested in the `to`.
		for _, edge := range edges {
			path = append(path, edge)
			_, dest := edge.Nodes()
			if eq(dest, to) {
				return path
			}
			visited := false
			for _, s := range seen {
				if eq(s, dest) {
					visited = true
					break
				}
			}
			if !visited {
				seen = append(seen, dest)
				queue = append(queue, dest)
			}
		}
	}
	return nil
}
