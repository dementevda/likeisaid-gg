package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/dementevda/likeisaid-gg/backend/internal/api"
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
	CloseHandler(api)

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}

func CloseHandler(api *api.API) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		api.Stop()
		os.Exit(0)
	}()
}
