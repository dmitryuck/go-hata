package main

import (
	"flag"
	"log"

	"project/internal/apiserver"
	"project/internal/config"
	"project/internal/logger"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	logger.Instance = logger.New(config)

	apiserver.Instance = apiserver.New(config)

	if err := apiserver.Instance.Start(); err != nil {
		log.Fatal(err)
	}
}
