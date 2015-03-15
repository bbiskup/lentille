package fragments

import (
	"fmt"
)

type FontWeight int

const (
	NORMAL FontWeight = iota
	BOLD
)

type Color int

const ANSI_COLOR_BASE = 30

const (
	BLACK = iota + ANSI_COLOR_BASE
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

func Bold(msg string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", msg)
}

func Colorize(msg string, color Color) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, msg)
}
