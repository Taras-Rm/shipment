package services

import (
	"errors"

	"github.com/Taras-Rm/shipment/helpers"
	"github.com/Taras-Rm/shipment/models"
	"github.com/Taras-Rm/shipment/repositories"
)

type AddShipmentInput struct {
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

func (i AddShipmentInput) Validate() error {
	// check email
	fromErr := helpers.ValidateEmail(i.FromEmail)
	toErr := helpers.ValidateEmail(i.ToEmail)
	if fromErr != nil || toErr != nil {
		return errors.New("invalid email format")
	}

	// check names
	nameFromErr := helpers.ValidateName(i.FromName)
	nameToErr := helpers.ValidateName(i.ToName)
	if nameFromErr != nil || nameToErr != nil {
		return errors.New("invalid name format")
	}

	// check country codes
	codeErr := helpers.ValidateCountryCode(i.FromCountryCode)
	if codeErr != nil {
		return codeErr
	}
	codeErr = helpers.ValidateCountryCode(i.ToCountryCode)
	if codeErr != nil {
		return codeErr
	}

	// check addresses
	addressFromErr := helpers.ValidateAddress(i.FromAddress)
	addressToErr := helpers.ValidateAddress(i.ToAddress)
	if addressFromErr != nil || addressToErr != nil {
		return errors.New("invalid address format")
	}

	// check weight
	if i.Weight <= 0 || i.Weight > 1000 {
		return errors.New("invalid weight")
	}

	return nil
}

//go:generate mockgen -source=shipment.go -destination=mocks/shipment.go
type ShipmentService interface {
	GetAllShipments() ([]models.Shipment, error)
	AddShipment(inp AddShipmentInput) (float64, error)
	GetShipmentByID(id uint) (models.Shipment, error)
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

func (s *shipmentService) AddShipment(inp AddShipmentInput) (float64, error) {
	// determine Region Rules
	regionFactor := helpers.RegionRulesFactor(inp.FromCountryCode)

	// determine Weight Class Rules
	weightFactor := helpers.WeightClassRulesFactor(inp.Weight)

	// calculate price
	price := regionFactor * float64(weightFactor)

	shipment := models.Shipment{
		FromName:        inp.FromName,
		FromEmail:       inp.FromEmail,
		FromAddress:     inp.FromAddress,
		FromCountryCode: inp.FromCountryCode,
		ToName:          inp.ToName,
		ToEmail:         inp.ToEmail,
		ToAddress:       inp.ToAddress,
		ToCountryCode:   inp.ToCountryCode,
		Weight:          inp.Weight,
		Price:           price,
	}

	// add the new shipment to the database
	err := s.shipmentRepository.CreateShipment(shipment)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *shipmentService) GetShipmentByID(id uint) (models.Shipment, error) {
	shipment, err := s.shipmentRepository.GetShipmentByID(id)
	if err != nil {
		return models.Shipment{}, err
	}

	return shipment, nil
}
