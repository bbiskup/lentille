package fragments

type BaseFragment struct {
	subFragments []Fragment
}

func (b *BaseFragment) SubFragments() []Fragment {
	return b.subFragments
}
