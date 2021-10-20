package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, "abacc", LongestCommonSubsequence("xabayycczz", "pabqaqcqcr"))
	assert.Equal(t, "XYZW", LongestCommonSubsequence("ZXVVYZW", "XKYKZPW"))
	assert.Equal(t, "ac", LongestCommonSubsequence("abc", "ac"))
	assert.Equal(t, "", LongestCommonSubsequence("unrelated", "zzzzz"))
	assert.Equal(t, "", LongestCommonSubsequence("", "not empty string"))
}
