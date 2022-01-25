package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andey-robins/bhgo/proxy"
	"github.com/andey-robins/bhgo/pscan"
)

func main() {
	// define flags
	tool := flag.String("run", "help", "the tool to run")
	target := flag.String("target", "localhost", "an address must be specified")

	// process flags
	flag.Parse()

	// verify flags
	if tool == nil {
		log.Panic("-run requires a name of one of the tools")
	}

	switch *tool {
	case "pscan":
		scanHandler(target)
	case "proxy":
		proxyHandler(target)
	case "help":
		helpHandler()
	}
}

func scanHandler(address *string) {
	if address == nil {
		log.Fatalln("An address must be specified with -target")
	}

	fmt.Println(pscan.Scan(*address))
}

func proxyHandler(target *string) {
	if target != nil {
		log.Fatalln("An address must be specified with -target. This is the site you can't reach.")
	}

	proxy.Proxy(*target)
}

func helpHandler() {
	fmt.Printf("\n\n")
	fmt.Println("Unable to run the given configuration.\nLook at the available config options below.")
	fmt.Println()
	fmt.Println("Tools: pscan proxy ")
	fmt.Println("Flags: -target ")
	fmt.Println()
}
