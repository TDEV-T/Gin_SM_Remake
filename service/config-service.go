package service

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ConnectConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
	JWT      string
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
			HOST:     os.Getenv("DB_HOST"),
			PORT:     os.Getenv("DB_PORT"),
			USER:     os.Getenv("DB_USER"),
			PASSWORD: os.Getenv("DB_PASS"),
			DBNAME:   os.Getenv("DB_NAME"),
			JWT:      os.Getenv("JWT_SECRET"),
		}
	})

	return Config
}
