package main

import "fmt"

func main() {
	printItems(1, "two", 3.0, Four{4, "four"})
}

type Four struct {
	i int
	s string
}

func printItems(items ...interface{}) {
	for _, each := range items {
		fmt.Println(each)
	}
}
