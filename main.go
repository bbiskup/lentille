package main

import "fmt"
import "strings"

// Command line prompt
type Prompt struct {
	fragments []Fragment
}

func (p *Prompt) Render() string {
	var result []string
	for _, fragment := range p.fragments {
		result = append(result, fragment.Render())
	}
	return strings.Join(result, "")
}

func (p *Prompt) Add(fragment Fragment) {
	p.fragments = append(p.fragments, fragment)
}

type BaseFragment struct {
	subFragments []Fragment
}

func (b *BaseFragment) SubFragments() []Fragment {
	return b.subFragments
}

// Part of command line prompt
type Fragment interface {
	// Creates string for rendering prompt
	Render() string

	// Checks whether fragment is currently active.
	// e.g., a fragment showing the current git branch
	// might check whether the current directory is part of
	// a git workspace
	IsActive() bool

	SubFragments() []Fragment
}

type DummyFragment struct {
	BaseFragment
}

func (f *DummyFragment) Render() string {
	return "dummy"
}

func (f *DummyFragment) IsActive() bool {
	return true
}

func main() {
	fmt.Println("lentille")

	p := new(Prompt)
	p.Add(new(DummyFragment))
	fmt.Println(p.Render())
	fmt.Println("End")
}
