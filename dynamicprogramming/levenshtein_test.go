package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevenshteinDistance(t *testing.T) {
	assert.Equal(t, 3, LevenshteinDistance("kitten", "sitting"))
	assert.Equal(t, 2, LevenshteinDistance("anchorge", "achorae"))
	assert.Equal(t, 16, LevenshteinDistance("", "not empty string"))
	assert.Equal(t, 2, LevenshteinDistance("hi", ""))
}
