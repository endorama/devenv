package profile

import (
	"text/template"
	"strings"
	"context"
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/pkg/errors"
	"github.com/mitchellh/cli"
)

var (
	profilesDirectory string

	commonTemplate = `
#!/bin/bash
#
# This file has been automatically generated with devenv
# Please remember that running 'devenv rehash' will overwrite this file :)

export DEVENV_ACTIVE_PROFILE='{{.Name}}'
export DEVENV_ACTIVE_PROFILE_PATH='{{.Location}}'

# plugin BEGIN ##################
{{range .Plugins}}
# plugin: {{.}}
__devenv_plugin__{{.}}__generate_loader
{{end}}
# plugin END ####################`
	
	shellLoaderTemplate = fmt.Sprintf(`
%s

exec {{.Shell}} -l
`, commonTemplate)

	runTemplate = fmt.Sprintf(`
%s

eval $@
`, commonTemplate)
)

type Profile struct {
	Name string
	Location string
	Plugins []string
	Shell string

	shellLoader string
	runLoader string
}

func NewProfile(name string) (p *Profile, err error) {
	p = &Profile{Name: name}
	err = p.GetLocation()
	fmt.Println(fmt.Sprintf("%v", p))

	p.Plugins = []string{"aws", "bin", "email", "envs", "gpg", "ssh", "shellhistory"}
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
	_, err = os.Stat(guessedLocation);
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
	tmpl, err := template.New("shell-loader").
		Parse(shellLoaderTemplate)
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
	tmpl, err := template.New("runner").
		Parse(runTemplate)
	if err != nil {
		return b, errors.Wrap(err, "cannot parse template")
	}
	err = tmpl.Execute(&b, p)
	if err != nil {
		return b, errors.Wrap(err, "cannot execute template")
	}
	
	return b, nil
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