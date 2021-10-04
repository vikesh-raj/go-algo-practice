package datastructures_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikesh-raj/go-algo-practice/datastructures"
)

func TestTrie(t *testing.T) {
	trie := datastructures.NewTrie()
	fmt.Println(trie.Root)
	trie.Insert("this")
	trie.Insert("that")
	trie.Insert("title")
	trie.Insert("hello")
	trie.Insert("help")
	fmt.Println(trie.Root)

	assert.True(t, trie.Contains("hello"))
	assert.True(t, trie.Contains("that"))
	assert.False(t, trie.Contains("thi"))
}
