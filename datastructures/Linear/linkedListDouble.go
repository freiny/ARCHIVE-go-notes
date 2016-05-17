package main

import "fmt"

func main() {

	l := DoublyLinkedList{}
	l.append(0)
	l.append(1)
	l.append(2)
	l.prepend(-1)
	l.prepend(-2)
	l.prepend(-3)
	l.traverseNext(print)
	// OUTPUT:
	// -3 true false
	// -2 false false
	// -1 false false
	// 0 false false
	// 1 false false
	// 2 false true

	l.deleteHead()
	l.deleteHead()
	l.deleteTail()
	l.traverseNext(print)
	// OUTPUT:
	// -1 true false
	// 0 false false
	// 1 false true

	l.traversePrev(print)
	// OUTPUT:
	// 1 false true
	// 0 false false
	// -1 true false

}

func print(n *node) bool {
	fmt.Println(n.value, n.prev == nil, n.next == nil)
	return false
}

type node struct {
	value int
	prev  *node
	next  *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
}

func (l *DoublyLinkedList) prepend(value int) {
	n := node{value: value}
	if l.head == nil {
		l.head = &n
		l.tail = &n
	} else {
		n.next = l.head
		l.head.prev = &n
		l.head = &n
	}
}

func (l *DoublyLinkedList) append(value int) {
	n := node{value: value}
	if l.head == nil {
		l.head = &n
		l.tail = &n
	} else {
		n.prev = l.tail
		l.tail.next = &n
		l.tail = l.tail.next
	}
}

func (l *DoublyLinkedList) deleteHead() {
	n := l.head.next
	n.prev = nil
	l.head = n
}

func (l *DoublyLinkedList) deleteTail() {
	n := l.tail.prev
	n.next = nil
	l.tail = n
}

func (l *DoublyLinkedList) traverseNext(f func(*node) bool) {
	for n := l.head; n != nil; n = n.next {
		if f(n) {
			break
		}
	}
}

func (l *DoublyLinkedList) traversePrev(f func(*node) bool) {
	for n := l.tail; n != nil; n = n.prev {
		if f(n) {
			break
		}
	}
}
