package models

type Shipment struct {
	Id              uint
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
