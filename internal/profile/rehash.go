package profile

import (
	"context"
	"fmt"

	herectx "github.com/endorama/devenv/internal/context"
	"github.com/mitchellh/cli"
)

func RehashAllProfiles(ctx context.Context) error {
	profiles, err := getProfiles()
	if err != nil {
		return fmt.Errorf("cannot retrieve profiles: %w", err)
	}
	
	ui := ctx.Value(herectx.UI).(*cli.BasicUi)
	ui.Info(fmt.Sprintf("Profiles: %s", profiles))
	for _, name := range profiles {
		err = RehashSingle(ctx, name)
		if err != nil {
			ui.Error(err.Error())
		}
	}

	return nil
}

func RehashSingle(ctx context.Context, profileName string) error {
	profile, err := New(ctx, profileName)
	if err != nil {
		return err
	}

	profile.LoadPlugins()

	err = profile.LoadPluginConfigurations(ctx)
	if err != nil {
		return err
	}
	err = profile.SetupPlugins(ctx)
	if err != nil {
		return err
	}
	err = profile.RunPluginGeneration(ctx)
	if err != nil {
		return err
	}
	err = profile.GenerateShellLoadFile(ctx)
	if err != nil {
		return err
	}

	return nil
}
