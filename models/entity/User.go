package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model
	RoleID int  `json:"role_id"`
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Register struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Role            int    `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hash), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err
}
