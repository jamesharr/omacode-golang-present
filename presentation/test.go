package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 5

	var value int
	value = <-c
	fmt.Println(value)
}
