package main

import (
	"fmt"
	"strings"
)

func main() {
	b := []bool{
		strings.ContainsAny("puppy", "u"),
		strings.ContainsAny("puppy", "ui"),
		strings.ContainsAny("puppy", "aeiou"),
		// true

		strings.ContainsAny("puppy", "aeio"),
		strings.ContainsAny("puppy", "P"),
		strings.ContainsAny("", ""),
		strings.ContainsAny("ab", ""),
		// false
	}
	fmt.Println(b)

}
