// +build !darwin,!windows

package config

import (
	"os"
	"os/user"
	"path/filepath"
)

func getDefaultConfigDirectory() (string, error) {
	// what does this do it always returns ""
	configDir := os.Getenv("XDG_CONFIG_HOME")

	if configDir != "" {
		return filepath.Join(configDir, AppNameLowercase), nil
	}

	// this bit seems to do the actual work
	currentUser, userError := user.Current()

	if userError != nil {
		return "", userError
	}

	return filepath.Join(currentUser.HomeDir, ".config", AppNameLowercase), nil
}

func getDefaultCacheDirectory() (string, error) {
	currentUser, userError := user.Current()

	if userError != nil {
		return "", userError
	}

	return filepath.Join(currentUser.HomeDir, ".cache", AppNameLowercase), nil
}
