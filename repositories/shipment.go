package repositories

import (
	"github.com/Taras-Rm/shipment/models"
	"gorm.io/gorm"
)

// shipment model
type ShipmentModel struct {
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

func ShipmentModelToDomain(shipment ShipmentModel) models.Shipment {
	return models.Shipment{
		Id:              shipment.ID,
		FromName:        shipment.FromName,
		FromEmail:       shipment.FromEmail,
		FromAddress:     shipment.FromAddress,
		FromCountryCode: shipment.FromCountryCode,
		ToName:          shipment.ToName,
		ToEmail:         shipment.ToEmail,
		ToAddress:       shipment.ToAddress,
		ToCountryCode:   shipment.ToAddress,
		Weight:          shipment.Weight,
		Price:           shipment.Price,
	}
}

func ShipmentModelFromDomain(shipment models.Shipment) ShipmentModel {
	return ShipmentModel{
		FromName:        shipment.FromName,
		FromEmail:       shipment.FromEmail,
		FromAddress:     shipment.FromAddress,
		FromCountryCode: shipment.FromCountryCode,
		ToName:          shipment.ToName,
		ToEmail:         shipment.ToEmail,
		ToAddress:       shipment.ToAddress,
		ToCountryCode:   shipment.ToAddress,
		Weight:          shipment.Weight,
		Price:           shipment.Price,
	}
}

type ShipmentRepository interface {
	GetAllShipments() ([]models.Shipment, error)
	CreateShipment(shipment models.Shipment) error
	GetShipmentByID(shipmentID uint) (models.Shipment, error)
}

type shipmentRepository struct {
	db *gorm.DB
}

func InitShipmentRepository(db *gorm.DB) ShipmentRepository {
	return &shipmentRepository{db: db}
}

// get all shipments that have been sent to the system
func (r *shipmentRepository) GetAllShipments() ([]models.Shipment, error) {
	var shipments []models.Shipment
	res := r.db.Find(&shipments)

	return shipments, res.Error
}

// create a new shipment
func (r *shipmentRepository) CreateShipment(shipment models.Shipment) error {
	model := ShipmentModelFromDomain(shipment)

	res := r.db.Create(&model)

	return res.Error
}

// get a single shipment by it's ID
func (r *shipmentRepository) GetShipmentByID(shipmentID uint) (models.Shipment, error) {
	var shipment models.Shipment

	res := r.db.First(&shipment, shipmentID)

	return shipment, res.Error
}
