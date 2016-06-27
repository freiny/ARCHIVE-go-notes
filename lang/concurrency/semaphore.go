package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s := "abcdefgh"

	tasks := make([]func(), 0, len(s))

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		tasks = append(tasks, func() { disp(c) })
	}

	// no more than 3 tasks running at same time
	maxTasks := make(chan int, 3)

	for _, task := range tasks {
		go func(task func()) {
			maxTasks <- 1
			task()
			<-maxTasks
		}(task)
	}

	waitForUserInput()

	// output shows that no more than 3 tasks are running at any given time

	// POSSIBLE OUTPUT (unpredictable):
	// Press Enter
	// [d[f[gf][ed][be][ag][cb][ha]h]c]

	// POSSIBLE OUTPUT (unpredictable):
	// Press Enter
	// [b[a[ea][cb][dc][fe][gd][hf]h]g]
}

func disp(s string) {
	fmt.Print("[" + s)
	time.Sleep(time.Duration(rand.Int63n(64)) * time.Millisecond)
	fmt.Print(s + "]")
}

func waitForUserInput() {
	fmt.Println("Press Enter")
	var input string
	fmt.Scanln(&input)
}
