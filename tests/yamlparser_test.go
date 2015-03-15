package parser_test

import (
	"fmt"
	"lentille/parser"
	"testing"
)

func TestParse1(t *testing.T) {
	p, err := parser.Parse("../testdata/file1.yaml")
	if err != nil {
		t.Errorf("Parse failed: %s", err)
	}
	if len(p) != 4 {
		t.Errorf("Length of parsed data incorrect")
	}
	fmt.Printf("p: %s", p)
}
