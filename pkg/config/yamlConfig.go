package config

import (
	"fmt"
	"os"

	"DiscordTemplate/pkg/logger"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	Token string
	logger logger.Logger
}

func NewYamlConfig(file string, logger logger.Logger) *yamlConfig {
	cfg := &yamlConfig{
		logger: logger,
	}
	err := cfg.loadConfig(file)
	if err != nil {
		logger.Fatal("Failed to load config file: %v", err)
	}
	err = cfg.validateConfig()
	if err != nil {
		logger.Fatal("Failed to validate config file: %v", err)
	}
	return cfg
}

func (c *yamlConfig) loadConfig(file string) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *yamlConfig) validateConfig() error {
	if c.Token == "" {
		return fmt.Errorf("empty token")
	}
	return nil
}