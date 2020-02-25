package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// RepositorySystem defines operations for working with user data storage.
type RepositorySystem interface {
	Insert(ctx context.Context, user User) error
	FindByUserName(ctx context.Context, username string) (*User, error)
}

// Service implement business operations for working with user.
type Service struct {
	userRepo RepositorySystem
}

// NewService creates a new user service.
func NewService(userRepo RepositorySystem) Service {
	return Service{
		userRepo: userRepo,
	}
}

// CreateUser creates an user.
func (s *Service) CreateUser(ctx context.Context, user User) (*User, error) {
	err := s.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByUserName find user by its id.
func (s *Service) FindByUserName(ctx context.Context, userEmail string) (*User, error) {
	user, err := s.userRepo.FindByUserName(ctx, userEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
