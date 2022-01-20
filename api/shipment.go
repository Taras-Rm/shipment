package api

import (
	"net/http"
	"strconv"

	"github.com/Taras-Rm/shipment/services"
	"github.com/gin-gonic/gin"
)

func UseShipment(gr *gin.RouterGroup, shipmentService services.ShipmentService) {
	handler := gr.Group("shipment")

	// endpoints
	handler.GET("", getAllShipments(shipmentService))
	handler.POST("", addShipment(shipmentService))
	handler.GET(":id", getShipmentByID(shipmentService))
}

func getAllShipments(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get all shipments
		shipments, err := shipmentService.GetAllShipments()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "Shipments received!",
			"shipments": shipments,
		})
	}
}

func addShipment(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shipmentReq *services.ShipmentRequest

		// check shipment request
		err := c.BindJSON(&shipmentReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		// validate shipment request
		err = services.IsValidShipmentRequest(shipmentReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		// add new shipment to database
		price, err := shipmentService.AddShipment(shipmentReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Shipment is added!",
			"price":   price,
		})
	}
}

func getShipmentByID(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get ID param
		id := c.Param("id")
		shipmentId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		// get shipment by ID
		shipment, err := shipmentService.GetShipmentByID(uint(shipmentId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Shipment is getted!",
			"shipment": shipment,
		})
	}
}
