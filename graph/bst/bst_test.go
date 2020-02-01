package bst_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alexanderbez/gorithm/graph/bst"
	"github.com/stretchr/testify/require"
)

func TestTraverse(t *testing.T) {
	tree := bst.NewTree()
	values := []int{8, 3, 10, 1, 6, 4, 7, 14, 13}

	for _, value := range values {
		tree.Add(value)
	}

	require.Equal(t, 9, tree.Size())

	traversal := []int{}
	tree.Traverse(func(i int) {
		traversal = append(traversal, i)
	})

	expected := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	require.Equal(t, expected, traversal)
}

func TestSearch(t *testing.T) {
	tree := bst.NewTree()
	size := 10000
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	vals := make(map[int]struct{})

	for i := 0; i < size; i++ {
		v := rand.Int()
		_, ok := vals[v]
		for ok {
			v = rand.Int()
			_, ok = vals[v]
		}

		vals[v] = struct{}{}
		tree.Add(v)
	}

	require.Equal(t, size, tree.Size())

	for v := range vals {
		require.NotNil(t, tree.Search(v))
	}

	for i := 0; i < size; i++ {
		v := rand.Int()
		n := tree.Search(v)

		_, ok := vals[v]
		if ok {
			require.NotNil(t, n)
		} else {
			require.Nil(t, n)
		}
	}
}
