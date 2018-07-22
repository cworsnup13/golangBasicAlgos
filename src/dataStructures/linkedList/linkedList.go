//package linkedList
package main

import "fmt"

type item struct {
	value interface{}
	next  *item
}

type linkedList struct {
	length int
	start  *item
	end    *item
}

type LinkedList interface {
	Append(newitem interface{})
	Prepend(newitem interface{})
	ToString() string
}

func New() LinkedList {
	l := linkedList{}
	return &l
}

func (l *linkedList) Append(val interface{}) {
	newitem := item{value:val}

	if l.start == nil {
		l.start = &newitem
		l.end = &newitem
	}

	l.end.next = &newitem
	l.end = &newitem
}

func (l *linkedList) Prepend(val interface{}) {
	newitem := item{value:val, next:l.start}

	if l.end == nil {
		l.end = &newitem
	}

	l.start = &newitem
}

func (l *linkedList) ToString() string {
	str := ""
	curr := l.start
	for curr != nil {
		str += " " + fmt.Sprint(curr.value)
		curr = curr.next
	}
	return str
}

func main() {
	l := New()
	l.Append(8)
	l.Append("abc")
	l.Prepend(123.456)
	fmt.Println(l.ToString())
}