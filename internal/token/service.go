package token

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Service implement logical method to create token.
type Service struct {
}

// NewService create a new token service.
func NewService() *Service {
	return &Service{}
}

var loginExpirationDuration = time.Duration(1) * time.Hour * 24 * 365
var signingMethod = jwt.SigningMethodHS256

// CreateToken create token
func (as *Service) CreateToken(ctx context.Context, req *CreateTokenRequest) (*Token, error) {
	signatureKey := os.Getenv("JWT_SIGNATURE_KEY")
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(loginExpirationDuration).Unix(),
		},
		UserName:  req.UserName,
		Name:      req.Name,
		Role:      req.Role,
		TimeStamp: fmt.Sprint(req.CreatedAt.Unix()),
	}

	jtoken := jwt.NewWithClaims(
		signingMethod,
		claims,
	)

	ts, err := jtoken.SignedString([]byte(signatureKey))

	if err != nil {
		return nil, err
	}
	token := &Token{
		AccessToken: ts,
	}

	return token, nil
}
