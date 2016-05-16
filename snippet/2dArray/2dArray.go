package main

import "fmt"

func main() {
	arr := make([][4]byte, 4)

	counter := 0
	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr[0]); c++ {
			arr[r][c] = byte(counter)
			counter++
		}
	}

	fmt.Println(arr)
	// OUTPUT:
	// [[0 1 2 3][4 5 6 7][8 9 10 11][12 13 14 15]]

	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr[0]); c++ {
			fmt.Printf("%x", arr[r][c])
		}
		fmt.Println()
	}
	// OUTPUT:
	// 0123
	// 4567
	// 89ab
	// cdef
}
