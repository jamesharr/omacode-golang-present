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

func boring(msg string) Boring {
	b := Boring{}
	b.Comm = make(chan string)
	b.Quit = make(chan bool)
	go func() {
		for i := 0; ; i++ {
			select {
			case b.Comm <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-b.Quit: // HL
				// do some cleanup, then send the "we're all done" signal // HL
				b.Quit <- true // HL
				return // HL
			}
		}
	}()
	return b
}

func fanIn(a, b Boring) Boring {
	combined := Boring{
		Comm: make(chan string),
		Quit: make(chan bool),
	}
	go func() {
		for {
			// inside fanIn
			select {
			case tmp := <-a.Comm:
				combined.Comm <- tmp
			case tmp := <-b.Comm:
				combined.Comm <- tmp
			case <- combined.Quit:
				a.Quit <- true
				b.Quit <- true
				<-a.Quit // HL
				<-b.Quit // HL
				combined.Quit <- true // HL
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
		fmt.Println(<-people.Comm)
	}
	people.Quit <- true
	<-people.Quit // wait for complete // HL
}
