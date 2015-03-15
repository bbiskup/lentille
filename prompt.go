package main

import (
	frag "lentille/fragments"
	"lentille/parser"
	"strings"
)

const joinChar = ":"

// Command line prompt
type Prompt struct {
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

func NewPrompt(configFileName string) (prompt *Prompt, err error) {
	parser := parser.YAMLParser{}
	parsed, err := parser.Parse(configFileName)
	if err != nil {
		return nil, err
	}
	result := Prompt{
		fragments: parsed,
	}
	return &result, nil
}
