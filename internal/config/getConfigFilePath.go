package config

import "os"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := home + configFileName

	return configPath, nil
}
