package fragments

// A fragment that consists of one or more literal characters

type LiteralFragment struct {
	baseFragment
	Literal string
}

func (f *LiteralFragment) Render() string {
	return "dummy"
}

func (f *LiteralFragment) IsActive() bool {
	return true
}

func NewLiteralFragment(literal string) Fragment {
	return &LiteralFragment{Literal: literal}
}

func NewLiteralFragmentFromRune(literal rune) Fragment {
	return &LiteralFragment{Literal: string(literal)}
}
