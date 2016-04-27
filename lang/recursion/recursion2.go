package main

import "fmt"

func main() {
	a := node{value: "A"}
	b := node{value: "B"}
	c := node{value: "C"}
	d := node{value: "D"}
	e := node{value: "E"}

	a.next = &b
	b.next = &c
	c.next = &d
	d.next = &e

	a.iter(print)
	// OUTPUT:
	// A
	// B
	// C
	// D
	// E
}

type node struct {
	value string
	next  *node
}

func (n *node) iter(fn func(*node)) {
	fn(n)
	if n.next != nil {
		n.next.iter(fn)
	}
}

func print(n *node) {
	fmt.Println(n.value)
}
