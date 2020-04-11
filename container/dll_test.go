package container_test

import (
	"testing"

	"github.com/alexanderbez/gorithm/container"
	"github.com/stretchr/testify/require"
)

func TestDoubleLinkedList(t *testing.T) {
	dll := container.NewDoubleLinkedList()
	require.Equal(t, 0, dll.Len())

	require.Nil(t, dll.Front())
	require.Nil(t, dll.Back())

	size := 1000
	nodes := make([]*container.Node, size)

	for i := 0; i < size; i++ {
		var n *container.Node

		if i%2 == 0 {
			n = dll.InsertFront(i)
			require.NotNil(t, n)
			require.Equal(t, i, dll.Front().Value.(int))
			require.Equal(t, i, n.Value.(int))

			dll.MoveToBack(n)
			require.Equal(t, i, dll.Back().Value.(int))
		} else {
			n = dll.InsertBack(i)
			require.NotNil(t, n)
			require.Equal(t, i, dll.Back().Value.(int))
			require.Equal(t, i, n.Value.(int))

			dll.MoveToFront(n)
			require.Equal(t, i, dll.Front().Value.(int))
		}

		nodes[i] = n
	}

	require.Equal(t, size, dll.Len())

	for i, n := range nodes {
		v := dll.Remove(n)
		require.Equal(t, i, v.(int))
		require.Equal(t, size-(i+1), dll.Len())

		v = dll.Remove(n)
		require.Equal(t, i, v.(int))
		require.Equal(t, size-(i+1), dll.Len())
	}
}
