package app

import (
	"fmt"
)

type App struct {
	Name string
	Env  string
}

func New(cnf AppConfig) *App {
	return &App{
		Name: cnf.GetName(),
		Env:  cnf.GetEnv(),
	}
}

func (a *App) Greet() string {
	return fmt.Sprintf("Hello, %s!\nThis is running in %s environment.", a.Name, a.Env)
}
