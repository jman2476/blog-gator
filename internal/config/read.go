package config

import (
	"encoding/json"
	"os"
)

func Read() (*Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return &Config{}, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}
