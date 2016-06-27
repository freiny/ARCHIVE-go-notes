package main

import "fmt"

func main() {

	cap := 3

	tasks := make([]func(), 0, cap)
	for i := 0; i < cap; i++ {
		n := i
		tasks = append(tasks, func() { fmt.Print(n) })
	}
	for i := 0; i < cap; i++ {
		tasks[i]()
	}
	fmt.Println()
	// OUTPUT: 012

}
