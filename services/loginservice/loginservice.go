package loginservice

import (
	"errors"
	"ginhello/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// Authenticate authenticates user credentials
func Authenticate(username, password string) (*models.User, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Check if user exists and password is correct
	if user == nil || !user.CheckPassword(password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
