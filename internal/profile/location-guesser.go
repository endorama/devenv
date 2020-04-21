package profile

import (
	"context"
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

func getProfiles() ([]string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return []string{}, err
	}

	profilesDirectory := fmt.Sprintf("%s/%s", userHome, configs.ProfilesDirectory)
	file, err := os.Open(profilesDirectory)
	if err != nil {
		return []string{}, err
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		return []string{}, err
	}

	profiles := []string{}
	for _, name := range names {
		_, err := New(context.Background(), name)
		if err == nil {
			profiles = append(profiles, name)
		}
	}
	return profiles, nil
}
