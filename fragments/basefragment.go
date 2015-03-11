package fragments

// Base for all fragments

type baseFragment struct {
	subFragments []Fragment
}

func (b *baseFragment) SubFragments() []Fragment {
	return b.subFragments
}
