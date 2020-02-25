package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token represents user tokens for authentication & authorization process (TODO).
type Token struct {
	AccessToken string `json:"access_token"`
}

// Claims represents information that contained in claim
type Claims struct {
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	TimeStamp string `json:"timestamp"`
	jwt.StandardClaims
}

type (
	// CreateTokenRequest represents parameters to create token
	CreateTokenRequest struct {
		UserName  string
		Name      string
		Role      string
		CreatedAt *time.Time
	}
)
