package main

import (
	"flag"
	"fmt"
	"os"
)

var defaultConfFileName = os.Getenv("HOME") + string(os.PathSeparator) + ".lentille.yml"

func main() {
	fmt.Println("lentille")
	configFileName := flag.String("config", defaultConfFileName, "Configuration file name")
	flag.Parse()

	prompt, err := NewPrompt(*configFileName)

	if err != nil {
		fmt.Printf("Error while parsing %s: %s", configFileName, err)
		os.Exit(0)
	}

	fmt.Printf("prompt: %#v\n", prompt)
	fmt.Printf("\n%s\n", prompt.Render())
}
