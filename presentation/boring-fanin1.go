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

func fanIn(a, b chan string) chan string {
	combined := make(chan string)
	go func() {
		for {
			select { // HL
			case tmp := <-a:
				combined <- tmp
			case tmp := <-b:
				combined <- tmp
			}
		}
	}()
	return combined
}

func main() {
	rand.Seed(time.Now().UnixNano()) // OMIT
	people := fanIn(boring("Jay"), boring("Kevin")) // HL
	for i := 0; i < 10; i++ {
		fmt.Println(<-people)
	}
}
