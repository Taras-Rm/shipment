package services

import (
	"errors"
	"testing"

	"github.com/Taras-Rm/shipment/models"
	mock_repositories "github.com/Taras-Rm/shipment/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService_AddShipment(t *testing.T) {
	type mockBehaviur func(r *mock_repositories.MockShipmentRepository, shipment models.Shipment)

	testCases := []struct {
		name          string
		input         AddShipmentInput
		inputShipment models.Shipment
		mockBehaviur  mockBehaviur
		expectedPrice float64
		expectedError error
	}{
		{
			name: "Ok",
			input: AddShipmentInput{
				FromName:        "Mark",
				FromEmail:       "testFrom@g.c",
				FromAddress:     "Lviv, 45",
				FromCountryCode: "UA",
				ToName:          "Iryna",
				ToEmail:         "testTo@g.c",
				ToAddress:       "Toronto, 34",
				ToCountryCode:   "CA",
				Weight:          234.4,
			},
			inputShipment: models.Shipment{
				FromName:        "Mark",
				FromEmail:       "testFrom@g.c",
				FromAddress:     "Lviv, 45",
				FromCountryCode: "UA",
				ToName:          "Iryna",
				ToEmail:         "testTo@g.c",
				ToAddress:       "Toronto, 34",
				ToCountryCode:   "CA",
				Weight:          234.4,
				Price:           3000,
			},
			mockBehaviur: func(r *mock_repositories.MockShipmentRepository, shipment models.Shipment) {
				r.EXPECT().CreateShipment(gomock.Eq(shipment)).Return(nil)
			},
			expectedPrice: 3000,
			expectedError: nil,
		},
		{
			name: "failed to add in database",
			input: AddShipmentInput{
				FromName:        "Mark",
				FromEmail:       "testFrom@g.c",
				FromAddress:     "Lviv, 45",
				FromCountryCode: "UA",
				ToName:          "Iryna",
				ToEmail:         "testTo@g.c",
				ToAddress:       "Toronto, 34",
				ToCountryCode:   "CA",
				Weight:          234.4,
			},
			inputShipment: models.Shipment{
				FromName:        "Mark",
				FromEmail:       "testFrom@g.c",
				FromAddress:     "Lviv, 45",
				FromCountryCode: "UA",
				ToName:          "Iryna",
				ToEmail:         "testTo@g.c",
				ToAddress:       "Toronto, 34",
				ToCountryCode:   "CA",
				Weight:          234.4,
				Price:           3000,
			},
			mockBehaviur: func(r *mock_repositories.MockShipmentRepository, shipment models.Shipment) {
				r.EXPECT().CreateShipment(gomock.Eq(shipment)).Return(errors.New("some db error"))
			},
			expectedPrice: 0,
			expectedError: errors.New("some db error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			shipmentRepo := mock_repositories.NewMockShipmentRepository(c)
			tC.mockBehaviur(shipmentRepo, tC.inputShipment)

			service := InitShipmentService(shipmentRepo)

			// Call method
			actualPrice, err := service.AddShipment(tC.input)

			// Require
			require.Equal(t, tC.expectedPrice, actualPrice)
			require.Equal(t, tC.expectedError, err)
		})
	}
}
