package profile

import (
	"context"
)

func RehashSingleProfile(ctx context.Context, profileName string) error {
	profile, err := NewProfile(profileName)
	if err != nil {
		return err
	}
	err = profile.Rehash(ctx)
	if err != nil {
		return err
	}
	
	return nil
}
