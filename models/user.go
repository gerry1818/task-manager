package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	PasswordHash string `gorm:"not null"` // Ensure this field is defined
}

// HashPassword hashes the user's password and stores it in the PasswordHash field.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword) // Use PasswordHash field
	return nil
}

// CheckPassword compares the stored password hash with the provided password.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
