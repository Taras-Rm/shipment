package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Taras-Rm/shipment/models"
	"github.com/Taras-Rm/shipment/services"
	mock_services "github.com/Taras-Rm/shipment/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_getShipmentById(t *testing.T) {
	type mockBehaviur func(r *mock_services.MockShipmentService, id uint, shipment models.Shipment)

	testCases := []struct {
		name                 string
		inputId              uint
		outputShipment       models.Shipment
		mockBehaviur         mockBehaviur
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			inputId: 2,
			outputShipment: models.Shipment{
				Id:              2,
				FromName:        "Mark",
				FromEmail:       "testFrom@g.c",
				FromAddress:     "Lviv, 45",
				FromCountryCode: "UA",
				ToName:          "Iryna",
				ToEmail:         "testTo@g.c",
				ToAddress:       "Toronto, 34",
				ToCountryCode:   "CA",
				Weight:          234.4,
				Price:           99.99,
			},
			mockBehaviur: func(r *mock_services.MockShipmentService, id uint, shipment models.Shipment) {
				r.EXPECT().GetShipmentByID(gomock.Eq(id)).Return(shipment, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"shipment":{"Id":2,"FromName":"Mark","FromEmail":"testFrom@g.c","FromAddress":"Lviv, 45","FromCountryCode":"UA","ToName":"Iryna","ToEmail":"testTo@g.c","ToAddress":"Toronto, 34","ToCountryCode":"CA","Weight":234.4,"Price":99.99}}`,
		},
		{
			name:           "some internal error",
			inputId:        2,
			outputShipment: models.Shipment{},
			mockBehaviur: func(r *mock_services.MockShipmentService, id uint, shipment models.Shipment) {
				r.EXPECT().GetShipmentByID(gomock.Eq(id)).Return(shipment, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"some internal error"}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			shipment := mock_services.NewMockShipmentService(c)
			tC.mockBehaviur(shipment, tC.inputId, tC.outputShipment)

			// Init endpoint
			api := gin.New()
			api.GET("/:id", getShipmentByID(shipment))

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/%d", tC.inputId), bytes.NewBufferString(""))

			// Make request
			api.ServeHTTP(w, req)

			// Require
			require.Equal(t, tC.expectedStatusCode, w.Code)
			require.Equal(t, tC.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_getAllShipments(t *testing.T) {
	type mockBehaviur func(r *mock_services.MockShipmentService, shipments []models.Shipment)

	testCases := []struct {
		name                 string
		outputShipments      []models.Shipment
		mockBehaviur         mockBehaviur
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			outputShipments: []models.Shipment{
				{
					Id:              2,
					FromName:        "Mark",
					FromEmail:       "testFrom@g.c",
					FromAddress:     "Lviv, 45",
					FromCountryCode: "UA",
					ToName:          "Iryna",
					ToEmail:         "testTo@g.c",
					ToAddress:       "Toronto, 34",
					ToCountryCode:   "CA",
					Weight:          234.4,
					Price:           99.99,
				},
				{
					Id:              3,
					FromName:        "Tom",
					FromEmail:       "testFrom@g.c",
					FromAddress:     "Lutsk, 34",
					FromCountryCode: "UA",
					ToName:          "Viktor",
					ToEmail:         "testTo@g.c",
					ToAddress:       "London, 32",
					ToCountryCode:   "UK",
					Weight:          5,
					Price:           234.78,
				},
			},
			mockBehaviur: func(r *mock_services.MockShipmentService, shipments []models.Shipment) {
				r.EXPECT().GetAllShipments().Return(shipments, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"shipments":[{"Id":2,"FromName":"Mark","FromEmail":"testFrom@g.c","FromAddress":"Lviv, 45","FromCountryCode":"UA","ToName":"Iryna","ToEmail":"testTo@g.c","ToAddress":"Toronto, 34","ToCountryCode":"CA","Weight":234.4,"Price":99.99},{"Id":3,"FromName":"Tom","FromEmail":"testFrom@g.c","FromAddress":"Lutsk, 34","FromCountryCode":"UA","ToName":"Viktor","ToEmail":"testTo@g.c","ToAddress":"London, 32","ToCountryCode":"UK","Weight":5,"Price":234.78}]}`,
		},
		{
			name:            "without shipments",
			outputShipments: []models.Shipment{},
			mockBehaviur: func(r *mock_services.MockShipmentService, shipments []models.Shipment) {
				r.EXPECT().GetAllShipments().Return(shipments, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"shipments":[]}`,
		},
		{
			name: "some internal error",
			mockBehaviur: func(r *mock_services.MockShipmentService, shipments []models.Shipment) {
				r.EXPECT().GetAllShipments().Return(shipments, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"some internal error"}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			shipment := mock_services.NewMockShipmentService(c)
			tC.mockBehaviur(shipment, tC.outputShipments)

			// Init endpoint
			api := gin.New()
			api.GET("", getAllShipments(shipment))

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", bytes.NewBufferString(""))

			// Make request
			api.ServeHTTP(w, req)

			// Require
			require.Equal(t, tC.expectedStatusCode, w.Code)
			require.Equal(t, tC.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_addShipment(t *testing.T) {
	type mockBehaviur func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput)

	testCases := []struct {
		name                 string
		fixturePath          string
		inputShipment        services.AddShipmentInput
		mockBehaviur         mockBehaviur
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			fixturePath: "./fixtures/shipments/add.ok.json",
			inputShipment: services.AddShipmentInput{
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
			mockBehaviur: func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {
				r.EXPECT().AddShipment(gomock.Eq(shipment)).Return(1000.5, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"price":1000.5}`,
		},
		{
			name:                 "Missing fromName",
			fixturePath:          "./fixtures/shipments/add.no_fromName.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing fromEmail",
			fixturePath:          "./fixtures/shipments/add.no_fromEmail.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing fromAddress",
			fixturePath:          "./fixtures/shipments/add.no_fromAddress.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing fromCountryCode",
			fixturePath:          "./fixtures/shipments/add.no_fromCountryCode.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing toName",
			fixturePath:          "./fixtures/shipments/add.no_toName.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing toEmail",
			fixturePath:          "./fixtures/shipments/add.no_toEmail.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing toAddress",
			fixturePath:          "./fixtures/shipments/add.no_toAddress.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing toCountryCode",
			fixturePath:          "./fixtures/shipments/add.no_toCountryCode.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Missing weight",
			fixturePath:          "./fixtures/shipments/add.no_weight.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:                 "Invalid weight",
			fixturePath:          "./fixtures/shipments/add.invalid_weight.json",
			mockBehaviur:         func(r *mock_services.MockShipmentService, shipment services.AddShipmentInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			shipment := mock_services.NewMockShipmentService(c)
			tC.mockBehaviur(shipment, tC.inputShipment)

			// Init endpoint
			api := gin.New()
			api.POST("", addShipment(shipment))

			// Input body preparing
			fixturedData, err := os.ReadFile(tC.fixturePath)
			require.NoError(t, err)

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(fixturedData))

			// Make request
			api.ServeHTTP(w, req)

			// Require
			require.Equal(t, tC.expectedStatusCode, w.Code)
			require.Equal(t, tC.expectedResponseBody, w.Body.String())
		})
	}
}
