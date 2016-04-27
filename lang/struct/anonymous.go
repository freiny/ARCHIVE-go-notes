package main

import "fmt"

func main() {
	names := []struct {
		first, last string
	}{
		{"john", "pho"},
		{"pat", "doe"},
		{"jane", "snow"},
	}

	for _, name := range names {
		fmt.Println(name)
	}
	// OUTPUT:
	// {john pho}
	// {pat doe}
	// {jane snow}

}
