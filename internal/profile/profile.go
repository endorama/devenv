package profile

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
	
	plugins "github.com/endorama/devenv/internal/plugins"
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

	// Config contains profile specific configurations
	Config ProfileConfig

	// RunLoaderPath is the path to be used for the profile run script
	RunLoaderPath string
	// ShellLoaderPath is the path to be used for the profile load script
	ShellLoaderPath string
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

func (p *Profile) LoadConfig() error {
	configFile := filepath.Join(p.Location, "config.yaml")
	err := p.Config.LoadFromFile(configFile)
	if err != nil {
		return errors.Wrap(err, "cannot load config from profile configuration file")
	}
	return nil
}

// LoadPlugins load profile plugins
func (p Profile) LoadPlugins() {
	// We use a switch case to instantiate "statically" plugins to allow
	// enabling them.
	// This requires adding plugin here to allow devenv to be able to use them.
	// It's a compromise to avoid reflection, which at this moment I feel would
	// overengineer this part.
	for _, pluginName := range p.Config.Plugins {
		switch pluginName {
		case plugins.BinPluginName:
			plugin := plugins.NewBinPlugin()
			p.enablePlugin(plugin)
		case plugins.EmailPluginName:
			plugin := plugins.NewEmailPlugin()
			p.enablePlugin(plugin)
		case plugins.EnvsPluginName:
			plugin := plugins.NewEnvsPlugin()
			p.enablePlugin(plugin)
		default:
			continue
		}
	}
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
	err = persistFile(p.ShellLoaderPath, sb.String())
	if err != nil {
		return errors.Wrap(err, "cannot save shell loader")
	}
	ui.Info("Making shell load file executable")
	os.Chmod(p.ShellLoaderPath, 0700)

	return nil
}

// SetupPlugins run Setuppable plugins Setup
func (p Profile) SetupPlugins(ctx context.Context) error {
	ui := ctx.Value("ui").(*cli.BasicUi)
	errorOccurred := false

	for _, plugin := range p.Plugins {
		if setuppablePlugin, ok := plugin.(Setuppable); ok {
			ui.Info(fmt.Sprintf("perform setup: %s", plugin.Name()))
			err := setuppablePlugin.Setup(p.Location)
			if err != nil {
				ui.Error(err.Error())
				errorOccurred = true
			}
		}
	}
	if errorOccurred {
		return errors.New("plugin setup failed")
	}
	return nil
}

// LoadPluginConfigurations load each Configurable plugin configuration
func (p *Profile) LoadPluginConfigurations(ctx context.Context) error {
	ui := ctx.Value("ui").(*cli.BasicUi)
	errorOccurred := false

	for _, plugin := range p.Plugins {
		if configurablePlugin, ok := plugin.(Configurable); ok {
			ui.Info(fmt.Sprintf("configuring: %s", plugin.Name()))
			err := configurablePlugin.LoadConfig(p.Location)
			if err != nil {
				ui.Error(err.Error())
				errorOccurred = true
			}
			ui.Info(fmt.Sprintf("%+v\n", configurablePlugin.Config()))
		}
	}
	if errorOccurred {
		return errors.New("error occurred loading plugin configuration")
	}
	return nil
}

// RunPluginGeneration run each Generator plugin generation
func (p *Profile) RunPluginGeneration(ctx context.Context) error {
	ui := ctx.Value("ui").(*cli.BasicUi)
	errorOccurred := false

	for _, plugin := range p.Plugins {
		if generatorPlugin, ok := plugin.(Generator); ok {
			ui.Info(fmt.Sprintf("perform generation: %s", plugin.Name()))
			err := generatorPlugin.Generate(p.Location)
			if err != nil {
				ui.Error(err.Error())
				errorOccurred = true
			}
		}
	}
	if errorOccurred {
		return errors.New("error occurred running plugin generation")
	}
	return nil
}
