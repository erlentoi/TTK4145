package main

import (
	. "fmt"
	"runtime"
)

func number_server(add_number <-chan int, number chan<- int) {
	var i = 0
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
		}

	}
}

func incrementing(add_number chan<- int, finished chan bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- 1
	}
	finished <- true
}

func decrementing(add_number chan<- int, finished chan bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- -1
	}
	finished <- true
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	ch_change_i := make(chan int)
	ch_i_value := make(chan int)
	ch_finished := make(chan bool)

	go incrementing(ch_change_i, ch_finished)
	go decrementing(ch_change_i, ch_finished)

	go number_server(ch_change_i, ch_i_value)

	//<-ch_finished
	//<-ch_finished
	//time.Sleep(time.Second * 3)
	println("asdasd")
	Println("The magic number is:", <-ch_i_value)
}
