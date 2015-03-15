package fragments

type ConfDict map[string]string

// Part of command line prompt
type Fragment interface {
	// Creates string for rendering prompt
	Render() string

	// Checks whether fragment is currently active.
	// e.g., a fragment showing the current git branch
	// might check whether the current directory is part of
	// a git workspace
	IsActive() bool

	SubFragments() []Fragment
}
