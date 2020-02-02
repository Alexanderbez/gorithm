package container_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alexanderbez/gorithm/container"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := container.NewQueue()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	s := 1000
	e := make([]int, s)

	require.Nil(t, q.Pop())

	for i := 0; i < s; i++ {
		v := r.Int()
		e[i] = v
		q.Push(v)
	}

	for i := 0; i < s; i++ {
		v := e[i]
		require.Equal(t, v, q.Pop().(int))
	}

	require.Nil(t, q.Pop())
}
