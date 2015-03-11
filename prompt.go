package main

import (
	frag "lentille/fragments"
	"strings"
)

const joinChar = ":"

// Command line prompt
type Prompt struct {
	config    Config
	fragments []frag.Fragment
}

func (p *Prompt) Render() string {
	var result []string
	for _, fragment := range p.fragments {
		result = append(result, fragment.Render())
	}
	return strings.Join(result, joinChar)
}

func (p *Prompt) Add(fragment frag.Fragment) {
	p.fragments = append(p.fragments, fragment)
}

func NewPrompt(configFileName string) *Prompt {
	result := Prompt{config: NewConfig(configFileName)}
	return &result
}
