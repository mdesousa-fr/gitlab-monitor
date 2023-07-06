package main

import (
	"fmt"
	"log"

	"github.com/mdesousa-fr/gitlab-monitor/internal/config"
)

func main() {
	config, err := config.ReadConfig("example.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}