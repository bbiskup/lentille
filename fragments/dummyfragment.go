package fragments

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

type DummyFragment struct {
	baseFragment
}

func (f *DummyFragment) Render() string {
	return "dummy"
}

func (f *DummyFragment) IsActive() bool {
	return true
}

func NewDummyFragment(args *yaml.File) *DummyFragment {
	return &DummyFragment{baseFragment{args: args}}
}
