package command

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/pkg/errors"

	"github.com/mitchellh/cli"

	"github.com/endorama/devenv/internal/profile"
)

type Shell struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Shell) Synopsis() string {
	return "Load a shell preconfigured with profile environment"
}

// Help return command help text
func (cmd Shell) Help() string {
	return fmt.Sprintf(`%s

Usage:
  devenv shell <PROFILE_NAME>
  devenv shell -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Shell) Run(args []string) int {
	// fmt.Println(args)
	// ctx := context.WithValue(context.Background(), "ui", cmd.UI)

	if len(args) != 1 {
		cmd.UI.Error("PROFILE_NAME is required")
		return 1
	}

	profileName := args[0]

	cmd.UI.Info(fmt.Sprintf("creating shell for profile: %s", profileName))

	p, err := profile.New(context.Background(), profileName)
	if err != nil {
		cmd.UI.Error(errors.Wrap(err, "cannot load profile").Error())
		return 1
	}

	cmd.UI.Info(fmt.Sprintf("shell load file is: %s", p.ShellLoaderPath))

	err = syscall.Exec("/usr/bin/env", []string{"bash", p.ShellLoaderPath}, os.Environ())
	if err != nil {
		cmd.UI.Error(errors.Wrap(err, "cannot exec load file").Error())
		return 1
	}

	cmd.UI.Info("Done")
	return 0
}
