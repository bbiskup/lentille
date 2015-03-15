package parser_test

import (
	"fmt"
	"lentille/parser"
	"testing"
)

func TestParse1(t *testing.T) {
	parser := parser.YAMLParser{}
	p, err := parser.Parse("../testdata/file1.yaml")
	if err != nil {
		t.Errorf("Parse failed: %s", err)
	}

	fmt.Printf("p: %#v", p)

	if len(p) != 5 {
		t.Errorf("Length of parsed data incorrect")
	}
}
