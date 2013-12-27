package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Boring struct {
	Comm chan string
	Quit chan bool
}

func boring(msg string) Boring { // HL
	b := Boring{}
	b.Comm = make(chan string)
	b.Quit = make(chan bool)
	go func() {
		for i := 0; ; i++ {
			select {
			case b.Comm <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-b.Quit: // HL
				return // HL
			}
		}
	}()
	return b // HL
}

func fanIn(a, b Boring) Boring { // HL
	combined := Boring{
		Comm: make(chan string),
		Quit: make(chan bool),
	}
	go func() {
		for {
			select {
			case tmp := <-a.Comm:
				combined.Comm <- tmp
			case tmp := <-b.Comm:
				combined.Comm <- tmp
			case <- combined.Quit: // HL
				a.Quit <- true // HL
				b.Quit <- true // HL
				return
			}
		}
	}()
	return combined
}

func main() {
	rand.Seed(time.Now().UnixNano()) // OMIT
	people := fanIn(boring("Jay"), boring("Kevin"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-people.Comm) // HL
	}
	people.Quit <- true // HL
}
