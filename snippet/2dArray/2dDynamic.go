package main

import "fmt"

func main() {

	arr := make([][]int, 0)
	fmt.Println(arr)
	// OUTPUT: []

	rows, cols := 3, 4
	counter := 0
	for r := 0; r < rows; r++ {
		arr = append(arr, make([]int, cols, cols))
		for c := 0; c < cols; c++ {
			arr[r][c] = counter
			counter++
		}
	}
	fmt.Println(arr)
	// OUTPUT:
	// [[0 1 2 3][4 5  6 7][8 9 10 11]]
}
