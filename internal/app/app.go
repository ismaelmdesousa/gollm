package app

import (
	"fmt"

	"github.com/ismaelmdesousa/gollm/internal/pkg/logger"
)

type App struct {
	Name   string
	Env    string
	Logger *logger.Logger
}

type Option func(*App)

func WithLogger(l *logger.Logger) Option {
	return func(a *App) {
		a.Logger = l
	}
}

func New(cnf AppConfig, options ...Option) *App {
	app := &App{
		Name: cnf.GetName(),
		Env:  cnf.GetEnv(),
	}

	for _, option := range options {
		option(app)
	}

	if app.Logger == nil {
		panic("Logger is required for the application")
	}

	return app
}

func (a *App) Greet() string {
	return fmt.Sprintf("Hello, %s! This is running in %s environment.", a.Name, a.Env)
}
