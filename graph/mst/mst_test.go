package mst_test

import (
	"testing"

	"github.com/alexanderbez/gorithm/graph/mst"
	"github.com/stretchr/testify/require"
)

func TestExecPrim(t *testing.T) {
	// graph reflected as shown: http://www.geeksforgeeks.org/wp-content/uploads/Fig-11.jpg
	g := mst.NewGraph(9)
	g.InsertWeighted(0, 1, 4)
	g.InsertWeighted(0, 7, 8)
	g.InsertWeighted(1, 2, 8)
	g.InsertWeighted(1, 7, 11)
	g.InsertWeighted(7, 8, 7)
	g.InsertWeighted(7, 6, 1)
	g.InsertWeighted(2, 8, 2)
	g.InsertWeighted(2, 3, 7)
	g.InsertWeighted(2, 5, 4)
	g.InsertWeighted(8, 6, 6)
	g.InsertWeighted(3, 4, 9)
	g.InsertWeighted(3, 5, 14)
	g.InsertWeighted(6, 5, 2)

	require.Equal(t, 37, g.ExecPrim(0))
}
