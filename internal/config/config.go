package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFile = ".gatorconfig.json"

type Config struct {
	UserName string `json:"current_user_name"`
	DBUrl    string `json:"db_url"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	config := Config{}
	if err != nil {
		return config, err
	}
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil

}

func (cfg Config) SetUser(UserName string) error {
	cfg.UserName = UserName
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := path.Join(home, configFile)
	return filePath, nil
}
