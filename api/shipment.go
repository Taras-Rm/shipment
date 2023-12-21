package api

import (
	"errors"
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
			newErrorResponse(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"shipments": shipments,
		})
	}
}

func addShipment(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp services.AddShipmentInput
		if err := c.BindJSON(&inp); err != nil {
			newErrorResponse(c, http.StatusBadRequest, errors.New("invalid input body"))
			return
		}

		// validate add shipment request
		if err := inp.Validate(); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		// add new shipment to database
		price, err := shipmentService.AddShipment(inp)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"price": price,
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
