package api

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, errorResponse{
		Error: err.Error(),
	})
}
