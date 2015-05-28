package main

import (
	"net"
	"os"
)

//Asks the native programs state to the server, and returns:
//0 if OK
//1 if waiting
//2 if there is not a instance of the native program
//3 there is no server running

func main() {

	var state []byte
	state = []byte("1")

	port := parseArguments()

	service := ":" + port
	conn, err := net.Dial("tcp", service)
	if err != nil {
		os.Exit(3)
	}

	conn.Write([]byte("state"))
	conn.Read(state)

	if string(state) == "1" {
		os.Exit(2)
	} else if string(state) == "2" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
