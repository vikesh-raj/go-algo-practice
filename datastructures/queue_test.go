package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vikesh-raj/go-algo-practice/datastructures"
)

func TestQueue(t *testing.T) {
	queue := datastructures.NewQueue()
	require.NotNil(t, queue)

	require.True(t, queue.IsEmpty())
	queue.Enqueue(10)
	queue.Enqueue(12)

	require.Equal(t, "[10 12]", queue.String())
	require.False(t, queue.IsEmpty())

	require.Equal(t, 2, queue.Len())
	front, ok := queue.Peek()
	require.True(t, ok)
	require.Equal(t, 10, front)

	front, ok = queue.Dequeue()
	require.True(t, ok)
	require.Equal(t, 10, front)
	require.Equal(t, 1, queue.Len())

	front, ok = queue.Peek()
	require.True(t, ok)
	require.Equal(t, 12, front)

	front, ok = queue.Dequeue()
	require.True(t, ok)
	require.Equal(t, 12, front)
	require.Equal(t, 0, queue.Len())

	_, ok = queue.Dequeue()
	require.False(t, ok)

	_, ok = queue.Dequeue()
	require.False(t, ok)
	require.Equal(t, "[]", queue.String())

}
