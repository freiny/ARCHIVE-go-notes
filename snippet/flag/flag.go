package main

import (
	"flag"
	"fmt"
)

func main() {

	var name = flag.String("name", "Pat", "help message for name")
	var age = flag.Int("age", 30, "help message for age")

	fmt.Println("name", *name, "age", *age)
	flag.Parse()
	fmt.Println("name", *name, "age", *age)

	// INPUT:
	// $ command -age 22
	// OUTPUT:
	// name Pat age 30
	// name Pat age 22

	// INPUT:
	// $ command -name Rita -age 22
	// OUTPUT:
	// name Pat age 30
	// name Rita age 22
}
