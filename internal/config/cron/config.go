package cron

import (
	"github.com/ElDanto/HashService/internal/config/system"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type CronConfig struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	LogFlag  bool   `default:"false" yaml:"logFlag"`
	Enable   bool   `default:"true" yaml:"enable"`
}

func (c *CronConfig) Init(conf system.Conf) {

	data, err := ioutil.ReadFile(conf.CronConfigPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, &c)
}
