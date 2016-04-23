package main

import "fmt"

func main() {
	o := Object{}

	fmt.Println(o)
	// OUTPUT: {{0 }{0 }}

	o.add()
	fmt.Println(o)
	// OUTPUT: {{1 HelloWorld}{1 HelloWorld}}

}

// Object to test
type Object struct {
	so SubObject
	SubObject
}

func (o Object) init() Object {
	return o
}

func (o *Object) add() {
	o.so.i++
	o.i++
	o.so.s = "HelloWorld"
	o.s = "HelloWorld"
}

// SubObject to test
type SubObject struct {
	i int
	s string
}
