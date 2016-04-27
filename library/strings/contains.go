package main

import (
	"fmt"
	"strings"
)

func main() {

	b := []bool{
		strings.Contains("Battle Star", "e"),
		strings.Contains("Battle Star", "at"),
		strings.Contains("Battle Star", " "),
		strings.Contains("Battle Star", ""),
		strings.Contains("", ""),
		// true

		strings.Contains("Battle Star", "ate"),
		strings.Contains("Battle Star", "ba"),
		strings.Contains("Battle Star", "z"),
		// false
	}
	fmt.Println(b)
}
