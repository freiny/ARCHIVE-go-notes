package main

import "fmt"

func main() {

	cap := 3

	tasksA := make([]func(), 0, cap)
	for i := 0; i < cap; i++ {
		tasksA = append(tasksA, func() { fmt.Print(i) })
	}
	for i := 0; i < cap; i++ {
		tasksA[i]()
	}
	fmt.Println()
	// OUTPUT: 333

	tasksB := make([]func(), 0, cap)
	for i := 0; i < cap; i++ {
		n := i
		tasksB = append(tasksB, func() { fmt.Print(n) })
	}
	for i := 0; i < cap; i++ {
		tasksB[i]()
	}
	fmt.Println()
	// OUTPUT: 012

}
