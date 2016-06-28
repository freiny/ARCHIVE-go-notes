package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd("ls -a")
	cmd("mkdir newfolder")
	cmd("ls -a")
	cmd("rm -rf newfolder")
	cmd("ls -a")
}

func cmd(s string) {
	parts := strings.Split(s, " ")
	command := parts[0]
	args := parts[1:]

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command", err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}
