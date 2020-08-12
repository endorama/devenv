package profile

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pkg/errors"

	herectx "github.com/endorama/devenv/internal/context"
	plugins "github.com/endorama/devenv/internal/plugins"
	utils "github.com/endorama/devenv/internal/utils"
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
	ok, err := utils.Exists(p.Location)
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
		case plugins.AwsPluginName:
			plugin := plugins.NewAwsPlugin()
			p.enablePlugin(plugin)
		case plugins.BinPluginName:
			plugin := plugins.NewBinPlugin()
			p.enablePlugin(plugin)
		case plugins.EmailPluginName:
			plugin := plugins.NewEmailPlugin()
			p.enablePlugin(plugin)
		case plugins.EnvsPluginName:
			plugin := plugins.NewEnvsPlugin()
			p.enablePlugin(plugin)
		case plugins.ShellHistoryPluginName:
			plugin := plugins.NewShellHistoryPlugin()
			p.enablePlugin(plugin)
		case plugins.SSHPluginName:
			plugin := plugins.NewSSHPlugin()
			p.enablePlugin(plugin)
		case plugins.TmuxPluginName:
			plugin := plugins.NewTmuxPlugin()
			p.enablePlugin(plugin)
		default:
			panic(fmt.Sprintf("trying to load unknown plugin: %s", pluginName))
		}
	}
}

func (p *Profile) enablePlugin(plugin Pluggable) {
	if p.Plugins[plugin.Name()] == nil {
		p.Plugins[plugin.Name()] = plugin
	}
}

func (p *Profile) disablePlugin(plugin Pluggable) {
	if p.Plugins[plugin.Name()] != nil {
		p.Plugins[plugin.Name()] = nil
	}
}

// GenerateShellLoadFile generate profile shell loader file
func (p Profile) GenerateShellLoadFile(ctx context.Context) error {
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf(`#!/bin/bash
#
# This file has been automatically generated with devenv
# Please remember that running 'devenv rehash' will overwrite this file :)

if [ "$DEVENV_ACTIVE_PROFILE" != "" ]; then
	echo "A profile is already loaded ($DEVENV_ACTIVE_PROFILE). Loading a new profile on top may leak credentials."
	echo "Are you sure you want to continue?"
	select yn in "Yes" "No"; do
		case $yn in
			Yes ) break;;
			No ) exit;;
		esac
	done
fi

if [ "$DEVENV_ACTIVE_PROFILE" == "%s" ]; then
	echo "This profile is already loaded, stopping."
	exit 0
fi

export DEVENV_ACTIVE_PROFILE='%s'
export DEVENV_ACTIVE_PROFILE_PATH='%s'
`, p.Name, p.Name, p.Location))

	sb.WriteString("# plugins BEGIN ##################\n")
	for _, plugin := range p.Plugins {
		ui.Info(fmt.Sprintf("rendering plugin: %s", plugin.Name()))
		sb.WriteString(fmt.Sprintf("# plugin %s\n", plugin.Name()))
		sb.WriteString(plugin.Render(p.Name, p.Location))
	}
	sb.WriteString("# plugins END ####################\n")

	sb.WriteString(fmt.Sprintf("\nexec %s -l\n", p.Shell))

	ui.Info("Save shell load file")
	err := utils.PersistFile(p.ShellLoaderPath, sb.String())
	if err != nil {
		return errors.Wrap(err, "cannot save shell loader")
	}

	ui.Info("Making shell load file executable")
	os.Chmod(p.ShellLoaderPath, 0700)

	return nil
}

// GenerateRunFile generate profile run file
func (p Profile) GenerateRunFile(ctx context.Context) error {
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf(`#!/usr/bin/env bash
#
# This file has been automatically generated with devenv
# Please remember that running 'devenv rehash' will overwrite this file :)

if [ "$DEVENV_ACTIVE_PROFILE" != "" ]; then
	echo "A profile is already loaded ($DEVENV_ACTIVE_PROFILE). Loading a new profile on top may leak credentials."
	echo "Are you sure you want to continue?"
	select yn in "Yes" "No"; do
		case $yn in
			Yes ) break;;
			No ) exit;;
		esac
	done
fi

if [ "$DEVENV_ACTIVE_PROFILE" == "%s" ]; then
	echo "This profile is already loaded, stopping."
	exit 0
fi

export DEVENV_ACTIVE_PROFILE='%s'
export DEVENV_ACTIVE_PROFILE_PATH='%s'
`, p.Name, p.Name, p.Location))

	sb.WriteString("# plugins BEGIN ##################\n")
	for _, plugin := range p.Plugins {
		ui.Info(fmt.Sprintf("rendering plugin: %s", plugin.Name()))
		sb.WriteString(fmt.Sprintf("# plugin %s\n", plugin.Name()))
		sb.WriteString(plugin.Render(p.Name, p.Location))
	}
	sb.WriteString("# plugins END ####################\n")

	sb.WriteString("\neval \"$*\"\n")

	ui.Info("Save run load file")
	err := utils.PersistFile(p.RunLoaderPath, sb.String())
	if err != nil {
		return errors.Wrap(err, "cannot save run loader")
	}

	ui.Info("Making run load file executable")
	os.Chmod(p.RunLoaderPath, 0700)

	return nil
}

// SetupPlugins run Setuppable plugins Setup
func (p Profile) SetupPlugins(ctx context.Context) error {
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)
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
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)
	errorOccurred := false

	for _, plugin := range p.Plugins {
		if configurablePlugin, ok := plugin.(Configurable); ok {
			ui.Info(fmt.Sprintf("configuring: %s", plugin.Name()))
			err := configurablePlugin.LoadConfig(p.Location)
			if err != nil {
				ui.Error(err.Error())
				errorOccurred = true
			}
			ui.Info(fmt.Sprintf("  %+v\n", configurablePlugin.Config()))
		}
	}
	if errorOccurred {
		return errors.New("error occurred loading plugin configuration")
	}
	return nil
}

// RunPluginGeneration run each Generator plugin generation
func (p *Profile) RunPluginGeneration(ctx context.Context) error {
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)
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
