package datastructures

import "github.com/vikesh-raj/go-algo-practice/helpers"

type MaxHeap struct {
	data []int
}

func NewMaxHeap(array []int) *MaxHeap {
	heap := &MaxHeap{data: array}
	heap.BuildHeap(array)
	return heap
}

func (h *MaxHeap) BuildHeap(array []int) {
	for i := parentIndex(len(array)); i >= 0; i-- {
		h.siftDown(i, len(array))
	}
}

func (h *MaxHeap) siftDown(currentIndex, endIndex int) {
	for currentIndex < endIndex {
		left, right := childrenIndices(currentIndex)
		if left >= endIndex {
			return
		}
		maxIndex := left
		max := h.data[left]
		if right < endIndex && h.data[right] > max {
			maxIndex = right
		}

		if h.data[currentIndex] > h.data[maxIndex] {
			break
		}
		h.data[currentIndex], h.data[maxIndex] = h.data[maxIndex], h.data[currentIndex]
		currentIndex = maxIndex
	}
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MaxHeap) siftUp() {
	idx := len(h.data) - 1
	for idx != 0 {
		parent := parentIndex(idx)
		if h.data[idx] <= h.data[parent] {
			break
		}
		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		idx = parent
	}
}

func (h MaxHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *MaxHeap) Remove() int {
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	max := h.data[0]
	h.data[0] = last
	h.siftDown(0, len(h.data))
	return max
}

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.siftUp()
}

func (h *MaxHeap) String() string {
	return helpers.ArrayToString(h.data)
}
