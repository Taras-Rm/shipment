package main

import (
	"github.com/Taras-Rm/shipment/api"
	"github.com/Taras-Rm/shipment/config"
	"github.com/Taras-Rm/shipment/repositories"
	"github.com/Taras-Rm/shipment/services"
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
	group := handler.Group("api")

	// db connection
	db, err := setup.ConnectDB()
	if err != nil {
		panic(err)
	}

	shipmentRepository := repositories.InitShipmentRepository(db)
	shipmentService := services.InitShipmentService(shipmentRepository)
	api.UseShipment(group, shipmentService)

	// start server
	handler.Run(port)
}
