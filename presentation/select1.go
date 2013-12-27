package main

func main() { // OMIT
	select { // HL
	case val := <-ch1:
		fmt.Println("From ch1", val)
	case val := <-ch2:
		fmt.Println("From ch2", val)
	case ch3 <- 5:
		fmt.Println("Sent 5 to ch3")
	} // HL
} // OMIT

func withDefault { // OMIT
	select { // OMIT
	default: // HL
		fmt.Println("No one wants to talk.")
	} // OMIT
} // OMIT
