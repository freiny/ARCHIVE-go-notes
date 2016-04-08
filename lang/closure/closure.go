package main

import (
	"fmt"
	"math"
)

func main() {

	print := fn()

	for i := 0; i < 9; i++ {
		print()
	}

	// OUTPUT: 1 2 3 1 2 3 1 2 3
}

func fn() func() {
	n := 1
	max := 3
	return func() {
		fmt.Print(n, " ")
		n = int(math.Mod(float64(n), float64(max)))
		n++
	}
}
