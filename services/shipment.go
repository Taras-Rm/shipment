package services

import (
	"errors"

	"github.com/Taras-Rm/shipment/helpers"
	"github.com/Taras-Rm/shipment/models"
	"github.com/Taras-Rm/shipment/repositories"
)

// structure of shipment request
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

// structure of shipment response
type ShipmentResponse struct {
	ID              uint    `json:"id" binding:"required"`
	FromName        string  `json:"fromName" binding:"required"`
	FromEmail       string  `json:"fromEmail" binding:"required"`
	FromAddress     string  `json:"fromAddress" binding:"required"`
	FromCountryCode string  `json:"fromCountryCode" binding:"required"`
	ToName          string  `json:"toName" binding:"required"`
	ToEmail         string  `json:"toEmail" binding:"required"`
	ToAddress       string  `json:"toAddress" binding:"required"`
	ToCountryCode   string  `json:"toCountryCode" binding:"required"`
	Weight          float64 `json:"weight" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
}

type ShipmentService interface {
	GetAllShipments() ([]ShipmentResponse, error)
	AddShipment(shipmentReq *ShipmentRequest) (float64, error)
	GetShipmentByID(id uint) (*ShipmentResponse, error)
}

type shipmentService struct {
	shipmentRepository repositories.ShipmentRepository
}

func InitShipmentService(shipmentRepo repositories.ShipmentRepository) ShipmentService {
	return &shipmentService{shipmentRepository: shipmentRepo}
}

func (s *shipmentService) GetAllShipments() ([]ShipmentResponse, error) {
	shipments, err := s.shipmentRepository.GetAllShipments()
	if err != nil {
		return nil, err
	}

	// formation of all shipments response
	newShipments := prepareShipmentsToResponse(shipments)

	return newShipments, nil
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

	// add the new shipment to the database
	err := s.shipmentRepository.CreateShipment(&shipment)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *shipmentService) GetShipmentByID(id uint) (*ShipmentResponse, error) {
	shipment, err := s.shipmentRepository.GetShipmentByID(id)
	if err != nil {
		return nil, err
	}

	// formation of shipment response
	newShipment := &ShipmentResponse{
		ID:              shipment.ID,
		FromName:        shipment.FromName,
		FromEmail:       shipment.FromEmail,
		FromAddress:     shipment.FromAddress,
		FromCountryCode: shipment.FromCountryCode,
		ToName:          shipment.ToName,
		ToEmail:         shipment.ToEmail,
		ToAddress:       shipment.ToAddress,
		ToCountryCode:   shipment.ToCountryCode,
		Weight:          shipment.Weight,
		Price:           shipment.Price,
	}

	return newShipment, nil
}

func IsValidShipmentRequest(shipmentReq *ShipmentRequest) error {

	fromErr := helpers.ValidateEmail(shipmentReq.FromEmail)
	toErr := helpers.ValidateEmail(shipmentReq.ToEmail)
	if fromErr != nil || toErr != nil {
		return errors.New("uncorrect email format")
	}

	nameFromErr := helpers.ValidateName(shipmentReq.FromName)
	nameToErr := helpers.ValidateName(shipmentReq.ToName)
	if nameFromErr != nil || nameToErr != nil {
		return errors.New("uncorrect name format")
	}

	codeFromErr := helpers.ValidateCountryCode(shipmentReq.FromCountryCode)
	codeToErr := helpers.ValidateCountryCode(shipmentReq.ToCountryCode)
	if codeFromErr != nil || codeToErr != nil {
		return errors.New("uncorrect country code")
	}

	addressFromErr := helpers.ValidateAddress(shipmentReq.FromAddress)
	addressToErr := helpers.ValidateAddress(shipmentReq.ToAddress)
	if addressFromErr != nil || addressToErr != nil {
		return errors.New("uncorrect address format")
	}

	if shipmentReq.Weight <= 0 || shipmentReq.Weight > 1000 {
		return errors.New("invalid weight")
	}

	return nil
}

func prepareShipmentsToResponse(shipments []models.Shipment) []ShipmentResponse {
	var newShipments []ShipmentResponse

	for _, val := range shipments {
		shipment := ShipmentResponse{
			ID:              val.ID,
			FromName:        val.FromName,
			FromEmail:       val.FromEmail,
			FromAddress:     val.FromAddress,
			FromCountryCode: val.FromCountryCode,
			ToName:          val.ToName,
			ToEmail:         val.ToEmail,
			ToAddress:       val.ToAddress,
			ToCountryCode:   val.ToCountryCode,
			Weight:          val.Weight,
			Price:           val.Price,
		}

		newShipments = append(newShipments, shipment)
	}

	return newShipments
}
