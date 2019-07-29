package profile

import (
	"context"
)

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
