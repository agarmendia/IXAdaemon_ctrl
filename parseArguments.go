package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArguments() (string, string, []string) {

	var mainPort string
	flag.StringVar(&mainPort, "mainPort", "2101", "aplication port")

	var ctrlPort string
	flag.StringVar(&ctrlPort, "ctrlPort", "2102", "control port")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: IXAdaemon_ctrl (--mainPort= 2101) (--ctrlPort = 2102) \"command\"")
		os.Exit(1)
	}
	ports := make([]string, 2)
	ports[0] = "--mainPort=" + mainPort
	ports[1] = "--ctrlPort=" + ctrlPort
	fmt.Println(ports)
	command := "IXAdaemon_server"
	args = append(ports, args[0:]...)

	return ctrlPort, command, args

}
