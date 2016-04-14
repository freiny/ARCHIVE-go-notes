package main

import "fmt"

func main() {

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(list[0:1], list[5:6], list[7:7])
	// OUTPUT: [0] [5] []

	fmt.Println(list[0:2], list[4:6], list[7:9])
	// OUTPUT: [0 1] [4 5] [7 8]

	slice := list[3:6]
	fmt.Println(slice)
	// OUTPUT: [3 4 5]

	slice = append(slice, 7)
	fmt.Println(slice)
	// OUTPUT: [3 4 5 7]

	slice = append(slice, 8, 9)
	fmt.Println(slice)
	// OUTPUT: [3 4 5 7 8 9]

	extra := list[0:3]

	slice = append(slice, extra...)
	fmt.Println(slice)
	// OUTPUT: [3 4 5 7 8 9 0 1 2]

}
