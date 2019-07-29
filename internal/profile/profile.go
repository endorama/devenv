package profile

import (
	"strings"
	"context"
	
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"

	"github.com/endorama/devenv/internal/profile/template"
)

const (
	profilesDirectory   = "profiles"
	shellLoaderFilename = "load.sh"
	shellRunnerFilename = "run.sh"
)

// Profile holds information for a single profile
type Profile struct {
	// Name of the profile
	Name string
	// Location of the profile
	Location string
	// Plugins contains a map of Pluggable
	Plugins map[string]Pluggable
	// Shell is the shell to be used by profile
	Shell string

	// runLoaderPath is the path to be used for the profile run script
	runLoaderPath string
	// shellLoaderPath is the path to be used for the profile load script
	shellLoaderPath string
}

// Exists return where the profile exists
// Profile existance is determinated by profile Location existance
func (p Profile) Exists() bool {
	if p.Location == "" {
		return false
	}
	ok, err := exists(p.Location)
	if err != nil {
		return false
	}
	return ok
}

// LoadPlugins load profile plugins
func (p Profile) LoadPlugins() {
	// due to golang working, it's far easier to initialize plugins one by one
	// and then enabling then individually than trying to load them
	// dinamically
	// dummyPlugin := NewDummyPlugin()
	// p.enablePlugin(dummyPlugin)
}

func (p *Profile) enablePlugin(plugin Pluggable) {
	if p.Plugins[plugin.Name()] == nil {
		p.Plugins[plugin.Name()] = plugin
	}
}

func (p *Profile) dnablePlugin(plugin Pluggable) {
	if p.Plugins[plugin.Name()] != nil {
		p.Plugins[plugin.Name()] = nil
	}
}

// GenerateShellLoadFile generate profile shell loader file
func (p Profile) GenerateShellLoadFile(ctx context.Context) error {
	ui := ctx.Value("ui").(*cli.BasicUi)

	sb := strings.Builder{}

	ui.Info("Generating shell load file")
	tmpl, err := template.GetShellLoaderTemplate()
	if err != nil {
		return errors.Wrap(err, "cannot parse shell loader template")
	}
	err = tmpl.Execute(&sb, p)
	if err != nil {
		return errors.Wrap(err, "cannot execute shell loader template")
	}
	ui.Info("Save shell load file")
	err = persistFile(p.shellLoaderPath, sb.String())
	if err != nil {
		return errors.Wrap(err, "cannot save shell loader")
	}

	return nil
}