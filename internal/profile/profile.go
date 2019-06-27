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
}

func NewProfile(name string) (p *Profile, err error) {
	p = &Profile{Name: name}
	err = p.GetLocation()
	fmt.Println(fmt.Sprintf("%v", p))

	p.Plugins = []string{"aws", "bin", "email", "envs", "gpg", "ssh", "shellhistory"}
	p.Shell = os.Getenv("SHELL")
	
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
	err = p.GenerateShellLoadFile()
	if err != nil {
		return errors.Wrap(err, "error generating profile shell load file")
	}
	
	ui.Info("Generating run file")
	err = p.GenerateRunFile()
	if err != nil {
		return errors.Wrap(err, "error generating profile run file")
	}

	return nil
}

func (p Profile) GenerateShellLoadFile() error {
	var b strings.Builder
	
	tmpl, err := template.New("header").
		Parse(shellLoaderTemplate)
	if err != nil { panic(err) }
	err = tmpl.Execute(&b, p)
	if err != nil { panic(err) }
	
	fmt.Println(b.String())
	return nil
}

func (p Profile) GenerateRunFile() error {
	var b strings.Builder
	
	tmpl, err := template.New("header").
		Parse(runTemplate)
	if err != nil { panic(err) }
	err = tmpl.Execute(&b, p)
	if err != nil { panic(err) }
	
	fmt.Println(b.String())
	return nil
}