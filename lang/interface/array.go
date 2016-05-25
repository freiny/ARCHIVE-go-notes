package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func(value interface{}) {
		fmt.Println(value, " ", reflect.TypeOf(value))
	}

	arr := []interface{}{"abc", 'd', 1, 2.3}

	for _, each := range arr {
		f(each)
	}
	// OUTPUT:
	// abc string
	// 100 int32
	// 1 int
	// 2.3 float64
}
