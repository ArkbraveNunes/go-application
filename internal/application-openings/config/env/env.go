package env

import (
	"os"
)

type AppEnv struct {
	Port    string
	Version string
}

type DatabaseEnv struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func GetAppEnv() AppEnv {
	return newAppEnv()
}

func GetDatabaseEnv() DatabaseEnv {
	return newDatabaseEnv()
}

func newAppEnv() AppEnv {
	return AppEnv{
		Port:    os.Getenv("APPLICATION_OPENINGS__PORT"),
		Version: os.Getenv("APPLICATION_OPENINGS__VERSION"),
	}
}

func newDatabaseEnv() DatabaseEnv {
	return DatabaseEnv{
		Host:     os.Getenv("APPLICATION_OPENINGS__DB_HOST"),
		Port:     os.Getenv("APPLICATION_OPENINGS__DB_PORT"),
		Name:     os.Getenv("APPLICATION_OPENINGS__DB_NAME"),
		Username: os.Getenv("APPLICATION_OPENINGS__DB_USER"),
		Password: os.Getenv("APPLICATION_OPENINGS__DB_PASS"),
	}
}
