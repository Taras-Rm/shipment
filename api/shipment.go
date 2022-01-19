package api

import (
	"net/http"

	"github.com/Taras-Rm/shipment/services"
	"github.com/gin-gonic/gin"
)

func UseShipment(gr *gin.RouterGroup, shipmentService services.ShipmentService) {
	handler := gr.Group("shipment")
	handler.GET("", getAllShipments(shipmentService))
	//handler.POST("", addShipment(shipmentService))
	//handler.GET(":id", getShipment(shipmentService))
}

func getAllShipments(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {

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
