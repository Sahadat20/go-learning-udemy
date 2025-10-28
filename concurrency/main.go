package main

import (
	"fmt"
	"time"
)

func great(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true

}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)
	go great("nice to meet you!", done)
	go great("How are yu?", done)

	go slowGreet("How ... are ... you?", done)

	go great("I hope you're doing linking the course!", done)

	// <-done
	// <-done
	// <-done
	// <-done
	for range done {
		// fmt.Println(doneChan)
	}
}
