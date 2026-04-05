package config

import (
	"encoding/json"
	"os"
)

func write(c *Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
