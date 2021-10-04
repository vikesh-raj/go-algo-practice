package datastructures

import "strings"

type TrieNode struct {
	Word     string
	End      bool
	Children map[rune]TrieNode
}

type Trie struct {
	Root TrieNode
}

func NewTrieNode(word string, end bool) TrieNode {
	return TrieNode{
		Word:     word,
		End:      end,
		Children: make(map[rune]TrieNode),
	}
}

func (n TrieNode) String() string {
	var b strings.Builder
	if n.End {
		b.WriteRune('*')
	}
	b.WriteString(n.Word)
	if len(n.Children) != 0 {
		b.WriteRune(' ')
		b.WriteRune('[')
		i := 0
		total := len(n.Children)
		for k, value := range n.Children {
			b.WriteRune(k)
			b.WriteRune(':')
			b.WriteString(value.String())
			if i != total-1 {
				b.WriteRune(' ')
			}
			i++
		}
		b.WriteRune(']')
	}
	return b.String()
}

func NewTrie() *Trie {
	return &Trie{Root: NewTrieNode("", false)}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	runes := []rune(word)
	n := len(runes)
	for i, ch := range runes {
		child, ok := node.Children[ch]
		if !ok {
			node.Children[ch] = NewTrieNode(string(runes[:i+1]), i == n-1)
			child = node.Children[ch]
		}
		node = child
	}
}

func (t *Trie) Contains(str string) bool {
	node := t.Root
	for _, ch := range str {
		child, ok := node.Children[ch]
		if !ok {
			return false
		}
		node = child
	}
	return node.End
}
