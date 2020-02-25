package register

import (
	"net/http"
	"server1/internal/user"

	"github.com/gin-gonic/gin"
)

// Controller wraps an registrant service and implement gin context
type Controller struct {
	service Service
}

// Register regitser a user.
func (cr *Controller) Register(c *gin.Context) {
	var registrant user.User
	err := c.BindJSON(&registrant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return

	}

	ctx := c.Request.Context()

	username, err := cr.service.Register(ctx, registrant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "register succes, now you can login.",
		"username": username,
	})
}

// NewController create a new User Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
