package configs

import "time"

type WorkerConfig struct {
	CleanupInterval time.Duration `yaml:"cleanup_interval"`
	BatchSize       int           `yaml:"batch_size"`
	MaxAttempts     int           `yaml:"max_attempts"`
}
