package profile

import (
	"context"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// New creates a new Profile instance with appropriate configurations
func New(ctx context.Context, name string) (p *Profile, err error) {
	p = &Profile{Name: name}
	loc, err := getProfileLocation(*p)
	if err != nil {
		return p, errors.Wrap(err, "cannot get profile location")
	}
	p.Location = loc

	if !p.Exists() {
		return p, errors.New("profile does not exists")
	}

	p.Plugins = make(map[string]Pluggable)
	p.Shell = os.Getenv("SHELL")

	p.runLoaderPath = filepath.Join(p.Location, shellRunnerFilename)
	p.shellLoaderPath = filepath.Join(p.Location, shellLoaderFilename)

	return p, err
}
