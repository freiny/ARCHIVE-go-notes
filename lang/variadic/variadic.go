package main

import "fmt"

func main() {

	printStrings("a", "b", "c")
	// OUTPUT: a b c

	list := []string{"1", "2", "3"}
	printStrings(list...)
	// OUTPUT: 1 2 3

}

func printStrings(s ...string) {
	for _, each := range s {
		fmt.Print(each, " ")
	}
	fmt.Println()
}
