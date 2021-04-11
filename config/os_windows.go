package config

import (
	"os"
	"os/user"
	"path/filepath"
)

func getDefaultConfigDirectory() (string, error) {
	configDir := os.Getenv("APPDATA")

	if configDir != "" {
		return filepath.Join(configDir, AppNameLowercase), nil
	}

	currentUser, userError := user.Current()

	if userError != nil {
		return "", userError
	}

	return filepath.Join(currentUser.HomeDir, "AppData", "Roaming", AppNameLowercase), nil
}

func getDefaultCacheDirectory() (string, error) {
	configDir := os.Getenv("LOCALAPPDATA")

	if configDir != "" {
		return filepath.Join(configDir, AppNameLowercase, "cache"), nil
	}

	currentUser, userError := user.Current()

	if userError != nil {
		return "", userError
	}

	return filepath.Join(currentUser.HomeDir, "AppData", "Local", AppNameLowercase, "cache"), nil
}
