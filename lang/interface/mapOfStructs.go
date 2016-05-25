package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func(each interface{}) {
		var id, v interface{}
		typeof := fmt.Sprintf("%T", each)
		switch typeof {
		case "main.as":
			id = each.(as).id
			v = each.(as).value
		case "main.bs":
			id = each.(bs).id
			v = each.(bs).value
		default:
			// fmt.Println("ID:", value, " ", reflect.TypeOf(value))
			return
		}

		fmt.Println("ID:", id, reflect.TypeOf(id), "VALUE:", v, reflect.TypeOf(v))
	}

	a := as{id: 0, value: "avalue"}
	b := bs{id: 1, value: 4.3}

	imap := map[string]interface{}{"a": a, "b": b}

	for _, each := range imap {
		f(each)
	}
	// OUTPUT:
	// ID: 0 int VALUE: avalue string
	// ID: 1 int VALUE: 4.3 float32
}

type as struct {
	id    int
	value string
}

type bs struct {
	id    int
	value float32
}
