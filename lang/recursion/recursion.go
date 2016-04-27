package main

import "fmt"

func main() {
	print("abcde")
	// OUTPUT:
	// abcde
	// abcd
	// abc
	// ab
	// a
}

func print(s string) {
	if len(s) > 0 {
		fmt.Println(s)
		print(s[0 : len(s)-1])
	}
}
