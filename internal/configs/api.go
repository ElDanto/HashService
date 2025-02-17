package configs

import "time"

type APIConfig struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	RateLimit    int           `yaml:"rate_limit"`
	EnableHTTPS  bool          `yaml:"enable_https"`
}
