package profile

import (
	"context"
)

func RehashSingleProfile(ctx context.Context, profileName string) error {
	profile, err := NewProfile(profileName)
	if err != nil {
		return err
	}

	awsPlugin := NewAwsPlugin()
	profile.EnablePlugin(awsPlugin)
	binPlugin := NewBinPlugin()
	profile.EnablePlugin(binPlugin)
	shellHistoryPlugin := NewShellHistoryPlugin()
	profile.EnablePlugin(shellHistoryPlugin)
	emailPlugin := NewEmailPlugin()
	profile.EnablePlugin(emailPlugin)
	envsPlugin := NewEnvsPlugin()
	profile.EnablePlugin(envsPlugin)
	// profile.EnablePlugin("gpg")
	// profile.EnablePlugin("shell-history")
	// profile.EnablePlugin("ssh")

	err = profile.LoadPluginConfigurations(ctx)
	if err != nil {
		return err
	}

	err = profile.Rehash(ctx)
	if err != nil {
		return err
	}

	return nil
}
