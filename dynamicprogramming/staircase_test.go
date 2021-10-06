package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {
	assert.Equal(t, 8, StairCaseCombinations(5, 2))
	assert.Equal(t, 34, StairCaseCombinations(8, 2))
	assert.Equal(t, 8, StairCaseCombinations(4, 4))

	assert.Equal(t, 24, StairCaseCombinations(6, 3))
}
