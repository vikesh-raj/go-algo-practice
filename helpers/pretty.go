package helpers

import (
	"container/list"
	"fmt"
	"strings"
)

func ArrayToString(array []int) string {
	if array == nil {
		return "nil"
	}
	var b strings.Builder
	b.WriteRune('[')
	n := len(array)
	for i, value := range array {
		b.WriteString(fmt.Sprint(value))
		if i != n-1 {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(']')
	return b.String()
}

func ListToString(list *list.List) string {
	if list == nil {
		return "nil"
	}
	var b strings.Builder
	b.WriteRune('[')
	for e := list.Front(); e != nil; e = e.Next() {
		b.WriteString(fmt.Sprint(e.Value))
		if e.Next() != nil {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(']')
	return b.String()
}
