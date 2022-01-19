package services

import (
	"errors"

	"github.com/Taras-Rm/shipment/helpers"
	"github.com/Taras-Rm/shipment/models"
	"github.com/Taras-Rm/shipment/repositories"
)

type ShipmentRequest struct {
	FromName        string  `json:"fromName" binding:"required"`
	FromEmail       string  `json:"fromEmail" binding:"required"`
	FromAddress     string  `json:"fromAddress" binding:"required"`
	FromCountryCode string  `json:"fromCountryCode" binding:"required"`
	ToName          string  `json:"toName" binding:"required"`
	ToEmail         string  `json:"toEmail" binding:"required"`
	ToAddress       string  `json:"toAddress" binding:"required"`
	ToCountryCode   string  `json:"toCountryCode" binding:"required"`
	Weight          float64 `json:"weight" binding:"required"`
}

type ShipmentService interface {
	GetAllShipments() ([]models.Shipment, error)
	AddShipment(shipmentReq *ShipmentRequest) (float64, error)
	GetShipmentByID(id uint) (*models.Shipment, error)
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

func (s *shipmentService) AddShipment(shipmentReq *ShipmentRequest) (float64, error) {
	// determine Region Rules (coefficient)
	regionCoef := helpers.RegionRulesCoef(shipmentReq.FromCountryCode)
	// determine Weight Class Rules (coefficient)
	weightCoef := helpers.WeightClassRulesCoef(shipmentReq.Weight)

	// calculate price
	price := regionCoef * float64(weightCoef)

	shipment := models.Shipment{
		FromName:        shipmentReq.FromName,
		FromEmail:       shipmentReq.FromEmail,
		FromAddress:     shipmentReq.FromAddress,
		FromCountryCode: shipmentReq.FromCountryCode,
		ToName:          shipmentReq.ToName,
		ToEmail:         shipmentReq.ToEmail,
		ToAddress:       shipmentReq.ToAddress,
		ToCountryCode:   shipmentReq.ToCountryCode,
		Weight:          shipmentReq.Weight,
		Price:           price,
	}

	err := s.shipmentRepository.CreateShipment(&shipment)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *shipmentService) GetShipmentByID(id uint) (*models.Shipment, error) {
	shipment, err := s.shipmentRepository.GetShipmentByID(id)

	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func IsValidShipmentRequest(shipmentReq *ShipmentRequest) error {

	fromErr := helpers.ValidateEmail(shipmentReq.FromEmail)
	toErr := helpers.ValidateEmail(shipmentReq.ToEmail)
	if fromErr != nil || toErr != nil {
		return errors.New("uncorrect email")
	}

	nameFromErr := helpers.ValidateName(shipmentReq.FromName)
	nameToErr := helpers.ValidateName(shipmentReq.ToName)
	if nameFromErr != nil || nameToErr != nil {
		return errors.New("uncorrect name")
	}

	codeFromErr := helpers.ValidateCountryCode(shipmentReq.FromCountryCode)
	codeToErr := helpers.ValidateCountryCode(shipmentReq.ToCountryCode)
	if codeFromErr != nil || codeToErr != nil {
		return errors.New("uncorrect country code")
	}

	if shipmentReq.Weight <= 0 || shipmentReq.Weight > 1000 {
		return errors.New("invalid weight")
	}

	return nil
}
