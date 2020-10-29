package command

import (
	"context"
	"fmt"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	herectx "github.com/endorama/devenv/internal/context"
	"github.com/endorama/devenv/internal/profile"
	"github.com/mitchellh/cli"
)

// Backup is a command to create an encrypted backup of a profile
type Backup struct {
	UI cli.Ui
}

// Synopsis returns short synopsis of the command.
func (cmd Backup) Synopsis() string {
	return "Create encrypted backup of profiles"
}

// Help return command help text
func (cmd Backup) Help() string {
	return fmt.Sprintf(`%s
The encryption passphrase is automatically generated using a safe RNG function and printed after backup creation.

Usage:
	devenv backup
		(not yet implemented)
		Backup all profiles

	devenv backup [PROFILE_NAME]
		Backup a single profile

	devenv backup [PROFILE_NAME...]
		(not yet implemented)
		Backup multiple profile

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

	petname.NonDeterministicMode()
	password := petname.Generate(6, "-")

	switch {
	case len(args) == 0:
		cmd.UI.Info("Creating backup for all profiles")
		// err = profile.RehashAllProfiles(ctx)
	case len(args) == 1:
		p, err := profile.New(ctx, args[0])
		if err != nil {
			cmd.UI.Error(err.Error())
			return 1
		}
		if !p.Exists() {
			cmd.UI.Error("Specified profile does not exist")
		}
		cmd.UI.Info(fmt.Sprintf("Creating backup for profile: %s", p.Name))
		err = profile.BackupSingle(ctx, p.Name, password)
		cmd.UI.Info("Included files:")
		files, err := p.Files()
		if err != nil {
			cmd.UI.Error(err.Error())
			return 1
		}
		for _, file := range files {
			cmd.UI.Info(fmt.Sprintf(" - %s", file))
		}
		cmd.UI.Info(fmt.Sprintf("Encryption passphrase is: %s", password))
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
