package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 55, Fibonacci(10))
	assert.Equal(t, 144, Fibonacci(12))
}
