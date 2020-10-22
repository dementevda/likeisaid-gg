package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/dementevda/likeisaid-gg/backend/cmd/api"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/toml_config.toml", "Path to conifg file")
}

func main() {
	flag.Parse()

	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(config)
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("finish")
}
