package helpers_test

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vikesh-raj/go-algo-practice/helpers"
)

func TestArrayToString(t *testing.T) {
	require.Equal(t, "[-2 4]", helpers.ArrayToString([]int{-2, 4}))
	require.Equal(t, "[]", helpers.ArrayToString([]int{}))
	require.Equal(t, "[20]", helpers.ArrayToString([]int{20}))
	require.Equal(t, "nil", helpers.ArrayToString(nil))
}

func TestListToString(t *testing.T) {
	list := list.New()
	require.Equal(t, "[]", helpers.ListToString(list))
	list.PushBack(10)
	require.Equal(t, "[10]", helpers.ListToString(list))
	list.PushBack(20)
	list.PushBack(30)
	require.Equal(t, "[10 20 30]", helpers.ListToString(list))
}
