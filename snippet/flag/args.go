package main

import (
	"fmt"
	"os"
)

func main() {
	for _, each := range os.Args {
		fmt.Println(each)
	}
}
