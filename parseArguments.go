package main

import (
	"flag"
)

func parseArguments() string {

	var ctrlPort string
	flag.StringVar(&ctrlPort, "ctrlPort", "2102", "control port")

	flag.Parse()

	return ctrlPort

}
