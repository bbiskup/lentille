package fragments

import (
	"fmt"
)

type FontAppearance int

const (
	NORMAL      FontAppearance = 0
	BOLD        FontAppearance = 1
	NEGATIVE    FontAppearance = 7
	CROSSED_OUT FontAppearance = 9
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
	return fmt.Sprintf("\033[9m%s\033[0m", msg)
}

func Colorize(msg string, color Color) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, msg)
}
