package repositories

import (
	"github.com/Taras-Rm/shipment/models"
	"gorm.io/gorm"
)

type ShipmentRepository interface {
	GetAllShipments() ([]models.Shipment, error)
	CreateShipment(shipment *models.Shipment) error
	GetShipmentByID(shipmentID uint) (*models.Shipment, error)
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

	if res.Error != nil {
		return nil, res.Error
	}

	return shipments, nil
}

// create a new shipment
func (r *shipmentRepository) CreateShipment(shipment *models.Shipment) error {
	res := r.db.Create(&shipment)
	return res.Error
}

// get a single shipment by ID
func (r *shipmentRepository) GetShipmentByID(shipmentID uint) (*models.Shipment, error) {
	var shipment *models.Shipment
	res := r.db.First(&shipment, shipmentID)
	return shipment, res.Error
}
