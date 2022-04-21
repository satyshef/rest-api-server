package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/golang/restapi/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()
	conf := apiserver.MakeConfig()
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal(err)
	}
	a := apiserver.New(conf)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

}
