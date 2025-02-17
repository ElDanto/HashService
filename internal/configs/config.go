package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Env    string       `yaml:"env"`
	API    APIConfig    `yaml:"api"`
	Worker WorkerConfig `yaml:"worker"`
	Redis  RedisConfig  `yaml:"model"`
}

func Load(configPath string) (*AppConfig, error) {
	cfg := &AppConfig{
		Env: "development", // значение по умолчанию
	}

	// Чтение YAML-файла
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Валидация конфигурации
	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validate(cfg *AppConfig) error {
	if cfg.Redis.Addr == "" {
		return fmt.Errorf("model addr is required")
	}

	if cfg.API.Port < 1 || cfg.API.Port > 65535 {
		return fmt.Errorf("invalid API port")
	}

	return nil
}
