package env

import (
	"os"
	"path/filepath"
)

var ConfigFilePath = filepath.Join(os.Getenv("HOME"), ".go-commit-config.json")
var ValidProviders = []string{"open-ai", "anthropic", "google"}
