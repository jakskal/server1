package handler

import (
	"server1/internal/auth"
	"server1/internal/register"
	"server1/internal/token"
)

// Handler wrap software controller.
type Handler struct {
	AuthController     *auth.Controller
	RegisterController *register.Controller
	TokenController    *token.Controller
}

// NewHandler creates a new Handler.
func NewHandler(
	authController *auth.Controller,
	registerController *register.Controller,
	tokenController *token.Controller,
) *Handler {
	return &Handler{
		TokenController:    tokenController,
		AuthController:     authController,
		RegisterController: registerController,
	}
}
