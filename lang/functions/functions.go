package main

import (
	"fmt"
	"strconv"
)

func main() {

	printA()
	// OUTPUT: A

	b := returnB()
	fmt.Println(b)
	// OUTPUT: B

	sum := sum(3, 4)
	fmt.Println(sum)
	// OUTPUT: SUM: 7

	ada, bab := getLengths("Ada", "Babbage")
	fmt.Println(ada, bab)
	// OUTPUT: 3 7

	length, _ := getLengths("Carl", "")
	fmt.Println(length)
	// OUTPUT: 4

	funcA, funcB := returnFuncs()
	funcA()
	funcB()
	// OUTPUT:
	// AAA
	// BBB

	diff := func(a int, b int) int {
		return a - b
	}

	d := diff(5, 7)
	fmt.Println(d)
	// OUTPUT: -2

}

func printA() {
	fmt.Println("A")
}

func returnB() string {
	return "B"
}

func sum(a int, b int) string {
	return "SUM: " + strconv.Itoa(a+b)
}

func getLengths(a string, b string) (int, int) {
	return len(a), len(b)
}

func returnFuncs() (func(), func()) {
	return func() {
			fmt.Println("AAA")
		}, func() {
			fmt.Println("BBB ")
		}
}
