package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	cfg                     *Config
	ErrUnknownFileExtension = errors.New("unknown file extension")
)

func Parse(path string, cfg *Config) error {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYAML(path, cfg)
	default:
		return ErrUnknownFileExtension
	}
}

func parseYAML(path string, cfg *Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if e := file.Close(); err == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
}

func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

func SetConfig(c *Config) {
	cfg = c
}
