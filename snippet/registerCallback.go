package main

import "fmt"

func main() {
	lib := Library{}.init()

	fmt.Println(lib.cba())
	fmt.Println(lib.cbb())
	fmt.Println(lib.cbc())
	// OUTPUT:
	// 0
	//
	// 0

	lib.registerCallback(cbUserDefinedA)
	lib.registerCallback(cbUserDefinedB)
	lib.registerCallback(cbUserDefinedC)

	fmt.Println(lib.cba())
	fmt.Println(lib.cbb())
	fmt.Println(lib.cbc())
	// OUTPUT:
	// 1
	// one
	// 1.1
}

func cbUserDefinedA() int     { return 1 }
func cbUserDefinedB() string  { return "one" }
func cbUserDefinedC() float64 { return 1.1 }

// Library callbacks
type Library struct {
	Callbacks
}

func (lib Library) init() Library {
	lib.cba = cbaStub
	lib.cbb = cbbStub
	lib.cbc = cbcStub
	return lib
}

func (lib *Library) registerCallback(i interface{}) {
	switch i.(type) {
	case (func() int):
		lib.cba = i.(func() int)
	case (func() string):
		lib.cbb = i.(func() string)
	case (func() float64):
		lib.cbc = i.(func() float64)
	default:
		fmt.Println("Unknown function type signature")
		return
	}
}

// Callbacks manipulated by the library/framework, but defined in user app
type Callbacks struct {
	cba func() int
	cbb func() string
	cbc func() float64
}

// stub functions for undefined callbacks
func cbaStub() int     { return 0 }
func cbbStub() string  { return "" }
func cbcStub() float64 { return 0.0 }
