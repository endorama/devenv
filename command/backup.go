package command

import (
	"context"
	"fmt"
	"strings"

	herectx "github.com/endorama/devenv/internal/context"
	"github.com/endorama/devenv/internal/profile/backup"
	"github.com/mitchellh/cli"
)

// Backup is a command to create an encrypted backup of a profile
type Backup struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Backup) Synopsis() string {
	return "Backup all available profiles"
}

// Help return command help text
func (cmd Backup) Help() string {
	return fmt.Sprintf(`%s
Usage:
	devenv backup
		Backup all profiles
	devenv backup [PROFILE_NAME]
		Backup a single profile

	devenv backup -h | --help

Options:
	-h --help     Show this screen.
`, cmd.Synopsis())
}

// Run run the actual command
func (cmd Backup) Run(args []string) int {
	var err error
	fmt.Println(args)

	ctx := context.WithValue(context.Background(), herectx.UI, cmd.UI)

	switch {
	case len(args) == 0:
		cmd.UI.Info("Creating backup for all profiles")
		// err = profile.RehashAllProfiles(ctx)
	case len(args) == 1:
		cmd.UI.Info(fmt.Sprintf("Creating backup for profile: %s", args[0]))
		err = backup.BackupSingle(ctx, args[0])
	case len(args) > 1:
		cmd.UI.Info(fmt.Sprintf("Creating backup for multiple profiles: %s", strings.Join(args, ", ")))
		// err = profile.RehashMUltipleProfiles(context.TODO())
	}

	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	cmd.UI.Info("Done")
	return 0
}
