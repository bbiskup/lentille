package main

import (
	"flag"
	"fmt"
	"lentille/parser"
	"os"
)

var defaultConfFileName = os.Getenv("HOME") + string(os.PathSeparator) + ".lentille.yml"

func main() {
	fmt.Println("lentille")
	configFileName := flag.String("config", defaultConfFileName, "Configuration file name")
	flag.Parse()

	parsed, err := parser.Parse(*configFileName)

	if err != nil {
		fmt.Printf("Error while parsing %s: %s", configFileName, err)
		os.Exit(0)
	}

	fmt.Printf("parsed: %s", parsed)
}
