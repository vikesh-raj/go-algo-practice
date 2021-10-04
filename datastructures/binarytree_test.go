package datastructures_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikesh-raj/go-algo-practice/datastructures"
)

func TestBinaryTree(t *testing.T) {
	tree := datastructures.BST{}
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)

	fmt.Println(tree.Root)
	assert.Equal(t, 21, tree.Root.Value)
}
