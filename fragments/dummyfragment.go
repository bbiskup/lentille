package fragments

type DummyFragment struct {
	baseFragment
}

func (f *DummyFragment) Render() string {
	return "dummy"
}

func (f *DummyFragment) IsActive() bool {
	return true
}

func NewDummyFragment(conf ConfDict) *DummyFragment {
	return &DummyFragment{baseFragment{conf: conf}}
}
