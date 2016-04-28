package main

import "fmt"

func main() {
	myPrint("abc", 1, 2, 3, []int{4, 5, 6})
	// OUTPUT: abc 1 2 3 [4 5 6]
}

func myPrint(i ...interface{}) {
	fmt.Println(i...)
}
