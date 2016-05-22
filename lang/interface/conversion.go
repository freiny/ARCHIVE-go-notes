package main

import "fmt"

func main() {
	e := &Element{Value: 3}
	fmt.Println(e.Value)
	// OUTPUT: 3

	f := func(e *Element) {
		var i interface{}
		i = e.Value.(int) + 1
		e.Value = i
	}
	e.run(f)
	fmt.Println(e.Value)
	// OUTPUT: 4
}

type Element struct {
	Value interface{}
}

func (e *Element) run(f func(*Element)) {
	f(e)
}
