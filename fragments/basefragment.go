package fragments

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"strconv"
)

// Base for all fragments

type baseFragment struct {
	args         *yaml.File
	subFragments []Fragment
	isActive     bool
}

func (b *baseFragment) SubFragments() []Fragment {
	return b.subFragments
}

func (b *baseFragment) IsActive() bool {
	return b.isActive
}

func newBaseFragment(args *yaml.File) *baseFragment {
	isActiveStr := GetOptionalParam(args, "active", "false")
	isActive, err := strconv.ParseBool(isActiveStr)
	if err != nil {
		log.Fatalf("incorrect bool")
	}
	return &baseFragment{args: args, isActive: isActive}
}
