package main

import (
	"flag"

	"github.com/ismaelmdesousa/gollm/internal/app"
	"github.com/ismaelmdesousa/gollm/internal/pkg/config"
	"github.com/ismaelmdesousa/gollm/internal/pkg/logger"
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

	log := logger.NewLogger(cfg.GetApp().GetLoggerLevel())

	log.Info("Configuration loaded successfully")

	log.Info("Start application")
	app := app.New(cfg.GetApp(), app.WithLogger(log))
	log.Info(app.Greet())
}
