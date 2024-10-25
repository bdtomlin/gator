package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	var cfg Config

	filepath, err := getConfigFilePath()
	if err != nil {
		return &cfg, err
	}

	rawJson, err := os.ReadFile(filepath)
	if err != nil {
		return &cfg, fmt.Errorf("unable to read config file: %w", err)
	}

	err = json.Unmarshal(rawJson, &cfg)
	if err != nil {
		return &cfg, fmt.Errorf("Error reading json: %w", err)
	}
	return &cfg, nil
}

func (cfg *Config) SetUser(uname string) error {
	cfg.CurrentUserName = uname

	jsn, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %w", err)
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	os.WriteFile(path, jsn, 0766)
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get filepath: %w", err)
	}
	return homeDir + "/" + configFileName, nil
}
