package main

import "fmt"

func main() {

	g := Group{
		"Accounting",
		[]string{"Jane", "Bob", "Pat"},
	}

	fmt.Println(g)
	// OUTPUT:
	// {Accounting [Jane Bob Pat]}

}

type Group struct {
	title   string
	members []string
}
