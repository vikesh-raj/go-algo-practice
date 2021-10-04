package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikesh-raj/go-algo-practice/helpers"
)

func TestParseArrayOfInt(t *testing.T) {
	table := []struct {
		input string
		array []int
		fail  bool
	}{
		{input: "1 2 3", array: []int{1, 2, 3}},
		{input: "", array: []int{}},
		{input: "1 hello", fail: true},
		{input: "2 -1 3 1 1", array: []int{2, -1, 3, 1, 1}},
	}

	for _, item := range table {
		array, err := helpers.ParseArrayOfInt(item.input)
		if item.fail {
			assert.Error(t, err)
		} else {
			assert.Equal(t, item.array, array)
		}
	}
}

func TestParseArrayOfFloat(t *testing.T) {
	table := []struct {
		input string
		array []float64
		fail  bool
	}{
		{input: "1 2.3 3", array: []float64{1, 2.3, 3}},
		{input: "", array: []float64{}},
		{input: "1 hello", fail: true},
		{input: "2 -1.1 3.11 12.9 1", array: []float64{2, -1.1, 3.11, 12.9, 1}},
	}

	for _, item := range table {
		array, err := helpers.ParseArrayOfFloat(item.input)
		if item.fail {
			assert.Error(t, err)
		} else {
			assert.Equal(t, item.array, array)
		}
	}
}

func TestParseMatrixInt(t *testing.T) {
	table := []struct {
		input []string
		array [][]int
		fail  bool
	}{
		{input: []string{"1 2 3", "4 5 6", "7 8 9"}, array: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{input: []string{"1 2", "4 rr"}, fail: true},
	}

	for _, item := range table {
		array, err := helpers.ParseMatrixInt(item.input)
		if item.fail {
			assert.Error(t, err)
		} else {
			assert.Equal(t, item.array, array)
		}
	}
}
