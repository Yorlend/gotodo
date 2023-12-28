package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() (*AppConfig, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ExtractDBConfig(cfg *AppConfig) *DBConfig {
	return &cfg.DB
}
