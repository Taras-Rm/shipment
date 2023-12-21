package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Taras-Rm/shipment/helpers"
	"github.com/Taras-Rm/shipment/services"
	mock_services "github.com/Taras-Rm/shipment/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_getShipmentById(t *testing.T) {
	type mockBehaviour func(s *mock_services.MockShipmentService, shipment *services.ShipmentResponse, shipmentId uint)

	type responseBody struct {
		Message  string                     `json:"message"`
		Shipment *services.ShipmentResponse `json:"shipment"`
		Error    *string                    `json:"error"`
	}

	shipment := services.ShipmentResponse{
		ID:              5,
		FromName:        "Taras",
		FromEmail:       "testEmailFrom@g.c",
		FromAddress:     "Kyiv",
		FromCountryCode: "UA",
		ToName:          "Vitaliy",
		ToEmail:         "testEmailTo@g.c",
		ToAddress:       "Warshaw",
		ToCountryCode:   "PL",
		Weight:          5.5,
	}

	testCases := []struct {
		name                 string
		mockBehavior         mockBehaviour
		requestIdParam       uint
		outputShipment       *services.ShipmentResponse
		expectedStatusCode   int
		expectedResponseBody responseBody
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_services.MockShipmentService, shipment *services.ShipmentResponse, shipmentId uint) {
				s.EXPECT().GetShipmentByID(gomock.Eq(shipmentId)).Return(shipment, nil)
			},
			requestIdParam:     5,
			outputShipment:     &shipment,
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: responseBody{
				Message:  "Shipment is getted!",
				Shipment: &shipment,
			},
		},
		{
			name: "can not found shipment with id",
			mockBehavior: func(s *mock_services.MockShipmentService, shipment *services.ShipmentResponse, shipmentId uint) {
				s.EXPECT().GetShipmentByID(gomock.Eq(shipmentId)).Return(shipment, errors.New("can not fount shipment"))
			},
			requestIdParam:     12,
			outputShipment:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: responseBody{
				Message: "Server error",
				Error:   helpers.StrToPointerStr("can not fount shipment"),
			},
		},
		{
			name: "invalid shipment id",
			mockBehavior: func(s *mock_services.MockShipmentService, shipment *services.ShipmentResponse, shipmentId uint) {
				s.EXPECT().GetShipmentByID(gomock.Eq(shipmentId)).Return(shipment, errors.New("can not fount shipment"))
			},
			requestIdParam:     12,
			outputShipment:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: responseBody{
				Message: "Server error",
				Error:   helpers.StrToPointerStr("can not fount shipment"),
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Initialize dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			serv := mock_services.NewMockShipmentService(c)

			tC.mockBehavior(serv, tC.outputShipment, tC.requestIdParam)

			// Initialize endpoint
			r := gin.New()
			r.GET("/:id", getShipmentByID(serv))

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+strconv.Itoa(int(tC.requestIdParam)), bytes.NewBufferString(""))

			// Make request
			r.ServeHTTP(w, req)

			// Require status code
			require.Equal(t, tC.expectedStatusCode, w.Code)

			// Require response body
			var actual responseBody
			err := json.Unmarshal(w.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tC.expectedResponseBody, actual)
		})
	}
}

func TestHandler_getAllShipments(t *testing.T) {
	type mockBehaviour func(s *mock_services.MockShipmentService, shipments []services.ShipmentResponse)

	type responseBody struct {
		Message   string                      `json:"message"`
		Shipments []services.ShipmentResponse `json:"shipments"`
		Error     *string                     `json:"error"`
	}

	shipment1 := services.ShipmentResponse{
		ID:              1,
		FromName:        "Taras",
		FromEmail:       "testEmailFrom@g.c",
		FromAddress:     "Kyiv",
		FromCountryCode: "UA",
		ToName:          "Vitaliy",
		ToEmail:         "testEmailTo@g.c",
		ToAddress:       "Warshaw",
		ToCountryCode:   "PL",
		Weight:          5.5,
		Price:           1200,
	}

	shipment2 := services.ShipmentResponse{
		ID:              2,
		FromName:        "Tom",
		FromEmail:       "testEmailFrom@g.c",
		FromAddress:     "Lviv",
		FromCountryCode: "UA",
		ToName:          "Tim",
		ToEmail:         "testEmailTo@g.c",
		ToAddress:       "Paris",
		ToCountryCode:   "FR",
		Weight:          12.57,
		Price:           245.5,
	}

	testCases := []struct {
		name                 string
		mockBehavior         mockBehaviour
		outputShipments      []services.ShipmentResponse
		expectedStatusCode   int
		expectedResponseBody responseBody
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_services.MockShipmentService, shipments []services.ShipmentResponse) {
				s.EXPECT().GetAllShipments().Return(shipments, nil)
			},
			expectedStatusCode: http.StatusOK,
			outputShipments:    []services.ShipmentResponse{shipment1, shipment2},
			expectedResponseBody: responseBody{
				Message:   "Shipments received!",
				Shipments: []services.ShipmentResponse{shipment1, shipment2},
			},
		},
		{
			name: "with internal error",
			mockBehavior: func(s *mock_services.MockShipmentService, shipments []services.ShipmentResponse) {
				s.EXPECT().GetAllShipments().Return(shipments, errors.New("some internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			outputShipments:    nil,
			expectedResponseBody: responseBody{
				Message: "Server error",
				Error:   helpers.StrToPointerStr("some internal error"),
			},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// Initialize dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			serv := mock_services.NewMockShipmentService(c)

			tC.mockBehavior(serv, tC.outputShipments)

			// Initialize endpoint
			r := gin.New()
			r.GET("", getAllShipments(serv))

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", bytes.NewBufferString(""))

			// Make request
			r.ServeHTTP(w, req)

			// Require status code
			require.Equal(t, tC.expectedStatusCode, w.Code)

			// Require response body
			var actual responseBody
			err := json.Unmarshal(w.Body.Bytes(), &actual)
			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, tC.expectedResponseBody, actual)
		})
	}
}
