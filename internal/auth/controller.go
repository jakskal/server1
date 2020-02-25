package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an login service and implement gin context.
type Controller struct {
	service Service
}

// NewController create a new login controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Authenticate logging in an user.
func (cr *Controller) Authenticate(c *gin.Context) {
	var request *Auth
	err := c.Bind(&request)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	response, err := cr.service.Authenticate(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to auth",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
