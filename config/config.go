package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// get server port from .env
func GetServerPort() string {
	str, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		logrus.Error("can`t read .env file (server port)")
		return ""
	}
	return str
}

// get DB user from .env
func GetDBUser() string {
	str, ok := os.LookupEnv("DB_USER")
	if !ok {
		logrus.Error("can`t read .env file (DB user)")
		return ""
	}
	return str
}

// get DB name from .env
func GetDBName() string {
	str, ok := os.LookupEnv("DB_NAME")
	if !ok {
		logrus.Error("can`t read .env file (DB name)")
		return ""
	}
	return str
}

// get DB host from .env
func GetDBHost() string {
	str, ok := os.LookupEnv("DB_HOST")
	if !ok {
		logrus.Error("can`t read .env file (DB host)")
		return ""
	}
	return str
}

// get DB password from .env
func GetDBPassword() string {
	str, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		logrus.Error("can`t read .env file (DB password)")
		return ""
	}
	return str
}
