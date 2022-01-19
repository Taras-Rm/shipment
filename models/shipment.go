package models

import "gorm.io/gorm"

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
	Weight          uint
	Price           uint
}
