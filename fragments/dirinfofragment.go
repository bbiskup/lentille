package fragments

// Information about working directory

import (
	"os"
)

type DirInfoFragment struct {
	BaseFragment
}

func (f *DirInfoFragment) Render() string {
	result, err := os.Getwd()
	if err != nil {
		return "Error getting directory"
	} else {
		return result
	}
}

func (f *DirInfoFragment) IsActive() bool {
	return true
}
