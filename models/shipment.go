package models

import "gorm.io/gorm"

// shipment model
type Shipment struct {
	gorm.Model
	FromName        string
	FromEmail       string
	FromAddress     string
	FromCountryCode string
	ToName          string
	ToEmail         string
	ToAddress       string
	ToCountryCode   string
	Weight          float64
	Price           float64
}
