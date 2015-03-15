package fragments

import (
	"log"
	"strconv"
)

// Base for all fragments

type baseFragment struct {
	conf         ConfDict
	subFragments []Fragment
	isActive     bool
}

func (b *baseFragment) SubFragments() []Fragment {
	return b.subFragments
}

func (b *baseFragment) IsActive() bool {
	return b.isActive
}

func newBaseFragment(conf ConfDict) *baseFragment {
	isActiveStr := GetOptionalParam(conf, "active", "false")
	isActive, err := strconv.ParseBool(isActiveStr)
	if err != nil {
		log.Fatalf("incorrect bool")
	}
	return &baseFragment{conf: conf, isActive: isActive}
}
