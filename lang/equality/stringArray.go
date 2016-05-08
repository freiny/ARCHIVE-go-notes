package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a []string
	fmt.Println(a, []string{})
	fmt.Println(reflect.DeepEqual(a, []string{}))
	// OUTPUT:
	// [] []
	// false

	fmt.Println(reflect.DeepEqual([]string{}, []string{}))
	fmt.Println(reflect.DeepEqual([]string{}, []string{""}))
	fmt.Println(reflect.DeepEqual([]string{""}, []string{""}))
	// OUTPUT:
	// true
	// false
	// true

	fmt.Println(reflect.DeepEqual([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
	fmt.Println(reflect.DeepEqual([]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}))
	// OUTPUT:
	// true
	// false

	fmt.Println()
}
