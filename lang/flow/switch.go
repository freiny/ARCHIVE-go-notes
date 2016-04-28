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
}
