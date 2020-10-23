package command

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/pkg/errors"

	"github.com/mitchellh/cli"

	"github.com/endorama/devenv/internal/profile"
)

// Run is the CLI command struct
type Run struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Run) Synopsis() string {
	return "Run a command within a profile environment"
}

// Help return command help text
func (cmd Run) Help() string {
	return fmt.Sprintf(`%s

Usage:
  devenv run <PROFILE_NAME> <COMMAND_STRING>
  devenv run -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Run) Run(args []string) int {
	// fmt.Println(args)
	// ctx := context.WithValue(context.Background(), "ui", cmd.UI)
	if len(args) < 2 {
		cmd.UI.Error("PROFILE_NAME is required")
		cmd.UI.Error("COMMAND_STRING is required")
		return 1
	}

	profileName := args[0]
	commandString := strings.Join(args[1:], " ")

	cmd.UI.Info(fmt.Sprintf("running in profile: %s", profileName))

	p, err := profile.New(context.Background(), profileName)
	if err != nil {
		cmd.UI.Error(errors.Wrap(err, "cannot load profile").Error())
		return 1
	}

	cmd.UI.Info(fmt.Sprintf("run load file is: %s", p.RunLoaderPath))

	err = syscall.Exec("/usr/bin/env",
		[]string{"bash", p.RunLoaderPath, commandString},
		os.Environ())
	if err != nil {
		cmd.UI.Error(errors.Wrap(err, "cannot exec load file").Error())
		return 1
	}

	cmd.UI.Info("Done")
	return 0
}
