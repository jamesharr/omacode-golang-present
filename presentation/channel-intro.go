package main

func main() {

	// BEGIN OMIT
	// Create // HL
	var c chan int
	c = make(chan int)
	// or
	c := make(chan int)

	// Send // HL
	c <- 5

	// Receive // HL
	var value int
	value = <-c
	// or
	value := <-c

	// Close // HL
	close(c)
	// END OMIT
} // OMIT
