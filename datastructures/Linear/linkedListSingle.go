package main

import "fmt"

func main() {

	l := LinkedList{}
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.traverse(print)
	// OUTPUT:
	// 1 true
	// 2 true
	// 3 false

}

func print(n *node) {
	fmt.Println(n.value, n.next != nil)
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

func (l LinkedList) traverse(f func(*node)) {
	for n := l.head; n != nil; n = n.next {
		f(n)
	}
}
