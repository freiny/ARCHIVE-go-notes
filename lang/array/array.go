package main

import "fmt"

func main() {

	evens := []int{2, 4, 6, 8}

	fmt.Println(evens, len(evens))
	// OUTPUT: [2 4 6 8] 4

	evens = append(evens, 10, 12, 14)
	fmt.Println(evens, len(evens))
	// OUTPUT: [2 4 6 8 10 12 14] 7

	fmt.Println(evens[2], evens[3])
	// OUTPUT: 6 8

	var list [3]string
	list[0] = "a"
	list[1] = "bb"
	list[2] = "ccc"

	for i, v := range list {
		fmt.Println(i, v)
	}
	// OUTPUT:
	// 0 a
	// 1 bb
	// 2 ccc
}
