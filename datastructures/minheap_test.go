package datastructures_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikesh-raj/go-algo-practice/datastructures"
)

func TestMinHeap(t *testing.T) {
	array := []int{5, 9, 3, 0, 2, 1}
	heap := datastructures.NewMinHeap(array)
	fmt.Println(heap)

	top, ok := heap.Peek()
	assert.True(t, ok)
	assert.Equal(t, 0, top)

	top = heap.Remove()
	assert.Equal(t, 0, top)
	assert.Equal(t, 5, heap.Len())
	top, ok = heap.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, top)
}
