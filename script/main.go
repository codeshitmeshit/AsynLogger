package main

import (
	"fmt"
	"runtime"
)

func logCallerInfo() {
	// Skip 2 levels to get the caller of the function that called logCallerInfo
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Could not retrieve caller info")
		return
	}
	fmt.Printf("Called from file: %s, line: %d\n", file, line)
}

func someFunction() {
	logCallerInfo()
}

func main() {
	someFunction()
}
