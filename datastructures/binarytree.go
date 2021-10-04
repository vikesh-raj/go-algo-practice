package datastructures

import (
	"strconv"
	"strings"
)

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func (n BSTNode) String() string {
	var b strings.Builder
	b.WriteRune('{')
	b.WriteString(strconv.FormatInt(int64(n.Value), 10))
	if n.Left != nil {
		b.WriteString(" L:")
		b.WriteString(n.Left.String())
	}
	if n.Right != nil {
		b.WriteString(" R:")
		b.WriteString(n.Right.String())
	}
	b.WriteRune('}')
	return b.String()
}

func (b *BST) Insert(value int) {
	b.Root = insertNode(b.Root, value)
}

func insertNode(root *BSTNode, value int) *BSTNode {
	if root == nil {
		root = &BSTNode{Value: value}
		return root
	}
	if value < root.Value {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}
	return root
}
