package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	go boring("boring!") // HL

	time.Sleep(4 * time.Second)
	fmt.Println("Bye")
}

// END OMIT

func boring(msg string) {
	rand.Seed(time.Now().UnixNano()) // OMIT
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
