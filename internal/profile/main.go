package profile

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	configs "github.com/endorama/devenv/internal/configs"
)

// New creates a new Profile instance with appropriate configurations
func New(ctx context.Context, name string) (p *Profile, err error) {
	p = &Profile{Name: name, Config: ProfileConfig{}}
	loc, err := getProfileLocation(*p)
	if err != nil {
		return p, errors.Wrap(err, "cannot get profile location")
	}
	p.Location = loc

	if !p.Exists() {
		return p, errors.New("profile does not exists")
	}

	err = p.LoadConfig()
	if err != nil {
		return p, fmt.Errorf("cannot load profile configurations: %w", err)
	}

	p.Plugins = make(map[string]Pluggable)
	p.Shell = os.Getenv("SHELL")

	p.RunLoaderPath = filepath.Join(p.Location, configs.ShellRunnerFilename)
	p.ShellLoaderPath = filepath.Join(p.Location, configs.ShellLoaderFilename)

	return p, err
}
