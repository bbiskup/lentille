package fragments

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

type FragmentArgs yaml.Map

// Base for all fragments

type baseFragment struct {
	args         *FragmentArgs
	subFragments []Fragment
}

func (b *baseFragment) SubFragments() []Fragment {
	return b.subFragments
}

func newBaseFragment(args *FragmentArgs) *baseFragment {
	return &baseFragment{args: args}
}
