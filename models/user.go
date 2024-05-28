package models

import (
	"ginhello/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, err // Other error occurred
	}
	return &user, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	u, err := GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil // User not found
		}
		return "", err // Other error occurred
	}

	if !u.CheckPassword(password) {
		return "", bcrypt.ErrMismatchedHashAndPassword // Password mismatch
	}

	token, err := token.GenerateToken(u.Id)
	if err != nil {
		return "", err // Error generating token
	}

	return token, nil
}

// SetPassword hashes the plain password and sets the HashedPassword field
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CreateUser creates a new user in the database
func CreateUser(user *User) error {

	return DB.Create(user).Error
}
