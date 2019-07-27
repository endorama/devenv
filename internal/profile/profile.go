package profile

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pkg/errors"

	"github.com/endorama/devenv/internal/template"
)

var (
	profilesDirectory string
)

type Profile struct {
	Name     string
	Location string
	Plugins  map[string]Pluggable
	Shell    string

	runLoader   string
	shellLoader string
}

func NewProfile(name string) (p *Profile, err error) {
	p = &Profile{Name: name}
	err = p.GetLocation()
	fmt.Println(fmt.Sprintf("%v", p))

	p.Plugins = map[string]Pluggable{}
	p.Shell = os.Getenv("SHELL")

	p.shellLoader = filepath.Join(p.Location, "load.sh")
	p.runLoader = filepath.Join(p.Location, "run.sh")

	return p, err
}

func (p *Profile) GetLocation() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profilesDirectory = fmt.Sprintf("%s/profiles", userHome)
	guessedLocation := filepath.Join(profilesDirectory, p.Name)
	_, err = os.Stat(guessedLocation)
	if os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("Profile does not exists at %s", guessedLocation))
	} else if err != nil {
		return err
	}
	p.Location = guessedLocation

	return nil
}

func (p Profile) Rehash(ctx context.Context) (err error) {
	ui := ctx.Value("ui").(*cli.BasicUi)

	ui.Info("Generating shell load file")
	file, err := p.GenerateShellLoadFile()
	if err != nil {
		return errors.Wrap(err, "error generating profile shell load file")
	}
	fmt.Println(file.String())
	ui.Info("Save shell load file")
	err = persistFile(p.shellLoader, file.String())
	if err != nil {
		ui.Error(err.Error())
		return err
	}

	ui.Info("Generating run file")
	file, err = p.GenerateRunFile()
	if err != nil {
		return errors.Wrap(err, "error generating profile run file")
	}
	fmt.Println(file.String())
	ui.Info("Save run file")
	err = persistFile(p.runLoader, file.String())
	if err != nil {
		ui.Error(err.Error())
		return err
	}

	return nil
}

func (p Profile) GenerateShellLoadFile() (b strings.Builder, err error) {
	tmpl, err := template.GetShellLoaderTemplate()
	if err != nil {
		return b, errors.Wrap(err, "cannot parse template")
	}
	err = tmpl.Execute(&b, p)
	if err != nil {
		return b, errors.Wrap(err, "cannot execute template")
	}

	return b, nil
}

func (p Profile) GenerateRunFile() (b strings.Builder, err error) {
	tmpl, err := template.GetRunnerTemplate()
	if err != nil {
		return b, errors.Wrap(err, "cannot parse template")
	}
	err = tmpl.Execute(&b, p)
	if err != nil {
		return b, errors.Wrap(err, "cannot execute template")
	}

	return b, nil
}

func (p *Profile) EnablePlugin(plugin Pluggable) {
	if p.Plugins[plugin.Name()] == nil {
		p.Plugins[plugin.Name()] = plugin
	}
}

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
		return fmt.Errorf("error occurred loading plugin configuration")
	}
	return nil
}

func (p *Profile) RunPluginGeneration(ctx context.Context) error {
	ui := ctx.Value("ui").(*cli.BasicUi)
	errorOccurred := false

	for _, plugin := range p.Plugins {
		if generatorPlugin, ok := plugin.(Generator); ok {
			ui.Info(fmt.Sprintf("perform generation: %s", plugin.Name()))
			err := generatorPlugin.Generate(*p)
			if err != nil {
				ui.Error(err.Error())
				errorOccurred = true
			}
		}
	}
	if errorOccurred {
		return fmt.Errorf("error occurred running plugin generation")
	}
	return nil
}

func persistFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}
