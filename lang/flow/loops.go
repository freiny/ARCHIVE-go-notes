package main

import "fmt"

func main() {

	// infinite for loop
	i := 0
	for {
		if i > 2 {
			break
		} else {
			fmt.Printf("%d ", i)
			i++
		}
	}
	fmt.Println()

	// OUTPUT:
	// 0 1 2

	// Standard for loop
	for i := 0; i < 4; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// OUTPUT:
	// 0 1 2 3

	// Range for loop
	str := "hey"
	for index, element := range str {
		fmt.Println(index, string(key))
	}
	fmt.Println()

	// OUTPUT:
	// 0 h
	// 1 e
	// 2 y

}
