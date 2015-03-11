package main

import (
	"fmt"
	frag "lentille/fragments"
)

func main() {
	fmt.Println("lentille")

	p := new(frag.Prompt)
	p.Add(new(frag.DummyFragment))
	fmt.Println(p.Render())
	fmt.Println("End")
}
