package main

func main() { // OMIT
	select {
	case val := <-ch1:
		fmt.Println("From ch1", val)
	case val := <-ch2:
		fmt.Println("From ch2", val)
	case ch3 <- 5:
		fmt.Println("Sent 5 to ch3")
	case <-time.After(3*time.Second): // HL
		fmt.Println("(╯°□°）╯   Y NO 1 TALK")
	}
} // OMIT

// package time
func After(d time.Duration) chan bool {
	ch := make(chan bool, 1) // buffered channel of size 1 // HL
	go func(){
		time.Sleep(d)
		ch <- true
	}()
	return ch
}
