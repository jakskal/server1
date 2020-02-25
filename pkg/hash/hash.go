package hash

import "golang.org/x/crypto/bcrypt"

// Password hashing plain password
func Password(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	return string(bytes), err
}
