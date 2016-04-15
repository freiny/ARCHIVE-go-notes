package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	ls()
	// OUTPUT:
	// .
	// ..
	// runCommand.go
	// runCommand

	mkdir("xyzFolder")
	ls()
	// OUTPUT:
	// .
	// ..
	// runCommand.go
	// runCommand
	// xyzFolder

	rm("xyzFolder")
	ls()
	// OUTPUT:
	// .
	// ..
	// runCommand.go
	// runCommand
}

func mkdir(s string) {
	err := exec.Command("mkdir", s).Run()
	if err != nil {
		fmt.Println(err)
	}
}

func rm(s string) {
	err := exec.Command("rm", "-rf", s).Run()
	if err != nil {
		fmt.Println(err)
	}
}

func ls() {
	cmd := "ls"
	args := []string{"-a"}

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
