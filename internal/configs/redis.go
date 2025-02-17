package configs

import "time"

type RedisConfig struct {
	Addr        string        `yaml:"addr"`
	PoolSize    int           `yaml:"pool_size"`
	TokenTTL    time.Duration `yaml:"token_ttl"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
}
