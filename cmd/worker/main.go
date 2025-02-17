package main

import (
	"fmt"
	"github.com/ElDanto/ShortGen/internal/configs"
	"log"
	"time"
)

func main() {
	cnf, err := configs.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println(cnf)

	time.Sleep(time.Minute * 1)
}
