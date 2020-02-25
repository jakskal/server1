package auth

import (
	"context"
	"errors"

	"server1/internal/token"
	"server1/internal/user"

	"golang.org/x/crypto/bcrypt"
)

// Service implements auth service interface.
type Service struct {
	userService  user.Service
	tokenService *token.Service
}

// NewService creates a new auth Service.
func NewService(userService user.Service, tokenService *token.Service) Service {
	return Service{
		userService:  userService,
		tokenService: tokenService,
	}
}

// Authenticate logging in an user.
func (s *Service) Authenticate(ctx context.Context, req *Auth) (*token.Token, error) {
	user, err := s.userService.FindByUserName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	v, err := comparePasswords(user.Password, req.Password)
	if v == false {
		return nil, errors.New("invalid password" + err.Error())
	}

	token, err := s.tokenService.CreateToken(ctx,
		&token.CreateTokenRequest{
			UserName:  user.UserName,
			Name:      user.Name,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		})
	if err != nil {
		return nil, err
	}

	return token, nil

}

func comparePasswords(hashedPassword string, plainPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false, err
	}
	return true, nil
}
