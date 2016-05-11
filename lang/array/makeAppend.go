package main

import "fmt"

func main() {

	out := func(i []int) {
		fmt.Println(len(i), cap(i), i)
	}

	a := make([]int, 2)
	out(a)
	// OUTPUT: 2 2 [0 0]

	a = append(a, 7)
	out(a)
	// OUTPUT: 3 4 [0 0 7]

	a = append(a, 8, 9)
	out(a)
	// OUTPUT: 5 8 [0 0 7 8 9]

	b := make([]int, 0)
	out(b)
	// OUTPUT: 0 0 []
	b = append(b, 7)
	out(b)
	// OUTPUT: 1 1 [7]

	c := make([]int, 0, 8)
	out(c)
	// OUTPUT: 0 8 []

	c = append(c, 7)
	out(c)
	// OUTPUT: 1 8 [7]

	s := make([]string, 0, 256)
	fmt.Println(len(s), cap(s), s)
	// OUTPUT: 0 256 []

	s = append(s, "a", "bc")
	s = append(s, "def")
	fmt.Println(len(s), cap(s), s)
	// OUTPUT: 3 256 [a bc def]
}
