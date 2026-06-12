package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration structure.
type App struct {
	Name        string `yaml:"name"`
	Env         string `yaml:"env"`
	LoggerLevel string `yaml:"logger_level"`
}

func (a *App) GetName() string {
	return a.Name
}

func (a *App) GetEnv() string {
	return a.Env
}

func (a *App) GetLoggerLevel() string {
	return a.LoggerLevel
}

// Config is the main configuration struct that holds all configuration sections.
type Config struct {
	App App `yaml:"app"`
}

func (c *Config) GetApp() *App {
	return &c.App
}

// Load reads the configuration from the specified YAML file and unmarshals it into a Config struct.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
