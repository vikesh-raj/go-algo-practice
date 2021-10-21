package dynamicprogramming

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		items         []Item
		cap           int
		expectedValue int
		expected      []int
	}{
		{items: []Item{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, cap: 10, expectedValue: 10, expected: []int{1, 3}},
		{items: []Item{
			{465, 100}, {400, 85}, {255, 55}, {350, 45}, {650, 130},
			{1000, 190}, {455, 100}, {100, 25}, {1200, 190}, {320, 65},
			{750, 100}, {50, 45}, {550, 65}, {100, 50}, {600, 70}, {255, 40}},
			cap:           200,
			expectedValue: 1505,
			expected:      []int{7, 12, 14, 15},
		},
	}

	for _, test := range tests {
		v, got := Knapsack(test.items, test.cap)
		fmt.Println(got)
		assert.Equal(t, test.expectedValue, v)
		assert.Equal(t, test.expected, got)
	}
}
