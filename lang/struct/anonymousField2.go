package main

import "fmt"

func main() {
	f := Framework{}
	f.speak()
	// OUTPUT: Hello

	nf := NewFramework{Framework{}}
	nf.speak()
	// OUTPUT: Hello World

	nf.speakBoth()
	// OUTPUT:
	// Hello
	// Hello World
}

// Framework ...
type Framework struct {
}

func (f Framework) speak() {
	fmt.Println("Hello")
}

// NewFramework ...
type NewFramework struct {
	Framework
}

func (nf NewFramework) speak() {
	fmt.Println("Hello World")
}

// NewFramework still has access to Framework methods
func (nf NewFramework) speakBoth() {
	nf.Framework.speak()
	nf.speak()
}
