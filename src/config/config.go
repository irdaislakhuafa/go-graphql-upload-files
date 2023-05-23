package config

import (
	"context"
	"encoding/json"
	"os"
)

type Config struct {
	Application struct {
		Name string
		Port string
	}
	Database struct {
		Directory      string
		DirectoryFiles string
	}
}

func Init(ctx context.Context, configFilePath string) (Config, error) {
	fileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := json.Unmarshal(fileBytes, &config); err != nil {
		return Config{}, err
	}

	_, err = os.Stat(config.Database.Directory)
	if os.IsNotExist(err) {
		err := os.MkdirAll(config.Database.Directory, 0755)
		if err != nil {
			return Config{}, err
		}
	}

	return config, nil
}
