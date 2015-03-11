package fragments

import "strings"

// Command line prompt
type Prompt struct {
	fragments []Fragment
}

func (p *Prompt) Render() string {
	var result []string
	for _, fragment := range p.fragments {
		result = append(result, fragment.Render())
	}
	return strings.Join(result, "")
}

func (p *Prompt) Add(fragment Fragment) {
	p.fragments = append(p.fragments, fragment)
}
