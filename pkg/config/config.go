package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type StorageConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	BotToken string `yaml:"bot_token"`
}

func LoadConfig(path string) (*StorageConfig, error) {
	cfg := &StorageConfig{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
