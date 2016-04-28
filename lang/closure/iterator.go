package main

import "fmt"

func main() {
	iter := iterator()

	iter(5)
	for iter() {
		fmt.Println("hello")
	}
	// OUTPUT:
	// hello
	// hello
	// hello
	// hello
	// hello

}

func iterator() func(...int) bool {
	counter := 0
	running := false
	return func(i ...int) bool {
		if i == nil {
			if counter <= 0 {
				running = false
			}
			counter--
		} else {
			counter = i[0]
			running = true
		}
		return running
	}
}
