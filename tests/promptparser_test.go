package main_test

import (
	par "lentille/parser"
	"testing"
)

func TestEmpty(t *testing.T) {
	p := par.NewPromptParser("")
	if p.PromptSpec != "" {
		t.Fatalf("prompt spec not empty")
	}
}

func TestTokenization(t *testing.T) {

}
