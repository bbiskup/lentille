package fragments

import (
	"fmt"
	"log"
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

var colorMap map[string]Color = map[string]Color{
	"black":   BLACK,
	"red":     RED,
	"green":   GREEN,
	"yellow":  YELLOW,
	"blue":    BLUE,
	"magenta": MAGENTA,
	"cyan":    CYAN,
	"white":   WHITE,
}

func Bold(msg string) string {
	return fmt.Sprintf("\033[9m%s\033[0m", msg)
}

func Colorize(msg string, color Color) string {
	log.Printf("Colorizing: %#v", color)
	if color == 0 {
		log.Fatal("Not a valid color")
	}
	return fmt.Sprintf("\033[%dm%s\033[0m", color, msg)
}
