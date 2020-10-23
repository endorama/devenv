package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"

	configs "github.com/endorama/devenv/internal/configs"
)

// List is a command to list available profiles
type List struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd List) Synopsis() string {
	return "List all available profiles"
}

// Help return command help text
func (cmd List) Help() string {
	return fmt.Sprintf(`%s
Usage:
  devenv list
  devenv list -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd List) Run(args []string) int {
	var err error

	userHome, err := os.UserHomeDir()
	if err != nil {
		cmd.UI.Error(fmt.Errorf("cannot detect user home folder: %w", err).Error())
		return 1
	}
	profilesLocation := filepath.Join(userHome, configs.ProfilesDirectory)
	cmd.UI.Info(fmt.Sprintf("Listing profiles from %s", profilesLocation))

	files, err := ioutil.ReadDir(profilesLocation)
	if err != nil {
		cmd.UI.Error(fmt.Errorf("cannot read folder content: %w", err).Error())
		return 1
	}

	for _, f := range files {
		matches, err := filepath.Glob(filepath.Join(profilesLocation, f.Name(), "config.yaml"))
		if err != nil {
			cmd.UI.Error(fmt.Errorf("bad glob patter: %w", err).Error())
			return 1
		}

		if len(matches) == 1 {
			cmd.UI.Output(f.Name())
		}
	}

	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	cmd.UI.Info("Done")
	return 0
}
