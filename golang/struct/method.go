package main

import "fmt"

func main() {

	name := Name{"Bruce", "Wayne"}
	name.greeting()

	// OUTPUT:
	// Hi my name is Bruce Wayne

}

type Name struct {
	first string
	last  string
}

func (n Name) greeting() {
	fmt.Println("Hi my name is ", n.first, n.last)
}
