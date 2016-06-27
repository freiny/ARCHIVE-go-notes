package main

import "fmt"

func main() {

	messages := make(chan string)
	go disp("abcdefghijklmnopqrstuvwxyz", messages, "disp-0-done")
	go disp("ABCDEFGHIJKLMNOPQRSTUVWXYZ", messages, "disp-1-done")
	numTasks := 2

	out := ""
	for i := 0; i < numTasks; i++ {
		out += "\n[" + <-messages + "]"
	}
	fmt.Println(out)

	// POSSIBLE OUTPUT (unpredictable):
	// AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtuUvVwWxXyYzZ
	// [disp-0-done]
	// [disp-1-done]

	// POSSIBLE OUTPUT (unpredictable):
	// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	// [disp-1-done]
	// [disp-0-done]
}

func disp(s string, messages chan string, doneMsg string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	messages <- doneMsg
}
