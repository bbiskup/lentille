package main_test

import (
	frag "lentille/fragments"
	par "lentille/parser"
	"testing"
)

func TestEmpty(t *testing.T) {
	p := par.NewPromptParser("")
	if p.PromptSpec != "" {
		t.Error("prompt spec not empty")
	}
}

var parseTests = []struct {
	PromptSpec string
	Expected   []frag.Fragment
}{
	{"", []frag.Fragment{}},
	{"xyz", []frag.Fragment{
		frag.NewLiteralFragment("xyz"),
	}},
}

func TestParse(t *testing.T) {
	for _, r := range parseTests {
		p := par.NewPromptParser(r.PromptSpec)
		parseResult := p.Parse()

		for i, _ := range parseResult {
			if parseResult[i] != r.Expected[i] {
				t.Errorf("Expected: %s, got: %s",
					r.Expected, parseResult)
			}
		}

	}
}
