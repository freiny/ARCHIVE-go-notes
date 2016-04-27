package main

import (
	"fmt"
	"strings"
)

func main() {

	b := []bool{
		strings.Contains("Hello World!", ""),
		strings.Contains("Hello World!", " "),
		strings.Contains("Hello World!", "!"),
		strings.Contains("Hello World!", " Wo"),
		strings.Contains("Hello World!", "wo"),
	}
	fmt.Println(b)
	// OUTPUT: [true true true true false]
}
