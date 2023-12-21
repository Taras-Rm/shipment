package setup

import (
	"fmt"

	"github.com/Taras-Rm/shipment/config"
	"github.com/Taras-Rm/shipment/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var db *gorm.DB

	// get data from .env
	dbHost := config.GetDBHost()
	dbUser := config.GetDBUser()
	dbPass := config.GetDBPassword()
	dbName := config.GetDBName()

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate creating a table
	err = db.AutoMigrate(&repositories.ShipmentModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
