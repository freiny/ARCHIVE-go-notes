package main

import "fmt"

func main() {

	team := map[string]bool{}
	team["jane"] = true
	team["john"] = true
	team["mel"] = true

	delete(team, "jane")

	fmt.Println(team["bob"], team["jane"], team["mel"])
	// OUTPUT: false true true

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
