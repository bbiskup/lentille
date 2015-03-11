package parser

import (
	"fmt"
)

type PromptParser struct {
	PromptSpec string
}

const (
	// Opening bracket for command
	CMD_START = '{'

	// Closing bracket for command
	CMD_END = '}'

	// Separator between command and subcommand,
	CMD_SUBCMD_SEP = ':'
)

// parses a prompt specification
func (p *PromptParser) parse(promptSpec string) {
	for c := range promptSpec {
		fmt.Println(c)
	}
}

func NewPromptParser(promptSpec string) *PromptParser {
	return &PromptParser{PromptSpec: promptSpec}
}
