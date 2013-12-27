package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) chan string {
	ch := make(chan string)
	go func() { // Start a goroutine inside the func before return // HL
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return ch // HL
}

func main() {
	ch := boring("boring!") // HL
	for i := 0; i < 5; i++ {
		fmt.Printf("Boring thing says: %q\n", <-ch)
	}
	fmt.Println("Bye")
}
