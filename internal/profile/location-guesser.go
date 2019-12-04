package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	configs "github.com/endorama/devenv/internal/configs"
)

func getProfileLocation(p Profile) (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	profiles := fmt.Sprintf("%s/%s", userHome, configs.ProfilesDirectory)
	guessedLocation := filepath.Join(profiles, p.Name)
	_, err = os.Stat(guessedLocation)
	if os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("Profile does not exists at %s", guessedLocation))
	} else if err != nil {
		return "", err
	}
	return guessedLocation, nil

}
