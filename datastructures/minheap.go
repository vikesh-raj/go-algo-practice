package datastructures

import "github.com/vikesh-raj/go-algo-practice/helpers"

type MinHeap struct {
	data []int
}

func NewMinHeap(array []int) *MinHeap {
	heap := &MinHeap{data: array}
	heap.BuildHeap(array)
	return heap
}

func parentIndex(idx int) int {
	return (idx - 1) / 2
}

func childrenIndices(idx int) (int, int) {
	return 2*idx + 1, 2*idx + 2
}

func (h *MinHeap) BuildHeap(array []int) {
	for i := parentIndex(len(array)); i >= 0; i-- {
		h.siftDown(i, len(array))
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	for currentIndex < endIndex {
		left, right := childrenIndices(currentIndex)
		if left >= endIndex {
			return
		}
		minIndex := left
		min := h.data[left]
		if right < endIndex && h.data[right] < min {
			minIndex = right
		}

		if h.data[currentIndex] < h.data[minIndex] {
			break
		}
		h.data[currentIndex], h.data[minIndex] = h.data[minIndex], h.data[currentIndex]
		currentIndex = minIndex
	}
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeap) siftUp() {
	idx := len(h.data) - 1
	for idx != 0 {
		parent := parentIndex(idx)
		if h.data[idx] >= h.data[parent] {
			break
		}
		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		idx = parent
	}
}

func (h MinHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *MinHeap) Remove() int {
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	min := h.data[0]
	h.data[0] = last
	h.siftDown(0, len(h.data))
	return min
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.siftUp()
}

func (h *MinHeap) String() string {
	return helpers.ArrayToString(h.data)
}
