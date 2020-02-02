package dfs_test

import (
	"testing"

	"github.com/alexanderbez/gorithm/graph/dfs"
	"github.com/stretchr/testify/require"
)

func TestDFS(t *testing.T) {
	q := dfs.NewGraph()

	nodes := make(map[int]*dfs.Node)
	for i := 1; i <= 6; i++ {
		nodes[i] = &dfs.Node{Value: i}
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
	q.DFS(nodes[1], func(n *dfs.Node) (halt bool) {
		traversal = append(traversal, n.Value)
		return false
	})

	require.Equal(t, []int{1, 3, 5, 6, 2, 4}, traversal)
}
