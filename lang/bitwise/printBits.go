package main

import (
	"fmt"
	"strconv"
)

func main() {
	var u uint64
	u = 18446744073709551615
	printBinary(u)
	printHex(u)
}

func printBinary(u uint64) {
	fmt.Println(strconv.FormatUint(u, 2))
}

func printHex(u uint64) {
	fmt.Println(strconv.FormatUint(u, 16))
}
