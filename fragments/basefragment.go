package fragments

// Base for all fragments

type baseFragment struct {
	conf         ConfDict
	subFragments []Fragment
}

func (b *baseFragment) SubFragments() []Fragment {
	return b.subFragments
}

func newBaseFragment(conf ConfDict) *baseFragment {
	return &baseFragment{conf: conf}
}
