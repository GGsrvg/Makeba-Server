package main

import (
	"flag"
	"log"
	"makeba/internal/server"
	"makeba/internal/util"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {

	dir, err := util.CurrentDir()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&configPath, "config-path", dir+"/configs/main.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
