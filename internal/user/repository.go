package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Repository implements operations for working with user data storage.
type Repository struct {
	db *gorm.DB
}

// NewRepository create a new user repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindByUserName find user by its email.
func (r *Repository) FindByUserName(ctx context.Context, userName string) (*User, error) {

	var user User
	if err := r.db.Where(&User{UserName: userName}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert a user into data storage.
func (r *Repository) Insert(ctx context.Context, user User) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
