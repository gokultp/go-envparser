package commands

import (
	"github.com/gokultp/envparser/internal/version"
)

// Version is the Version command
type Version struct {
}

// NewVersion returns a new instalce of Version
func NewVersion() *Version {
	return &Version{}
}

//InitFlags initialises the flags if any
func (Version) InitFlags() {}

// ParseFlags will parse given flags
func (Version) ParseFlags(args []string) {}

// Help prints Version text for the command, not needed here
func (Version) Help() {}

// Exec will print the build version details
func (Version) Exec() error {
	version.PrintVersion()
	return nil
}
