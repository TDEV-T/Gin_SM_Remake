package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"Gin_Remake/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB_Con  *gorm.DB
	db_once sync.Once
)

func LoadConnect() *gorm.DB {
	var db *gorm.DB

	db_once.Do(func() {
		db = setUPDB()

		db.AutoMigrate(&models.User{})
	})

	DB_Con = db

	return db
}

func CloseConnect(db *gorm.DB) {

}

func setUPDB() *gorm.DB {
	env := LoadConfig()

	port, err := strconv.Atoi(env.port)

	if err != nil {
		panic("Can't Change String to interger")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.host, port, env.user, env.password, env.dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {

		fmt.Println(dsn)
		fmt.Println(err.Error())
		panic("Error Can't Connect Database ")
	}

	return db
}
