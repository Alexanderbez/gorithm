package container_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alexanderbez/gorithm/container"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	s := container.NewStack()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	k := 1000
	e := make([]int, k)

	require.Nil(t, s.Pop())

	for i := 0; i < k; i++ {
		v := r.Int()
		e[i] = v
		s.Push(v)
	}

	for i := k - 1; i >= 0; i-- {
		v := e[i]
		require.Equal(t, v, s.Pop().(int))
	}

	require.Nil(t, s.Pop())
}
