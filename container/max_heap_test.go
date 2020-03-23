package container_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/alexanderbez/gorithm/container"
	"github.com/stretchr/testify/require"
)

func TestMaxHeap(t *testing.T) {
	mh := container.NewMaxHeap()
	rng := rand.New(rand.NewSource(time.Now().Unix()))

	vals := make([]int, 1024)
	for i := range vals {
		val := rng.Int()
		vals[i] = val
		mh.Insert(val)
	}

	sort.Ints(vals)
	for i := len(vals) - 1; i >= 0; i-- {
		max, ok := mh.Peek()
		require.Equal(t, vals[i], max)
		require.True(t, ok)

		max2, ok := mh.Pop()
		require.True(t, ok)
		require.Equal(t, vals[i], max2)

		require.Equal(t, max, max2)
	}

	max, ok := mh.Peek()
	require.Equal(t, 0, max)
	require.False(t, ok)

	max, ok = mh.Pop()
	require.Equal(t, 0, max)
	require.False(t, ok)
}
