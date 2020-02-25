package user

import "time"

// User represents user that use application.
type User struct {
	UserName  string     `json:"user_name"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

/*
FindByEmailOrCreateUserRequest represents paramater to find user by email
or created new user if not founded.
*/
type FindByEmailOrCreateUserRequest struct {
	Email string
	Name  string
}
