package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetServerPort() string {
	str, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		logrus.Error("can`t read .env file")
		return ""
	}
	return str
}
