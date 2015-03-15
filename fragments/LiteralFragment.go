package fragments

import (
	"log"
)

// A fragment that consists of one or more literal characters

type LiteralFragment struct {
	baseFragment
	Text string
}

func (f *LiteralFragment) Render() string {
	return Bold(Colorize(f.Text, MAGENTA)) + Colorize("nix", BLUE)
}

func (f *LiteralFragment) IsActive() bool {
	return true
}

func NewLiteralFragment(conf ConfDict) Fragment {
	text, ok := conf["text"]
	if !ok {
		log.Fatalf("Missing mandatory parameter '%s'", text)
	}

	return &LiteralFragment{
		Text: string(text),
	}
}
