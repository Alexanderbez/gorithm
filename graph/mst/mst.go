package mst

import (
	"math"
)

// Graph reflects an undirected graph represented as an adjacency matrix. Each
// value can reflect the edge weight between two vertices or a non-zero value
// to reflect a connection.
type Graph struct {
	graph [][]int
}

// NewGraph builds a undirected graph represented as an n x n adjacency matrix.
func NewGraph(n int) Graph {
	edges := make([][]int, n)
	for i := range edges {
		edges[i] = make([]int, n)
	}

	return Graph{graph: edges}
}

// InsertUnweighted inserts an edge into the graph g without a specified weight
// w. An edge exists between two vertices if the value is non-zero.
func (g Graph) InsertUnweighted(v, u int) { g.graph[v][u] = 1; g.graph[u][v] = 1 }

// InsertWeighted inserts an edge into the graph g with a specified weight w.
func (g Graph) InsertWeighted(v, u, w int) { g.graph[v][u] = w; g.graph[u][v] = w }

// ExecPrim implements Prim’s Minimum Spanning Tree (MST) algorithm. As input,
// a starting vertex is given. The graph is represented via an adjacency matrix
// where each entry reflects the edge weight between two vertices. The MST total
// edge weight is returned.
//
// Note: Each edge weight is assumed to be non-negative.
//
// Algorithm:
// 1) Create a set, mstSet, that keeps track of vertices already included in
// MST.
// 2) Assign a key value to all vertices in the input graph. Initialize all key
// values as some infinite/dummy value. Assign key value as 0 for a given
// initial vertex so that it is picked first.
// 3) While mstSet doesn’t include all vertices
// ...a) Pick a vertex u which is not there in mstSet and has minimum key value.
// ...b) Include u to mstSet.
// ...c) Update key value of all adjacent vertices of u. To update the key
// values, iterate through all adjacent vertices. For every adjacent vertex v,
// if weight of edge u-v is less than the previous key value of v, update the
// key value as weight of u-v.
func (g Graph) ExecPrim(s int) int {
	mstValue := 0
	mstSet := make(map[int]bool)
	vertexValues := make([]int, len(g.graph))

	// Set vertex key values
	for i := range vertexValues {
		vertexValues[i] = math.MaxInt64
	}
	vertexValues[s] = 0

	// While all the vertices have not been added to the MST set
	for len(mstSet) != len(g.graph) {
		// Pick a vertex (u) which is not there in mstSet and has minimum key
		// value.
		min := math.MaxInt64
		var minVertex int

		for vertex := range g.graph {
			if !mstSet[vertex] && vertexValues[vertex] <= min {
				min = vertexValues[vertex]
				minVertex = vertex
			}
		}

		// Include minVertex (u) to the MST set
		mstSet[minVertex] = true

		// Update key value of all adjacent vertices not in the MST set
		for neighbor, weight := range g.graph[minVertex] {
			if weight != 0 && !mstSet[neighbor] && weight < vertexValues[neighbor] {
				vertexValues[neighbor] = weight
			}
		}
	}

	for i := 0; i < len(vertexValues); i++ {
		mstValue += vertexValues[i]
	}

	return mstValue
}
