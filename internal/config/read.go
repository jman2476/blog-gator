package config

import (
	"encoding/json"
	"os"
)

func Read() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return &Config{}, err
	}

	configFile := home + "/.gatorconfig.json"
	data, err := os.ReadFile(configFile)
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
