package linkedList

import "fmt"

type Item struct {
	value interface{}
	next  *Item
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
}

func New() LinkedList {
	l := linkedList{}
	return &l
}

func (l *linkedList) Append(val interface{}) {
	newitem := Item{value: val}

	if l.head == nil {
		l.head = &newitem
		l.tail = &newitem
		l.tail.next = nil
		return
	}

	l.tail.next = &newitem
	l.tail = &newitem
}

func (l *linkedList) Prepend(val interface{}) {
	newitem := Item{value: val, next: l.head}

	if l.tail == nil {
		l.tail = &newitem
	}

	l.head = &newitem
}

func (l *linkedList) DeleteHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

func (l *linkedList) DeleteTail() {
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}

	curr := l.head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
	l.tail = curr
}

func (l *linkedList) ToString() string {
	str := ""
	curr := l.head
	for curr != nil {
		str += " " + fmt.Sprint(curr.value)
		curr = curr.next
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
