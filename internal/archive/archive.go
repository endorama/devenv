package archive

import (
	"bytes"
	"fmt"
	"io"
)

// Create creates an encrypted gzipped tar archive file
func Create(out io.Writer, files []string, password string) error {
	var targz bytes.Buffer

	err := createArchive(files, &targz)
	if err != nil {
		return fmt.Errorf("cannot create archive: %w", err)
	}

	err = encrypt(password, &targz, out)
	if err != nil {
		return fmt.Errorf("cannot encrypt archive: %w", err)
	}

	return nil
}
