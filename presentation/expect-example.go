package main

import "github.com/jamesharr/expect"


func usage() { // OMIT
	exp, err := expect.Spawn( "ssh", "localhost" )
	if err != nil { panic(err) }
	defer exp.Close()

	exp.SetTimeout(5 * time.Second) // HL

	exp.Expect(`[Pp]assword:`) // HL
	exp.Sendln("terriblepassword")

	exp.Expect(`\$`) // HL
	exp.Sendln("ls -lh")
	exp.Expect("ls -lh") // Cut out remote-echo

	m, _ := exp.Expect(`(?m)^.*\$`) // HL
	fmt.Println("Directory Listing:", m.Before)

	exp.Sendln("exit") // HL
	exp.ExpectEOF() // HL
} // OMIT

func problem() { // OMIT
	myPty := pty.Start(exec.Command("ssh", "localhost"))
	buffer := make([]byte, 4096)
	bytesRead, err := myPty.Read(buffer) // HL
	buffer = buffer[0:bytesRead] // Chop buffer slice to only contain what was read
} // OMIT
