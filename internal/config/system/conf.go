package system

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Conf struct {
	CronConfigPath   string `yaml:"cron_config_path"`
	ServerConfigPath string `yaml:"server_config_path"`
}

func (c *Conf) Init() {
	data, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, &c)
}
