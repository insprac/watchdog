package config

import (
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	configDir := os.Getenv("WATCHDOG_DIR")
	if configDir == "" {
		panic("Environment variable WATCHDOG_DIR is not set")
	}
	return configDir
}

func GetConfigFile() string {
	configDir := GetConfigDir()
	return filepath.Join(configDir, "config.yaml")
}

func GetStateFile() string {
	configDir := GetConfigDir()
	return filepath.Join(configDir, "state.yaml")
}
