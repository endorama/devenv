package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"

	"github.com/endorama/devenv/internal/profile"
)

type Rehash struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Rehash) Synopsis() string {
	return "Rebuild profiles loader and shims"
}

// Help return command help text
func (cmd Rehash) Help() string {
	return fmt.Sprintf(`%s
Usage:
  devenv rehash [PROFILE_NAME]
  devenv rehash -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Rehash) Run(args []string) int {
	var err error
	fmt.Println(args)

	ctx := context.WithValue(context.Background(), "ui", cmd.UI)

	switch {
	case len(args) == 0:
		cmd.UI.Info("Rehashing all profiles")
		// err = profile.RehashAllProfiles(context.TODO())
	case len(args) == 1:
		cmd.UI.Info(fmt.Sprintf("Rehashing %s", args[0]))
		err = profile.RehashSingle(ctx, args[0])
	case len(args) > 1:
		cmd.UI.Info(fmt.Sprintf("Rehashing multiple profiles: %s", strings.Join(args, ", ")))
		// err = profile.RehashMUltipleProfiles(context.TODO())
	}

	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	cmd.UI.Info("Done")
	return 0
}
