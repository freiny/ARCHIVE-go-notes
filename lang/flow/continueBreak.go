package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 0; i < 8; i++ {
		if i > 3 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	// OUTPUT: 0 1 2 3

	for i := 0; i < 8; i++ {
		if math.Mod(float64(i), float64(2)) > 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	// OUTPUT: 0 2 4 6

}
