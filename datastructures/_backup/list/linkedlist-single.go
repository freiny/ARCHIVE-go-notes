package main

import "fmt"

func main() {

	l := list{}.new()
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	l.insert(5)
	l.insert(4)

	print := func(n *node) bool {
		fmt.Print(n.data, " ")
		return false
	}

	l.forEach(print)
	fmt.Println()
	// OUTPUT: 1 2 3 4 5 4

	add2 := func(n *node) bool {
		n.data = n.data + 2
		return false
	}

	l.forEach(add2)
	l.forEach(print)
	fmt.Println()
	// OUTPUT: 3 4 5 6 7 6

	l.delete(3)
	l.delete(5)
	l.delete(6)
	l.forEach(print)
	// OUTPUT: 4 7 6

}

type list struct {
	head *node
	tail *node
	size int
}

func (l list) new() list {
	return l
}

func (l list) forEach(fn func(*node) bool) {
	for node := l.head; node != nil; node = node.next {
		// if any fn iteration returns true then foreach is complete, so break
		if fn(node) {
			break
		}
	}
}

func (l *list) insert(data int) {
	newNode := node{data, nil}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
	}
}

func (l *list) delete(data int) {

	delete := func(n *node) bool {
		if n == l.tail {
			return true
		}

		if n == l.head {
			if data == n.data {
				l.head = n.next
				n = nil
				return true
			}
		}

		if data == n.next.data {
			n.next = n.next.next
			return true
		}
		return false
	}
	l.forEach(delete)
}

type node struct {
	data int
	next *node
}
