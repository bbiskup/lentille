package main

import (
	"flag"
	"fmt"
	frag "lentille/fragments"
	"os"
)

var defaultConfFileName = os.Getenv("HOME") + string(os.PathSeparator) + ".lentille.yml"

var (
	configFileName = flag.String("configFile", defaultConfFileName, "Configuration file name")
)

func main() {
	fmt.Println("lentille")

	conf := NewConfig("configFileName")

	p := NewPrompt(configFileName)
	p.Add(new(frag.DummyFragment))
	p.Add(new(frag.DirInfoFragment))
	fmt.Println(p.Render())
	fmt.Println("End")
}
