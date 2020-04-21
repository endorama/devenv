package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"

	herectx "github.com/endorama/devenv/internal/context"
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
  devenv rehash
		Rehash all profiles
  devenv rehash [PROFILE_NAME]
		Rehash a single profile
	
  devenv rehash -h | --help

Options:
  -h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Rehash) Run(args []string) int {
	var err error
	fmt.Println(args)

	ctx := context.WithValue(context.Background(), herectx.UI, cmd.UI)

	switch {
	case len(args) == 0:
		cmd.UI.Info("Rehashing all profiles")
		err = profile.RehashAllProfiles(ctx)
	case len(args) == 1:
		cmd.UI.Info(fmt.Sprintf("Rehashing profile: %s", args[0]))
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
