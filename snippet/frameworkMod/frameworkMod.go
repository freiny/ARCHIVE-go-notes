package main

import (
	"fmt"
	"strconv"
)

func main() {
	framework := Framework{"WindowFramework", 1280, 720}
	fmt.Println(framework.GetSize(), framework.GetVer())
	// OUTPUT: [1280 720] Framework

	renamed := FrameworkRename{framework}
	fmt.Println(renamed.GetSize(), renamed.GetVer())
	// OUTPUT: [1280 720] Framework

	mod := FrameworkMod{framework}
	fmt.Println(mod.GetSize(), mod.GetVer())
	// OUTPUT: [WIDTH: 1280 HEIGHT: 720] Framework

}

type Framework struct {
	name   string
	width  int
	height int
}

func (f Framework) GetSize() []int {
	return []int{f.width, f.height}
}
func (f Framework) GetVer() string {
	return "Framework"
}

type FrameworkRename struct {
	Framework
}

type FrameworkMod struct {
	Framework
}

func (fm FrameworkMod) GetSize() []string {
	return []string{"WIDTH: " + strconv.Itoa(fm.width), "HEIGHT: " + strconv.Itoa(fm.height)}
}
