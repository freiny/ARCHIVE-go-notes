package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// When iterating over map keys, key order is not guaranteed

	// One Possible Output:
	// b 2
	// c 3
	// a 1

	// Another Possible Output:
	// c 3
	// a 1
	// b 2

	// Another Possible Output:
	// a 1
	// b 2
	// c 3
}
