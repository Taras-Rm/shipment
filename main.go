package main

import (
	"github.com/Taras-Rm/shipment/config"
	"github.com/Taras-Rm/shipment/setup"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// getting server port
	port := config.GetServerPort()

	// server handler
	handler := setup.ServerStart()

	// db connection
	_, err = setup.ConnectDB()
	if err != nil {
		panic(err)
	}

	// start server
	handler.Run(port)
}
