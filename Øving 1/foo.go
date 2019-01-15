// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	var count int

	for count < 1000000 {
		i++
		count++
	}
}

func decrementing() {
	var count int
	for count < 1000000 {
		i--
		count++
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go incrementing()
	go decrementing()

	time.Sleep(1000 * time.Millisecond)

	Println("The magic number is:", i)

}
