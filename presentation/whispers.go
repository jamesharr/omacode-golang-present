package main

import "fmt"
import "time"

func whisperer(input chan int) chan int {
	output := make(chan int)
	go func() {
		for item := range input { // HL
			item += 1
			output <- item
		}
	}()
	return output
}

func main() {
	tStart := time.Now()

	// Construct 100k whisperers // HL
	first := make(chan int)
	last := first
	for i := 0; i < 100000; i++ {
		last = whisperer(last)
	}

	// Start the conversation // HL
	tSend := time.Now()
	first <- 1

	// Read the find // HL
	msg := <-last
	tRecv := time.Now()
	fmt.Println(msg)

	fmt.Println("Setup:", tSend.Sub(tStart))
	fmt.Println("Comms:", tRecv.Sub(tSend))
}
