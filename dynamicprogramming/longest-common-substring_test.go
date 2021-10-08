package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstring(t *testing.T) {
	assert.Equal(t, "ood", LongestCommonSubstring("tofoodie", "toody"))
	assert.Equal(t, "programming", LongestCommonSubstring("dynamicprogramming", "functionalprogramming"))
	assert.Equal(t, "", LongestCommonSubstring("unrelated", "zzzzz"))
	assert.Equal(t, "", LongestCommonSubstring("", "not empty string"))
}
