package queue

import (
	"code.uber.internal/personal/golangBasicAlgos/src/dataStructures/linkedList"
)

type Queue interface {
	IsEmpty() bool
	Peek() *linkedList.Item
	Enqueue(interface{})
	Dequeue() *linkedList.Item
	ToString() string
}

type queue struct {
	ll linkedList.LinkedList
}

func New() Queue {
	l := linkedList.New()
	q := queue{
		ll: l,
	}
	return &q
}

func (q *queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}

func (q *queue) Peek() *linkedList.Item {
	return q.ll.Head()
}

func (q *queue) Enqueue(newitem interface{}) {
	q.ll.Append(newitem)
}

func (q *queue) Dequeue() *linkedList.Item {
	item := q.ll.Head()
	q.ll.DeleteHead()
	return item
}

func (q *queue) ToString() string {
	return q.ll.ToString()
}