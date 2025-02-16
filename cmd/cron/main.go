package main

import (
	"fmt"
	"github.com/ElDanto/HashService/internal/config/cron"
	"github.com/ElDanto/HashService/internal/config/system"
	"github.com/ElDanto/HashService/internal/models/redis"
	"time"
)

var sysConf system.Conf
var config cron.CronConfig

func init() {
	sysConf.Init()
	config.Init(sysConf)
}

func main() {

	fmt.Println(config)
	redisHash := redis.Hash{}
	hash := redisHash.Generate()

	fmt.Println(hash)

	time.Sleep(time.Minute * 1)
}
