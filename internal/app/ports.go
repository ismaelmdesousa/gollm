package app

type AppConfig interface {
	GetName() string
	GetEnv() string
}
