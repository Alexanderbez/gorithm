package bfs

import (
	"github.com/alexanderbez/gorithm/container"
)

type (
	Node struct {
		Value int
		// neighbors []*Node
	}

	Graph struct {
		nodes map[int][]*Node
	}
)

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int][]*Node),
	}
}

// AddEdge adds a unidirectional edge between a source node and a destination node.
func (g *Graph) AddEdge(src, dst *Node) {
	_, ok := g.nodes[src.Value]
	if !ok {
		g.nodes[src.Value] = make([]*Node, 0)
	}

	g.nodes[src.Value] = append(g.nodes[src.Value], dst)
}

// BFS performs a bread first search over the graph starting at a given source
// Node. For each visited node, the provided match callback is invoked. If the
// callback returns true, it indicates a match and traversal halts.
func (g *Graph) BFS(src *Node, match func(n *Node) (halt bool)) {
	visited := make(map[int]bool)

	queue := container.NewQueue()
	queue.Push(src)

	for queue.Size() != 0 {
		node := queue.Pop().(*Node)

		if match(node) {
			return
		}

		// add all neighbors to queue
		for _, neighbor := range g.nodes[node.Value] {
			if !visited[neighbor.Value] {
				visited[neighbor.Value] = true
				queue.Push(neighbor)
			}
		}
	}
}
