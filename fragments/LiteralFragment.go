package fragments

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"reflect"
)

// A fragment that consists of one or more literal characters

type LiteralFragment struct {
	baseFragment
	Text string
}

func (f *LiteralFragment) Render() string {
	return "dummy"
}

func (f *LiteralFragment) IsActive() bool {
	return true
}

func NewLiteralFragment(args *yaml.File) Fragment {
	log.Printf("NewLiteralFragment: %s", reflect.TypeOf(args))
	text, err := args.Get("text")
	if err != nil {
		log.Fatalf("Missing mandatory parameter '%s'", text)
	}

	return &LiteralFragment{
		baseFragment: baseFragment{args: args},
		Text:         string(text),
	}
}
