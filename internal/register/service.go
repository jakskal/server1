package register

import (
	"context"

	"server1/internal/user"
	"server1/pkg/hash"
)

// Service implement business operations for working with registrant.
type Service struct {
	userService user.Service
}

// NewService creates a new registrant service.
func NewService(userService user.Service) Service {
	return Service{
		userService: userService,
	}
}

// Register creates an registrant.
func (s *Service) Register(ctx context.Context, req user.User) (*string, error) {
	hashedPassword, err := hash.Password(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hashedPassword

	registeredUser, err := s.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &registeredUser.UserName, nil
}
