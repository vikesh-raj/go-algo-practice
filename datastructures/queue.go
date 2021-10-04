package datastructures

import (
	"container/list"

	"github.com/vikesh-raj/go-algo-practice/helpers"
)

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	return &Queue{data: list.New()}
}

func (q *Queue) Len() int {
	return q.data.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Peek() (int, bool) {
	front := q.data.Front()
	if front == nil {
		return 0, false
	}
	value := front.Value.(int)
	return value, true
}

func (q *Queue) Enqueue(value int) {
	q.data.PushBack(value)
}

func (q *Queue) Dequeue() (int, bool) {
	front := q.data.Front()
	if front == nil {
		return 0, false
	}
	value := front.Value.(int)
	q.data.Remove(front)
	return value, true
}

func (q *Queue) String() string {
	return helpers.ListToString(q.data)
}
