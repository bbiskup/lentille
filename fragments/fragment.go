package fragments

type ConfDict map[string]string

// Part of command line prompt
type Fragment interface {
	// Creates string for rendering prompt
	Render() string

	SubFragments() []Fragment
}
