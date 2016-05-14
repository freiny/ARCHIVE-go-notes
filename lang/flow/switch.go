package main

import "fmt"

func main() {

	mode := "Production"

	switch mode {
	case "Development":
		fmt.Println("Dev")
	case "Production":
		fmt.Println("Prod")
	default:
		fmt.Println("None")
	}
	// OUTPUT: Prod

	msg := "hello"

	switch {
	case len(msg) < 5:
		fmt.Println("len < 5")
	case len(msg) > 5:
		fmt.Println("len > 5")
	case msg == "hello":
		fmt.Println(msg, "world")
	default:
		fmt.Println("default")
	}
	// OUTPUT: hello world

}
