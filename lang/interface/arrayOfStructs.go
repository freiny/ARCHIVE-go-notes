package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func(value interface{}) {
		fmt.Println(value, " ", reflect.TypeOf(value))
	}

	a := as{id: 0, value: "avalue"}
	b := bs{id: 1, value: 4.3}
	arr := []interface{}{a, b}

	for _, each := range arr {
		f(each)
	}
	// OUTPUT:
	// {0 avalue} main.as
	// {}1 4.3} main.bs
}

type as struct {
	id    int
	value string
}

type bs struct {
	id    int
	value float32
}
