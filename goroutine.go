package tool

import (
	"fmt"
	"runtime"
)

// Go goroutine with recover
func Go(cb func()) {
	go func() {
		defer handlePanic()
		cb()
	}()
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Printf("Recovered in HandlePanic: %v \n", r)

		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		fmt.Printf("goroutine panic stack: %s \n", string(buf[:n]))
	}
}
