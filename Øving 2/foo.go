// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

// Control signals

func number_server(add_number <-chan int, number chan<- int, finished chan<- bool) {
	var i = 0
	var nmb_finished = 0
	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {

		case a := <-add_number:
			if a == -1 {
				i--
			}
			if a == 1 {
				i++
			}

		case number <- i:

			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.

		case t:= <-finished:
			nmb_finished++
			if (nmb_finished==2){}
				break
			}

		}

	}
}

func incrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- 1
	}

	finished <- true

	//TODO: signal that the goroutine is finished
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- -1
	}
	finished <- true
	//TODO: signal that the goroutine is finished
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels

	ch_change_i := make(chan int)

	ch_i_value := make(chan int)

	ch_finished := make(chan bool)
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.

	// TODO: Spawn the required goroutines

	// TODO: block on finished from both "worker" goroutines

	go incrementing(ch_change_i, ch_finished)
	go decrementing(ch_change_i, ch_finished)

	number_server(ch_change_i, ch_i_value)
	time.Sleep(1000 * time.Millisecond)
	Println("The magic number is:", <-ch_i_value)
}
