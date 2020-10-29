package backup

import (
	"fmt"
	"os"

	"github.com/endorama/devenv/internal/archive"
)

// EncryptedBackup create an encrypted backup archive from the specified set of files, using the
// specified passphrase. To support relative paths, it allows a cwd parameter to change directory
// before creating the file.
// Archive will be created in the current folder
func EncryptedBackup(name, password, cwd string, files []string) error {
	// create output file
	out, err := os.Create(fmt.Sprintf("%s.tar.gz.age", name))
	if err != nil {
		return fmt.Errorf("cannot create archive: %w", err)
	}
	defer out.Close()

	// change folder to specified location, so is possible to use relative
	// paths in the archive
	err = os.Chdir(cwd)
	if err != nil {
		return err
	}

	// perform archive creation, compression and encryption
	err = archive.Create(out, files, password)
	if err != nil {
		return fmt.Errorf("failed creating encrypted backup archive: %w", err)
	}

	return nil
}
