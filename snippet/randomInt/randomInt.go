package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	for i := 0; i < 10; i++ {
		n := randomInt(2)
		fmt.Print(n, " ")
	}
	// Some possible OUTPUTS:

	// A possilbe output:
	// 0 0 1 0 1 1 0 0 1 1

	// Another possilbe output:
	// 1 0 1 0 0 0 1 0 0 1

	// Another possilbe output:
	// 0 0 1 1 1 0 1 1 0 0
}

func randomInt(max int) int {
	return rand.Intn(max)
}
