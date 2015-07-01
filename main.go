package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {

	ctrlPort, cmd, args := parseArguments()
	counter := 0

	conn, err := net.Dial("tcp", "127.0.0.1:"+ctrlPort)
	if err != nil {
		if err := exec.Command(cmd, args...).Start(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	time.Sleep(1 * time.Second)
	conn, err = net.Dial("tcp", "127.0.0.1:"+ctrlPort)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn.Write([]byte("state\n"))
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		if message == "1\n" {
			fmt.Print("...")
		}
		if message == "0\n" {
			break
		}
		if message == "3\n" {
			fmt.Print("---")
			counter++
			if counter == 5 {
				break
			}
		}

		time.Sleep(2 * time.Second)

	}
	conn.Close()

	if counter != 5 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}
