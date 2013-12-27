package main

import (
	"fmt"
	"github.com/jamesharr/omacode-golang-present/lockmanager/lockmanager"
	"math/rand"
)

var lm lockmanager.LockManager

const NUM_ROUTINES = 1000
const NUM_ITERATIONS = 100000 / NUM_ROUTINES
const NUM_RESOURCES = NUM_ROUTINES / 2

func main() {
	lm = lockmanager.Create()

	// Kick off a start by closing this channel
	start := make(chan bool)

	// All simulators will send a bool here when they're done (to make sure they all return)
	end := make(chan bool)

	// Make a list of fake resource names to lock
	resources := make([]string, NUM_RESOURCES)
	for i := 0; i < NUM_RESOURCES; i++ {
		resources[i] = fmt.Sprintf("res.%d", i)
	}

	// Prep routines
	for i := 0; i < NUM_ROUTINES; i++ {
		go simulator(i, resources, start, end)
	}

	// GO!
	close(start)

	// Wait for them to end.
	for i := 0; i < NUM_ROUTINES; i++ {
		<-end
	}
}

func simulator(proc_num int, resources []string, start chan bool, done chan bool) {
	<-start
	for i := 0; i < NUM_ITERATIONS; i++ {
		resIdx := rand.Intn(len(resources))
		lock := lm.Lock(resources[resIdx])
		fmt.Println("Process", proc_num, "got lock for resource", resIdx)
		lock.Release()
		fmt.Println("Process", proc_num, "released lock for resource", resIdx)
	}
	fmt.Println("process", proc_num, "is done")
	done <- true
}
