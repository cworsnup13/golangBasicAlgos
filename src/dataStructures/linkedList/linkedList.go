package linkedList

import (
	"fmt"
)

type Item struct {
	Value interface{}
	Next  *Item
}

type linkedList struct {
	length int
	head   *Item
	tail   *Item
}

type LinkedList interface {
	Append(newitem interface{})
	Prepend(newitem interface{})
	DeleteHead()
	DeleteTail()
	ToString() string
	IsEmpty() bool
	Head() *Item
	Tail() *Item
	Find(Value interface{}, callback func(interface{}) bool) *Item
	Delete(val interface{}) *Item
}

func New() LinkedList {
	l := linkedList{}
	return &l
}

func (l *linkedList) Append(val interface{}) {
	newitem := Item{Value: val}

	if l.head == nil {
		l.head = &newitem
		l.tail = &newitem
		l.tail.Next = nil
		return
	}

	l.tail.Next = &newitem
	l.tail = &newitem
}

func (l *linkedList) Prepend(val interface{}) {
	newitem := Item{Value: val, Next: l.head}

	if l.tail == nil {
		l.tail = &newitem
	}

	l.head = &newitem
}

func (l *linkedList) DeleteHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.Next
}

func (l *linkedList) DeleteTail() {
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}

	curr := l.head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	l.tail = curr
}

func (l *linkedList) Delete(val interface{}) (deletedNode *Item) {
	if l.head == nil {
		return
	}

	for l.head != nil && l.head.Value == val {
		deletedNode = l.head
		l.head = l.head.Next
	}

	currentNode := l.head
	if currentNode != nil {
		for currentNode.Next != nil {
			if currentNode.Value == val {
				deletedNode = currentNode.Next
				currentNode.Next = currentNode.Next.Next
			} else {
				currentNode = currentNode.Next
			}
		}
	}

	if l.tail.Value == val {
		l.tail = currentNode
	}

	return
}

func (l *linkedList) ToString() string {
	str := ""
	curr := l.head
	for curr != nil {
		str += " " + fmt.Sprint(curr.Value)
		curr = curr.Next
	}
	return str
}

func (l *linkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *linkedList) Head() *Item {
	return l.head
}

func (l *linkedList) Tail() *Item {
	return l.tail
}

func (l *linkedList) Find(Value interface{}, callback func(interface{}) bool) *Item {
	if l.head == nil {
		return nil
	}

	currentNode := l.head
	for currentNode != nil {
		if callback != nil && callback(currentNode.Value) {
			return currentNode
		}

		if Value != nil && currentNode.Value == Value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}
