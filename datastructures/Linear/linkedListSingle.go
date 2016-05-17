package main

import "fmt"

func main() {

	l := LinkedList{}
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	l.traverse(print)
	// OUTPUT:
	// 1 false
	// 2 false
	// 3 false
	// 4 true

	l.delete()
	l.delete()
	l.traverse(print)
	// OUTPUT:
	// 1 false
	// 2 true

}

func print(n *node) bool {
	fmt.Println(n.value, n.next == nil)
	return false
}

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) insert(value int) {
	n := node{value: value}
	if l.head == nil {
		l.head = &n
		l.tail = &n
	}
	l.tail.next = &n
	l.tail = l.tail.next
}

func (l *LinkedList) delete() {
	f := func(n *node) bool {
		if n.next.next == nil {
			l.tail = n
			l.tail.next = nil
			return true
		}
		return false
	}

	l.traverse(f)
}

func (l *LinkedList) traverse(f func(*node) bool) {
	for n := l.head; n != nil; n = n.next {
		if f(n) {
			break
		}
	}
}
