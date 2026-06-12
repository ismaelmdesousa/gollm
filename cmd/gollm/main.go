package main

import (
	"flag"

	"github.com/ismaelmdesousa/gollm/internal/app"
	"github.com/ismaelmdesousa/gollm/internal/pkg/config"
)

var (
	configPath = flag.String("config", "config.yml", "Path to the configuration file")
)

func main() {
	flag.Parse()
	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	app := app.New(cfg.GetApp())
	println(app.Greet())
}
