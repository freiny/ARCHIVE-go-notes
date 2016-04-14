package main

import "fmt"

func main() {
	fn := func() {
		func(a int, b int) { fmt.Println(a, b) }(1, 2)
	}
	fn()
}
