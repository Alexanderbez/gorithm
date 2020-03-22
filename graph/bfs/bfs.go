package bfs

import (
	"github.com/alexanderbez/gorithm/container"
)

type (
	Node struct {
		ID     string
		Weight int
	}

	// Graph defines a directional graph implementation.
	Graph struct {
		nodes map[string][]*Node
	}
)

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string][]*Node),
	}
}

// AddEdge adds a unidirectional edge between a source node and a destination node.
func (g *Graph) AddEdge(src, dst *Node) {
	_, ok := g.nodes[src.ID]
	if !ok {
		g.nodes[src.ID] = make([]*Node, 0)
	}

	g.nodes[src.ID] = append(g.nodes[src.ID], dst)
}

// BFS performs a breadth first search over the graph starting at a given source
// Node. For each visited node, the provided match callback is invoked. If the
// callback returns true, it indicates a match and traversal halts.
func (g *Graph) BFS(src *Node, match func(n *Node) (halt bool)) {
	visited := make(map[string]bool)

	queue := container.NewQueue()
	queue.Push(src)
	visited[src.ID] = true

	for queue.Size() != 0 {
		node := queue.Pop().(*Node)

		if match(node) {
			return
		}

		// add all neighbors to queue
		for _, neighbor := range g.nodes[node.ID] {
			if !visited[neighbor.ID] {
				visited[neighbor.ID] = true
				queue.Push(neighbor)
			}
		}
	}
}
