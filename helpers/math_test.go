package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vikesh-raj/go-algo-practice/helpers"
)

func TestAbs(t *testing.T) {
	require.Equal(t, 3, helpers.Abs(3))
	require.Equal(t, 3, helpers.Abs(-3))
	require.Equal(t, 0, helpers.Abs(0))
}

func TestMax(t *testing.T) {
	require.Equal(t, 3, helpers.Max(1, 3))
	require.Equal(t, 1, helpers.Max(1, -3))
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, helpers.Min(1, 3))
	require.Equal(t, -3, helpers.Min(1, -3))
}

func TestMaxArray(t *testing.T) {
	table := []struct {
		array []int
		max   int
	}{
		{array: []int{5, 19, -1, -4, 10}, max: 19},
		{array: []int{-1}, max: -1},
	}

	for _, item := range table {
		assert.Equal(t, item.max, helpers.MaxArray(item.array))
		assert.Equal(t, item.max, helpers.MaxMulti(item.array...))
	}
}

func TestMinArray(t *testing.T) {
	table := []struct {
		array []int
		min   int
	}{
		{array: []int{5, 19, -1, -4, 10}, min: -4},
		{array: []int{-5}, min: -5},
	}

	for _, item := range table {
		assert.Equal(t, item.min, helpers.MinArray(item.array))
		assert.Equal(t, item.min, helpers.MinMulti(item.array...))
	}
}

func TestMaxWithIndex(t *testing.T) {
	table := []struct {
		array    []int
		max      int
		maxIndex int
	}{
		{array: []int{5, 19, -1, -4, 10}, max: 19, maxIndex: 1},
		{array: []int{-5}, max: -5, maxIndex: 0},
	}

	for _, item := range table {
		max, maxIndex := helpers.MaxArrayWithIndex(item.array)
		assert.Equal(t, item.max, max)
		assert.Equal(t, item.maxIndex, maxIndex)
	}
}

func TestMinWithIndex(t *testing.T) {
	table := []struct {
		array    []int
		min      int
		minIndex int
	}{
		{array: []int{5, 19, -1, -4, 10}, min: -4, minIndex: 3},
		{array: []int{-5}, min: -5, minIndex: 0},
	}

	for _, item := range table {
		min, minIndex := helpers.MinArrayWithIndex(item.array)
		assert.Equal(t, item.min, min)
		assert.Equal(t, item.minIndex, minIndex)
	}
}
