package main

import "fmt"

func init() {
	fmt.Println("init() running...")
}

func main() {
	fmt.Println("main() running...")
	// OUTPUT:
	// init() running...
	// main() running...
}
