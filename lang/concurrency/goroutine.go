package main

import (
	"fmt"
	"time"
)

func main() {
	go dispLower()
	go dispUpper()

	// goroutines exit when main() ends
	// so, delay end of main()
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// POSSIBLE OUTPUT (unpredictable):
	// abcAdefghiBjCklmDnEoFpGqHrIstJuKvwLxMyNzOPQRSTUVWXYZ

	// POSSIBLE OUTPUT (unpredictable):
	// aAbcBdCeDfEgFhGHiIjJkKlLmMnNoOPpQqRrSsTtUuVvWwXYxZyz

	// POSSIBLE OUTPUT (unpredictable):
	// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

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
