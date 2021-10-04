package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vikesh-raj/go-algo-practice/datastructures"
)

func TestStack(t *testing.T) {
	stack := datastructures.NewStack()
	require.NotNil(t, stack)

	require.True(t, stack.IsEmpty())
	stack.Push(10)
	stack.Push(12)

	require.Equal(t, "[10 12]", stack.String())
	require.False(t, stack.IsEmpty())

	require.Equal(t, 2, stack.Len())
	top, ok := stack.Top()
	require.True(t, ok)
	require.Equal(t, 12, top)

	top, ok = stack.Pop()
	require.True(t, ok)
	require.Equal(t, 12, top)
	require.Equal(t, 1, stack.Len())

	top, ok = stack.Top()
	require.True(t, ok)
	require.Equal(t, 10, top)

	top, ok = stack.Pop()
	require.True(t, ok)
	require.Equal(t, 10, top)
	require.Equal(t, 0, stack.Len())

	_, ok = stack.Pop()
	require.False(t, ok)

	_, ok = stack.Top()
	require.False(t, ok)
	require.Equal(t, "[]", stack.String())
}
