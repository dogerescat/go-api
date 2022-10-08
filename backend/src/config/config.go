package config

import (
	"os"
)

type config struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

var Config config

func init() {
	Config.Host = os.Getenv("DB_HOST")
	Config.User = os.Getenv("DB_USER")
	Config.Password = os.Getenv("DB_PASSWORD")
	Config.Port = os.Getenv("DB_PORT")
	Config.Database = os.Getenv("DB_DATABASE")
}
