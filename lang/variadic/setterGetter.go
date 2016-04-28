package main

import "fmt"

func main() {
	f := Framework{}

	// set property: f.Property(value)
	f.Property("hello")

	// get property: f.Property()
	get := f.Property()

	fmt.Println(get)
	// OUTPUT: hello
}

type Framework struct {
	property string
}

func (f *Framework) Property(i ...interface{}) string {
	if len(i) > 0 {
		f.property = i[0].(string)
		return ""
	} else {
		return f.property
	}
}
