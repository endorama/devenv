package command

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Whoami is a command to know which profile is currently loaded
type Whoami struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Whoami) Synopsis() string {
	return "Rebuild profiles loader and shims"
}

// Help return command help text
func (cmd Whoami) Help() string {
	return fmt.Sprintf(`%s
Usage:
  devenv whoami
		Print current loaded profile, if any
	
  devenv whoami -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Whoami) Run(args []string) int {
	currentProfile := os.Getenv("DEVENV_ACTIVE_PROFILE")

	if currentProfile == "" {
		// there is no loaded profile
		return 128
	}

	cmd.UI.Output(currentProfile)
	return 0
}
