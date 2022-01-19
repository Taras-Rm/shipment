package services

import (
	"github.com/Taras-Rm/shipment/models"
	"github.com/Taras-Rm/shipment/repositories"
)

type ShipmentRequest struct {
	FromName        string `json:"fromName" binding:"required"`
	FromEmail       string `json:"fromEmail" binding:"required"`
	FromAddress     string `json:"fromAddress" binding:"required"`
	FromCountryCode string `json:"fromCountryCode" binding:"required"`
	ToName          string `json:"toName" binding:"required"`
	ToEmail         string `json:"toEmail" binding:"required"`
	ToAddress       string `json:"toAddress" binding:"required"`
	ToCountryCode   string `json:"toCountryCode" binding:"required"`
	Weight          uint   `json:"weight" binding:"required"`
}

type ShipmentService interface {
	GetAllShipments() ([]models.Shipment, error)
	//AddShipment(shipmentReq *ShipmentRequest) (uint, error)
	//GetShipmentByID(ID uint) (*models.Shipment, error)
}

type shipmentService struct {
	shipmentRepository repositories.ShipmentRepository
}

func InitShipmentService(shipmentRepo repositories.ShipmentRepository) ShipmentService {
	return &shipmentService{shipmentRepository: shipmentRepo}
}

func (s *shipmentService) GetAllShipments() ([]models.Shipment, error) {
	shipments, err := s.shipmentRepository.GetAllShipments()

	if err != nil {
		return nil, err
	}

	return shipments, nil
}
