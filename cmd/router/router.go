package router

import (
	"fmt"
	"server1/cmd/handler"
	"server1/internal/middleware"

	"github.com/gin-gonic/gin"
)

// API generate api
func API(handler handler.Handler, serverPort string) {
	tokenController := handler.TokenController
	registerController := handler.RegisterController
	authController := handler.AuthController
	r := gin.Default()

	r.POST("/register", registerController.Register)

	r.POST("/login", authController.Authenticate)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	officeAuthorized := r.Group("/")
	officeAuthorized.Use(middleware.AuthRequired())
	{
		officeAuthorized.GET("/token/check", tokenController.Check)

	}

	if serverPort == "" {
		serverPort = "8080"
		fmt.Print("using default port : 8080")
	}
	r.Run(`:` + serverPort)
}
