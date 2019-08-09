package version

import (
	"fmt"
)

var (
	version = "master"
	commit  = "unknown"
)

// BuildString returns full version information
func BuildString() string {
	return fmt.Sprintf("%s (from commit %s)", version, commit)
}

// Version return current version value
func Version() string {
	return version
}

// Commit return current commit value
func Commit() string {
	return commit
}