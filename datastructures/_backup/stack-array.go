package main

import (
	"errors"
	"fmt"
)

func main() {

	s := stack{}.new()
	// OUTPUT: {3 0 [0 0 0]}

	s.push(1)
	// OUTPUT: {3 1 [1 0 0]}

	s.push(2)
	// OUTPUT: {3 2 [1 2 0]}

	s.push(3)
	// OUTPUT: {3 3 [1 2 3]}

	// s.push(4)
	// // OUTPUT: panic: ERROR: push to stack failed top of stack reached

	fmt.Println(s.pop())
	// OUTPUT: {3 2 [1 2 3]}
	// 3

	fmt.Println(s.pop())
	// OUTPUT: {3 1 [1 2 3]}
	// 2

	fmt.Println(s.pop())
	// OUTPUT: {3 0 [1 2 3]}
	// 1

	// fmt.Print(s.pop())
	// // OUTPUT: panic: ERROR: pop to stack failed bottom of stack reached

}

const max = 3

type stack struct {
	sizeMax int
	size    int
	list    [max]int
}

func (s stack) new() stack {
	s.sizeMax = len(s.list)
	s.print()
	return s
}

func (s *stack) push(item int) {
	if s.size < s.sizeMax {
		s.list[s.size] = item
		s.size++
		s.print()
	} else {
		panic(errors.New("ERROR: push to stack failed top of stack reached"))
	}
}

func (s *stack) pop() int {
	if s.size > 0 {
		s.size--
		s.print()
		return s.list[s.size]
	} else {
		panic(errors.New("ERROR: pop to stack failed bottom of stack reached"))
	}
}

func (s stack) print() {
	fmt.Println(s)
}
