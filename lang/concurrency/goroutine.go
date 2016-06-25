package main

import "fmt"

func main() {
	go dispLower()
	go dispUpper()

	var input string
	fmt.Println("Press Enter")
	fmt.Scanln(&input)

	// OUTPUT (unpredictable):
	// Press Enter
	// abcdefghijklAmnBoCpDqErFstGuHvIwJxKyLzMNOPQRSTUVWXYZ
	// Done
}

func dispLower() {
	s := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
}

func dispUpper() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
}
