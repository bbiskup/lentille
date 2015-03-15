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
	result := f.Text

	log.Printf("conf: %#v", f.conf)
	color, hasColor := f.conf["color"]
	if hasColor {
		log.Printf("has color " + color)
		result = Colorize(result, colorMap[color])
	}
	return result
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
		baseFragment: baseFragment{conf: conf},
		Text:         string(text),
	}
}
