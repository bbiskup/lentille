package parser

import (
	"fmt"
	frag "lentille/fragments"
	"log"
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

	ESCAPE_CHAR = '\\'
)

type StateBase int

// Parser state
const (
	STATE_LITERAL StateBase = iota
	STATE_CMD               = iota
)

// parses a prompt specification into a list of fragments
func (p *PromptParser) parse(promptSpec string) []frag.Fragment {
	var result []frag.Fragment
	cmdState := STATE_LITERAL
	escape := true
	var currentCmd string
	var currentLiteral string
	for c := range promptSpec {
		if c == ESCAPE_CHAR {
			escape = true
			continue
		}

		if c == CMD_START && escape {
			result = append(result, frag.NewLiteralFragmentFromRune(CMD_START))
			continue
		}

		if c == CMD_START {
			cmdState = STATE_CMD
			currentCmd = ""
			if len(currentLiteral) > 0 {
				result = append(result, frag.NewLiteralFragmentFromRune(CMD_START))
			}
			continue
		} else if c == CMD_END {
			cmdState = STATE_LITERAL
			log.Printf("Command finished: %s", currentCmd)
			log.Print("TODO create appropriate Fragment instance")
			continue
		}

		if cmdState == STATE_LITERAL {
			currentLiteral += string(c)
		}

		fmt.Println(c)
	}
	return result
}

func NewPromptParser(promptSpec string) *PromptParser {
	return &PromptParser{PromptSpec: promptSpec}
}
