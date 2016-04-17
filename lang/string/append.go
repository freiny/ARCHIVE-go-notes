package main

import (
	"bytes"
	"fmt"
)

func main() {

	var buffer bytes.Buffer

	s := "abcde"

	for _, v := range s {
		buffer.WriteString(string(v))
		fmt.Println(buffer.String())
	}
	// OUTPUT:
	// a
	// ab
	// abc
	// abcd
	// abcde

}
