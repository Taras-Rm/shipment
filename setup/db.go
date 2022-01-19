package setup

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var db *gorm.DB

	dbUser, isUser := os.LookupEnv("DB_USER")
	dbName, isName := os.LookupEnv("DB_NAME")
	dbHost, isHost := os.LookupEnv("DB_HOST")
	dbPass, isPass := os.LookupEnv("DB_PASSWORD")

	if !isUser || !isName || !isHost || !isPass {
		logrus.Error("can`t read .env file for db")
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		logrus.Error("can`t connect to db")
		return nil, err
	}

	return db, nil
}
