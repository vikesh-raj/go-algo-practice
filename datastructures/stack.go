package datastructures

import (
	"github.com/vikesh-raj/go-algo-practice/helpers"
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Top() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

func (s *Stack) String() string {
	return helpers.ArrayToString(s.data)
}
