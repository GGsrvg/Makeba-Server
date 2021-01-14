package main

import (
	"flag"
	"log"
	"makeba/internal/server"
	"makeba/internal/util"
	"makeba/internal/web"
	"sync"

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

	dir, err := util.CurrentDir()

	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		startApiServer(dir)
		wg.Done()
	}()
	go func() {
		startWebServer(dir)
		wg.Done()
	}()

	wg.Wait()
}

func startApiServer(dir string) {
	config := server.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(dir, config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func startWebServer(dir string) {
	webConfig := web.NewConfig()

	_, err := toml.DecodeFile(configPath, webConfig)
	if err != nil {
		log.Fatal(err)
	}

	web := web.New(dir, webConfig)
	if err := web.Start(); err != nil {
		log.Fatal(err)
	}
}
