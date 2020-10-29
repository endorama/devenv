package archive

import (
	"fmt"
	"io"

	"filippo.io/age"
)

func encrypt(password string, in io.Reader, out io.Writer) error {
	r, err := age.NewScryptRecipient(password)
	if err != nil {
		return fmt.Errorf("failed creating scrypt recipient: %w", err)
	}

	w, err := age.Encrypt(out, r)
	if err != nil {
		return fmt.Errorf("failed encrypting: %w", err)
	}
	if _, err := io.Copy(w, in); err != nil {
		return fmt.Errorf("failed copying data: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("failed closing writer: %w", err)
	}
	return nil
}
