package service

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ConnectConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var (
	Config *ConnectConfig
	once   sync.Once
)

func LoadConfig() *ConnectConfig {
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			panic("Faile to Load ENV")
		}

		Config = &ConnectConfig{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASS"),
			dbname:   os.Getenv("DB_NAME"),
		}
	})

	return Config
}
