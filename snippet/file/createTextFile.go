package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	content := []string{
		"testing...",
		"testing...",
		"1...2...3...",
	}
	createTextFile("out.txt", content...)
}

func createTextFile(filename string, content ...string) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, each := range content {
		fmt.Fprintln(writer, each)
	}

	writer.Flush()
}
