package main

import (
	"fmt"
	"time"
)

func main() {
	fps := getFPS()

	for i := 0; i < 50000; i++ {
		delay(50)
		if isReady, frames := fps(); isReady {
			fmt.Println(isReady, frames)
		}
	}
}

func delay(n int) {
	for j := 0; j < n*1000; j++ {
		fmt.Print("")
	}
}

// GetTime returns current time in nano seconds
func getTime() int64 {
	return time.Now().UTC().UnixNano()
}

type fnBoolInt func() (bool, int)

// GetFPS sdf
func getFPS() fnBoolInt {
	isReady := false
	frames := 0
	var prev, curr int64
	return func() (bool, int) {
		if isReady {
			frames = 0
			prev = getTime()
			isReady = false
		}

		curr = getTime()

		if curr-prev >= 1000000000 {
			isReady = true
		}
		frames++
		return isReady, frames
	}
}
