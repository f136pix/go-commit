package configstore

import (
	"encoding/json"
	"go-commit/cmd/env"
	"os"
)

type Config struct {
	Providers map[string]string `json:"providers"`
}

func LoadConfig() *Config {
	cfg := &Config{Providers: map[string]string{}}
	file, err := os.Open(env.ConfigFilePath)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(cfg)
	}
	return cfg
}

func SaveConfig(cfg *Config) {
	file, err := os.Create(env.ConfigFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(cfg)
}
