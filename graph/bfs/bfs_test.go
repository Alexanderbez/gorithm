package bfs_test

import (
	"testing"

	"github.com/alexanderbez/gorithm/graph/bfs"
	"github.com/stretchr/testify/require"
)

func TestBFS(t *testing.T) {
	q := bfs.NewGraph()

	nodes := make(map[int]*bfs.Node)
	for i := 1; i <= 6; i++ {
		nodes[i] = &bfs.Node{Value: i}
	}

	q.AddEdge(nodes[1], nodes[2])
	q.AddEdge(nodes[1], nodes[3])
	q.AddEdge(nodes[2], nodes[4])
	q.AddEdge(nodes[2], nodes[5])
	q.AddEdge(nodes[3], nodes[5])
	q.AddEdge(nodes[4], nodes[5])
	q.AddEdge(nodes[4], nodes[6])
	q.AddEdge(nodes[5], nodes[6])

	traversal := []int{}
	q.BFS(nodes[1], func(n *bfs.Node) (halt bool) {
		traversal = append(traversal, n.Value)
		return false
	})

	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, traversal)
}
