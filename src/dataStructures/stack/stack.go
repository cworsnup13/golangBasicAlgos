package stack

import (
	"code.uber.internal/personal/golangBasicAlgos/src/dataStructures/linkedList"
)

type Stack interface {
	IsEmpty() bool
	Peek() *linkedList.Item
	Push(interface{})
	Pop() *linkedList.Item
	ToString() string
}

type stack struct {
	ll linkedList.LinkedList
}

func New() Stack {
	l := linkedList.New()
	q := stack{
		ll: l,
	}
	return &q
}

func (s *stack) IsEmpty() bool {
	return s.ll.IsEmpty()
}

func (s *stack) Peek() *linkedList.Item {
	return s.ll.Tail()
}

func (s *stack) Push(newitem interface{}) {
	s.ll.Append(newitem)
}

func (s *stack) Pop() *linkedList.Item {
	item := s.ll.Tail()
	s.ll.DeleteTail()
	return item
}

func (s *stack) ToString() string {
	return s.ll.ToString()
}