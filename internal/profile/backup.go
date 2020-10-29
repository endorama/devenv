package profile

import (
	"context"

	"github.com/endorama/devenv/internal/backup"
)

// BackupSingle create backup for single profile specified by name
func BackupSingle(ctx context.Context, profileName, password string) error {
	profile, err := New(ctx, profileName)
	if err != nil {
		return err
	}

	files, err := profile.Files()
	if err != nil {
		return err
	}

	err = backup.EncryptedBackup(profile.Name, password, profile.Location, files)
	if err != nil {
		return err
	}

	return nil
}
